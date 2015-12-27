package main

import (
	"bufio"
	"controlpanel"
	//"db"
	"fmt"
	_ "net"
	"strings"
	"watcher"
)

func main() {
	c := make(chan string)

	//Our Bot Instance
	bot := watcher.NewBot()

	//Connection instance copy
	conn, _ := bot.Connect()

	//defer statement to close the connection when this function ends
	defer conn.Close()

	//db.InitDB()

	//Write to the IRC server
	conn.Write([]byte("NICK " + bot.Nick + "\r\n"))                         //IRC server requests a nickname for the user
	conn.Write([]byte("USER " + bot.Nick + " 8 *  : " + bot.Nick + "\r\n")) //IRC server always requests a realname in this format
	conn.Write([]byte("JOIN " + bot.Channel + "\r\n"))                      //Using the irc JOIN command to join the channel our bot uses.
	conn.Write([]byte("PRIVMSG " + bot.Channel + " :Hello World!\r\n"))

	//Using a Go Routine to handle a Control Panel for the bot simultaniously as the bot is running
	go controlpanel.ControlPanel(conn, bot, c)

	//Using another Go Routine to handle a Control Panel that listens to input from the terminal
	go controlpanel.ControlPanelConsole(conn, bot)

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
				//c <- splitted[3]
				c <- str
			}
		}
		if err != nil {
			break //Break out of the loop if there is an error
		}
	}
}
