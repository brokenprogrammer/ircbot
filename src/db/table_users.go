package db

type User struct {
	id   int
	Name string
}

//Returns the string used for database queries connected with this table
func GetUsersTable() string {
	return "users(name)"
}

//Function used to insert a new user into the database
func NewUser() {
	//TODO: Check if user exists, If not then insert the user into the database
	//TODO: Apply this function to the chat reader.
}
