package controlpanel

import (
	"fmt"
	"net"
	//"time"
	"watcher"
)

//Experimenting with a ControlPanel for the bot.
func ControlPanel(conn net.Conn, bot *watcher.Watcher, c chan string) {
	for {
		//msg := <-c

		/*switch msg {
		default:
			conn.Write([]byte("PRIVMSG " + bot.Channel + " :" + msg + "\r\n"))
		}*/

		/*select {
		case msg := <-c:
			fmt.Println("Do work!", msg)
		case <-time.After(time.Millisecond * 1):
			fmt.Println("Def")
		}*/

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
