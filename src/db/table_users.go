package db

import (
	"log"
	"strconv"
)

type User struct {
	id   int
	Name string
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

//Deletes an user from the database
func DeleteUser(userid int, c *Crud) {
	//Deletes user with the id userid
	err := c.Delete(GetUsersTableRaw(), userid)

	//If we find errors then print it out
	if err != nil {
		log.Print(err)
	}

	//Print success message
	log.Print("Deleted user: ", userid)
}

//Returns the string used for database queries connected with this table
func GetUsersTable() string {
	return "users(name)"
}

//Returns raw version of the users table
func GetUsersTableRaw() string {
	return "users"
}

func GetUserByID() {

}

//Function that gets the user in the database by just its name
func GetUserByName(username string, c *Crud) int {
	var userid int

	//Get all the rows from the users table where the name is same as provided username
	rows, err := c.DBInstance.Query("SELECT * FROM users WHERE name='" + username + "'")

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
		var id int
		var name string
		//Scan the id and name from the found row
		rows.Scan(&id, &name)
		log.Print("User: ", id, name)

		//Set the user id to the found user
		userid = id

		//Log successmessage and return no errors
		log.Printf("Successfully Found User: %s \n", username)
		return userid
	}

	//Log successmessage and return no errors
	log.Printf("Couldn't Find User: %s \n", username)
	return 0
}

//Gets all the users messages printed out to a text file
func GetUserMessages() {

}

//Checks if the user is blocked from using the bot returns true if the user is blocked.
func IsUserBlocked(username string, c *Crud) bool {
	//Get the user id by checking the username.

	uid := GetUserByName(username, c)

	//If the userid is 0 (Doesn't exist).
	if uid == 0 {
		//Act as blocked and return true.
		log.Print("UserID was not found, Acting as blocked.")
		return true
	}

	//Get the row from the blocked database where userid is the same as we got from GetUserByName.
	rows, err := c.DBInstance.Query("SELECT * FROM blocked WHERE userid='" + strconv.Itoa(uid) + "'")

	//If an error ocurr, Print it out and return true.
	if err != nil {
		//Printing error and acting as blocked by returning true.
		log.Print(err)
		return true
	}

	//Close the rows when we are finished
	defer rows.Close()

	//Look at the rows we found in the database
	for rows.Next() {
		var id int     //The id of the found row
		var userid int //The user id bound to the found row

		//Scan the id and userid from the row
		rows.Scan(&id, &userid)
		log.Print("ID & Userid: ", id, userid) //Print the found values out to the log

		//Checking if the found id is same as the id for the specified user
		if uid == userid {
			//If true the user is actually blocked and we return true
			log.Print("User is blocked")
			return true
		}
	}

	//If the code reaches this then there was no blocked user found in the database
	log.Print("The user is not blocked ", username)
	return false
}

func IsUserAdmin() {

}

//Function to count the ammount of entries in the users table
func GetAmmountOfUsers(c *Crud) int {
	//Get all the rows from the users table. Not the most efficient SQL query but will do for this app.
	rows, err := c.DBInstance.Query("SELECT COUNT(*) FROM users")

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
		log.Printf("Successfully Found Users: %s", result)
		return result
	}

	return 0
}
