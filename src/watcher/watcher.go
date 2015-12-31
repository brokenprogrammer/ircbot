package watcher

import (
	"log"
	"net"
)

type Watcher struct {
	Server  string   //Server adress
	Port    string   //Server Port
	Nick    string   //Watchers Nickname
	Channel string   //Channel to join
	Conn    net.Conn //Connection
}

//A factory function for our Watcher structure
func NewBot(server, port, nick, channel string) *Watcher {
	return &Watcher{
		Server:  server,  //Server our IRC uses
		Port:    port,    //Port for IRC
		Nick:    nick,    //Name of our bot
		Channel: channel, //Channel bot is connecting to
		Conn:    nil,     //The conn will get initialized from the Connect function
	}
}

//This function starts a new dial connection using the information our Watcher structure has in it.
func (w *Watcher) Connect() (net.Conn, error) {
	//Start a new dial connection to the server.
	conn, err := net.Dial("tcp", w.Server+":"+w.Port)
	if err != nil {
		//If an error occurr we return the error message
		log.Fatal("Unable to connect to irc: ", err)
	}
	//Set the Watcher structures conn to the newly created connection
	w.Conn = conn
	//Print out the success connection message
	log.Printf("Connected to %s (%s)", w.Server, w.Conn.RemoteAddr())
	//Return the connection and no error.
	return w.Conn, nil
}
