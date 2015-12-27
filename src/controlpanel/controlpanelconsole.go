package controlpanel

import (
	"fmt"
	"net"
	"watcher"
)

//Control Panel that will undestand our commands by watching input in the terminal
func ControlPanelConsole(conn net.Conn, bot *watcher.Watcher) {
	for {
		fmt.Println("Waiting for input: ")
		var input string //Input variable to hold input

		fmt.Scan(&input) //Scan the input to read what was typed

		//Actions depending on scanned input
		switch input {
		case "Hello": //Print hello message
			conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + hello()))
		case "Quit": //Quit
			conn.Write([]byte("QUIT " + "\r\n"))
		}
	}
}

func hello() string {
	return "Hello, I'm At The ControlPanel!\r\n"
}
