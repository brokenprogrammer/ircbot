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

//function that extracts a message out of the info the irc server sends out
func ExtractMessage(str string) string {
	//New string that will hold the ressult of this function
	var newStr string

	//Split the given string so we can work with the length of different parts of the string.
	splitted := strings.Split(str, " ")

	//Integer to hold the length of the parts of the message we want to remove.
	var extraLen int

	//Looping through the parts we want to remove (Splitted 0-2)
	for i := 0; i <= 2; i++ {
		extraLen += len(splitted[i])
	}

	//Adding for spaces and the colon infront of all messages
	extraLen += 4

	//Making the new string everything after the part we want to remove
	newStr = str[extraLen:]

	//returning the new string
	return newStr
}
