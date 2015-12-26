package controlpanel

import (
	"fmt"
	"net"
	"watcher"
)

//Experimenting with a ControlPanel for the bot.
func ControlPanel(conn net.Conn, bot *watcher.Watcher) {
	for {
		fmt.Println("Waiting for input: ")

		var input string

		fmt.Scan(&input)

		switch input {
		case "Hello":
			conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + hello()))
		case "Quit":
			conn.Write([]byte("QUIT " + "\r\n"))
		}
	}
}

func hello() string {
	return "Hello, I'm At The ControlPanel!\r\n"
}
