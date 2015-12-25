package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"watcher"
)

func main() {
	fmt.Println("Hello World!")

	bot := watcher.NewBot()
	conn, _ := bot.Connect()
	defer conn.Close()

	fmt.Println(bot)

	//Write to the IRC server
	conn.Write([]byte("NICK " + bot.Nick + "\r\n"))                         //IRC server requests a nickname for the user
	conn.Write([]byte("USER " + bot.Nick + " 8 *  : " + bot.Nick + "\r\n")) //IRC server always requests a realname in this format
	conn.Write([]byte("JOIN " + bot.Channel + "\r\n"))                      //Using the irc JOIN command to join the channel our bot uses.
	conn.Write([]byte("PRIVMSG " + bot.Channel + " :Hello World!\r\n"))

	//Using a Go Routine to handle a Control Panel for the bot simultaniously as the bot is running
	go ControlPanel(conn, bot)

	//The bufio reader will read data we get from our connection and return it as a string.
	connBuff := bufio.NewReader(conn)

	for {
		str, err := connBuff.ReadString('\n')
		if len(str) > 0 { //If there is a message from the server
			fmt.Println(str) //Print it out

			//Staying connected to the IRC server
			splitted := strings.Split(str, " ") //Split the string into a slice
			if splitted[0] == "PING" {          //If the IRC Server is pinging us
				fmt.Println(splitted)
				conn.Write([]byte("PONG " + splitted[1] + "\r\n"))                            //Respond back with a PONG
				conn.Write([]byte("PRIVMSG " + bot.Channel + " :Hello I'm Still here! \r\n")) //Tell the chat you're still here
			}
		}
		if err != nil {
			break
		}
	}
}

//Experimenting with a ControlPanel for the bot.
func ControlPanel(conn net.Conn, bot *watcher.Watcher) {
	for {
		fmt.Println("Waiting for input: ")

		var input string

		fmt.Scan(&input)

		switch input {
		case "Hello":
			conn.Write([]byte("PRIVMSG " + bot.Channel + " :Hello I'm At The ControllPanel! \r\n"))
		case "Quit":
			conn.Write([]byte("QUIT " + "\r\n"))
		}
	}
}
