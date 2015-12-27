package controlpanel

import (
	"fmt"
	"net"
	"time"
	"watcher"
)

//Control Panel that watches input from the chat.
func ControlPanel(conn net.Conn, bot *watcher.Watcher, c chan string) {
	for {

		select {
		//Incase we got a message through the channel
		case msg := <-c:
			fmt.Println("Do work!", msg)
		//Incase we didn't we timeout after 1 Second
		case <-time.After(time.Second * 1):
			//fmt.Println("Def") //Timeout
		}
	}
}
