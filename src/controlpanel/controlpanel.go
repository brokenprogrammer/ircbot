package controlpanel

import (
	"db"
	"fmt"
	"net"
	"strconv"
	"strings"
	"time"
	"watcher"
)

//Control Panel that watches input from the chat.
func ControlPanel(conn net.Conn, bot *watcher.Watcher, c chan string, DBConn *db.Crud) {
	for {

		select {
		//Incase we got a message through the channel
		case message := <-c:
			//Call the readMessage function that handles code depending on commands
			readMessage(message, bot, conn, DBConn)
		//Incase we didn't we timeout after 1 Second
		case <-time.After(time.Second * 1):
			//fmt.Println("Def") //Timeout
		}
	}
}

//readMessage reads messages user sent to the chat
//If it recognize a command it calls a function depending on command.
func readMessage(msg string, bot *watcher.Watcher, conn net.Conn, DBConn *db.Crud) {
	//Splits the string into a slice so we can use multiple parts of the message
	splitted := strings.Split(msg, " ")

	//Switch through the incoming message
	switch strings.TrimSpace(splitted[3]) {
	//Incase user wrote !help, Display the help message
	case ":!help":
		//Calling function displaying help message
		helpCommand(bot, conn)
	case ":!status":
		//Calling the function displaying the status message
		statusCommand(bot, conn, DBConn)
	default:
		fmt.Println(msg)
	}
}

//helpCommand is a command function called by the readMessage function incase a
//user requested to display the help message.
func helpCommand(bot *watcher.Watcher, conn net.Conn) {
	conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + "BrokenBot Commands: !help - Display help message, !status - Bot Status \r\n"))
}

//statusCommand is a command function called by readMessage function incase a user
//Requested to view the status of the application.
func statusCommand(bot *watcher.Watcher, conn net.Conn, DBConn *db.Crud) {
	conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + "BrokenBot Status: Uptime: " + time.Since(bot.RanFor).String() + " Tracking: " + strconv.Itoa(db.GetAmmountOfUsers(DBConn)) + " users and tracking: " + strconv.Itoa(db.GetAmmountOfMessages(DBConn)) + " messages. \r\n"))
}
