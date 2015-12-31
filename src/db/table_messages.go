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

//Function to get ammount of messages stored in the db
func GetAmmountOfMessages(c *Crud) int {
	//Get all the rows from the messages table. This is not the most efficient way but will do for a small app.
	rows, err := c.DBInstance.Query("SELECT COUNT(*) FROM messages")

	//If there is an error we close down the function
	if err != nil {
		//User doesnt exist
		log.Print(err)
		return 0
	}

	//Defer close the Query
	defer rows.Close()

	//Loop through the one row and print it out to the console.
	for rows.Next() {
		var result int

		//Scan the id and name from the found row
		rows.Scan(&result)
		log.Print("Result: ", result)

		//Log successmessage and return no errors
		log.Printf("Successfully Found Messages: %s", result)
		return result
	}

	//Nothing found, return 0
	return 0
}
