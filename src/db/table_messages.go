package db

type Message struct {
	id      int
	Userid  int
	Message string
}

func StoreMessage() {

}

func GetMessages() {

}

func DeleteMessages() {

}

//Returns the string used for database queries connected with this table
func GetMessagesTable() string {
	return "messages(userid, message)"
}
