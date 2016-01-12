package controlpanel

import (
	"bufio"
	"db"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"watcher"
    "printer"
)

//Control Panel that will undestand our commands by watching input in the terminal
func ControlPanelConsole(conn net.Conn, bot *watcher.Watcher, DBConn *db.Crud) {
	//Using bufio reader to read in input from the os.Stdin which is the console
	bio := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Waiting for input: ")
		//var input string //Input variable to hold input
		//fmt.Scanln(&input) //Scan the input to read what was typed

		//Line will be an entire console line read by our bufio reader
		line, _, err := bio.ReadLine()

		if err != nil {
			//Logging any errors that might occurr
			log.Printf("%q\n", err)
		}

		//Splitting the input so we can use multiple parts of the inputted message
		splitted := strings.Split(string(line), " ")

		//Actions depending on scanned input
		switch splitted[0] {
		case "Hello": //Print hello message
			conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + helloCommand()))
		case "Block": //Block user from using the bot
			conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + blockCommand(splitted[1], DBConn)))
		case "UnBlock":
			conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + unblockCommand(splitted[1], DBConn)))
        case "Extract":
            conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + extractCommand(splitted[1], DBConn)))
		case "Quit": //Quit
			conn.Write([]byte("QUIT " + "\r\n"))
		}
	}
}

//Function that handles the hello command, prints a greetings message
func helloCommand() string {
	return "Hello, I'm At The ControlPanel!\r\n"
}

//Function that handles the block command, blocks given user from using the bot
func blockCommand(user string, c *db.Crud) string {

	//Check if the user is blocked already
	if db.IsUserBlocked(user, c) {
		return user + " is already blocked from using the bot or he/she doesn't exist!\r\n"
	}

	//if the user is not blocked then proceed with blocking the user
	db.BlockUser(user, c)
	return "Blocking " + user + " from using the bot!\r\n"
}

//Function that handles the unblock command, unblock given user so he/she can use the bot again
func unblockCommand(user string, c *db.Crud) string {

	//If the user is blocked proceed with unblocking
	if db.IsUserBlocked(user, c) {
		//Unblock user removing it form the block list
		db.UnBlockUser(user, c)
		return "UnBlocking " + user + " from the block list!\r\n"
	}

	//The user is not blocked and cannot be unblocked
	return user + " is not blocked from using the bot or he/she doesn't exist!\r\n"
}

func extractCommand(user string, c *db.Crud) string {
    userid := db.GetUserByName(user, c)
    messages := db.GetMessages(userid, c)

    printer.TextToFile(messages, user)

    return user + " is going to get logged.\r\n"
}
