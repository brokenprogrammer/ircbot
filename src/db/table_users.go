package db

type User struct {
	id   int
	Name string
}

//Returns the string used for database queries connected with this table
func GetUsersTable() string {
	return "users(name)"
}
