package db

type Blocked struct {
	id     int
	Userid int
}

func BlockUser() {

}

func UnBlockUser() {

}

//Returns the string used for database queries connected with this table
func GetBlockedTable() string {
	return "blocked(userid)"
}
