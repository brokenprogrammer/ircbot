package db

import (
	"log"
)

type Message struct {
	id      int
	Userid  int
	Message string
}

//Store a message into the message table
func StoreMessage(userid int, message string, c *Crud) {
	//Insert message into the database
	c.insertMessage(GetMessagesTable(), userid, message)

	//Log successmessage and return no errors
	log.Printf("Successfully Inserted Message To Database: %s \n", message)
}

func GetMessages() {

}

func DeleteMessages() {

}

//Returns the string used for database queries connected with this table
func GetMessagesTable() string {
	return "messages(userid, message)"
}
