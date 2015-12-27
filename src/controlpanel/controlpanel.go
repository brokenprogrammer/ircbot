package controlpanel

import (
	"fmt"
	"net"
	"strings"
	"time"
	"watcher"
)

//Control Panel that watches input from the chat.
func ControlPanel(conn net.Conn, bot *watcher.Watcher, c chan string) {
	for {

		select {
		//Incase we got a message through the channel
		case message := <-c:
			//Call the readMessage function that handles code depending on commands
			readMessage(message, bot, conn)
		//Incase we didn't we timeout after 1 Second
		case <-time.After(time.Second * 1):
			//fmt.Println("Def") //Timeout
		}
	}
}

//readMessage reads messages user sent to the chat
//If it recognize a command it calls a function depending on command.
func readMessage(msg string, bot *watcher.Watcher, conn net.Conn) {

	//Switch through the incoming message
	switch strings.TrimSpace(msg) {
	//Incase user wrote !help, Display the help message
	case ":!help":
		//Calling function displaying help message
		helpCommand(bot, conn)
	default:
		fmt.Println(msg)
	}
}

//helpCommand is a command function called by the readMessage function incase a
//user requested to display the help message.
func helpCommand(bot *watcher.Watcher, conn net.Conn) {
	conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + "BrokenBot Commands: !help - Display help message, !status - Bot Status \r\n"))
}

func status() {

}
