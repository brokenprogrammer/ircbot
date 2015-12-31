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
		case "Quit": //Quit
			conn.Write([]byte("QUIT " + "\r\n"))
		}
	}
}

func helloCommand() string {
	return "Hello, I'm At The ControlPanel!\r\n"
}

func blockCommand(user string, c *db.Crud) string {
	//TODO: Code that adds user to database, Code in controlpanel.go that checks if user is blocked
	if db.IsUserBlocked(user, c) {
		return user + " is already blocked from using the bot or he/she doesn't exist!\r\n"
	}

	db.BlockUser(user, c)
	return "Blocking " + user + " from using the bot!\r\n"
}
