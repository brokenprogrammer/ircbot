package db

import (
	"log"
)

type User struct {
	id   int
	Name string
}

//Returns the string used for database queries connected with this table
func GetUsersTable() string {
	return "users(name)"
}

//Function to check user towards the database
func CheckUser(username string, c *Crud) bool {
	//Get all the rows from the users table
	rows, err := c.DBInstance.Query("SELECT * FROM users")

	//If there is an error we close down the function
	if err != nil {
		//User doesnt exist
		log.Print(err)
		return true
	}

	//Defer close the Query
	defer rows.Close()

	//Loop through the rows and print them out to the console.
	for rows.Next() {
		var id int
		var name string
		//Scan the id and name from the current row
		rows.Scan(&id, &name)
		log.Print(id, name)
		//if the row name is same as the username provided to the function then it exists
		if name == username {
			//Print that user exists and return
			log.Print(name, " Already Exists")
			return true
		}
	}

	//Log successmessage and return no errors
	log.Printf("Successfully Checked User In Database: %s \n", username)

	//If the function got this far then the user doesn't exist in the DB.
	//Calling the function that adds users to the database
	NewUser(username, c)

	return false
}

//Function used to insert a new user into the database
func NewUser(username string, c *Crud) {
	//Insert user into the database
	c.Insert(GetUsersTable(), username)

	//Log successmessage and return no errors
	log.Printf("Successfully Inserted User To Database: %s \n", username)
}
