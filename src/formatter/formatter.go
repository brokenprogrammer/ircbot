package formatter

import (
	"strings"
)

//Function used to build a string depending on a length of values for the database queries
func StringBuilder(length int) string {
	var str string
	str = "?"

	for i := 1; i < length; i++ {
		str += ", ?"
	}

	//returning a string like "?, ?, ?" depending on how many values needed in the prepared statement
	return str
}

//Function used to format out the username out of input from the irc server
func GetUsername(s string) string {
	//Declare variable to hold the username
	var username string
	//All usernames is between : and ! so get the index of the ! mark
	usernameEnd := strings.Index(s, "!")

	//Get the username by getting the string between the colon and ! mark
	username = s[1:usernameEnd]

	return username
}
