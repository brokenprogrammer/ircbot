package main

import (
	"bufio"
	"config"
	"controlpanel"
	"db"
	"fmt"
	"formatter"
	_ "net"
	"strings"
	"watcher"
)

func main() {
	//Initialize the configurations structure
	config.InitCfg()

	//Initialize the database, migrating tables.
	db.InitDB(config.MainCFG.DBDriver, config.MainCFG.DBPath)

	//Create a new DBConnection.
	DBConn := db.NewCrud(config.MainCFG.DBDriver, config.MainCFG.DBPath)

	//DBConn.Insert(db.GetUsersTable(), "Larry")
	//DBConn.Update("users", "name", "NewLarry", 4)
	//DBConn.Delete("users", 4)
	db.CheckUser("BadBob", DBConn)

	//Channel to send Chat input to the different Go Routines
	c := make(chan string)

	//Our Bot Instance
	bot := watcher.NewBot(config.MainCFG.Server, config.MainCFG.Port, config.MainCFG.Nick, config.MainCFG.Channel)

	//Connection instance copy
	conn, _ := bot.Connect()

	//defer statement to close the connection when this function ends
	defer conn.Close()

	//Write to the IRC server
	conn.Write([]byte("NICK " + bot.Nick + "\r\n"))                         //IRC server requests a nickname for the user
	conn.Write([]byte("USER " + bot.Nick + " 8 *  : " + bot.Nick + "\r\n")) //IRC server always requests a realname in this format
	conn.Write([]byte("JOIN " + bot.Channel + "\r\n"))                      //Using the irc JOIN command to join the channel our bot uses.
	conn.Write([]byte("PRIVMSG " + bot.Channel + " :Hello World!\r\n"))

	//Using a Go Routine to handle a Control Panel for the bot simultaniously as the bot is running
	go controlpanel.ControlPanel(conn, bot, c, DBConn)

	//Using another Go Routine to handle a Control Panel that listens to input from the terminal
	go controlpanel.ControlPanelConsole(conn, bot, DBConn)

	//The bufio reader will read data we get from our connection and return it as a string.
	connBuff := bufio.NewReader(conn)

	//The infinite loop that reads the chat untill the connection is lost (Or we Quit).
	for {
		str, err := connBuff.ReadString('\n')

		if len(str) > 0 { //If there is a message from the server
			fmt.Println(str) //Print it out

			//Staying connected to the IRC server
			splitted := strings.Split(str, " ") //Split the string into a slice
			if splitted[0] == "PING" {          //If the IRC Server is pinging us
				conn.Write([]byte("PONG " + splitted[1] + "\r\n"))                            //Respond back with a PONG
				conn.Write([]byte("PRIVMSG " + bot.Channel + " :Hello I'm Still here! \r\n")) //Tell the chat you're still here
			}

			//Sending the string through channel to our controlpanel
			if splitted[1] == "PRIVMSG" {
				//TODO: Make use of the structs in all the table files. (Look at other API repos)

				//Check if the user is stored in our Database
				db.CheckUser(formatter.GetUsername(splitted[0]), DBConn)

				//Get the id of the user who typed the message
				userid := db.GetUserByName(formatter.GetUsername(splitted[0]), DBConn)

				//Store the message in the database
				db.StoreMessage(userid, formatter.ExtractMessage(str), DBConn)

				//Send the whole string to the channel if the user isn't blocked
				if !db.IsUserBlocked(formatter.GetUsername(splitted[0]), DBConn) {
					c <- str
				}
			}
		}
		if err != nil {
			break //Break out of the loop if there is an error
		}
	}
}
