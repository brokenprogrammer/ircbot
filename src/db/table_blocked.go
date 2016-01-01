package db

import (
	"log"
	"strconv"
)

type Blocked struct {
	id     int
	Userid int
}

//Function to unblock an user from using the commands.
func BlockUser(user string, c *Crud) {
	//Get the userid of the specified user by checking GetUserByName
	userid := GetUserByName(user, c)

	//If the user id is 0 its the same as it doesnt exist.
	if userid != 0 {
		//If everything goes as planned we insert the id to the database
		c.Insert(GetBlockedTable(), strconv.Itoa(userid))
	}
	//Print out an success message
	log.Print("Successfully blocked user: ", user)
}

//Function to unblock user letting them use the commands again
func UnBlockUser(user string, c *Crud) {
	//Get the userid of the specified user by checking GetUserByName
	userid := GetUserByName(user, c)

	//If the user id is 0 then it doesnt exist
	if userid != 0 {
		//If everything goes as planned we delete the record of the user here.
		c.Delete(GetBlockedTableRaw(), userid)
	}

	//Print out an success message
	log.Print("Successfully unblocked user: ", user)
}

//Returns the string used for database queries connected with this table
func GetBlockedTable() string {
	return "blocked(userid)"
}

//Returns raw version of the blocked table
func GetBlockedTableRaw() string {
	return "blocked"
}
