/*
	db_Crud is used to store our Crud structure which will act as our Database wrapper.
	Here methods for our Database connection is store which we easily can refer to when we want
	to do basic tasks such as: Insert, Delete, Update, Select. And even more specific tasks such as
	findById().

	In our Crud structure we store the Driver used for the database as well as the path for the Database
	together with the actual instance.
*/
package db

import (
	"database/sql"
	"formatter"
	"log"
)

//Our Crud structure which will be used for all the tasks when handling our database.
type Crud struct {
	DBDriver   string  //The database driver used.
	DBPath     string  //The path for our sqlite3 database.
	DBInstance *sql.DB //The actual instance for our database connection
}

//Sets up a new Crud instance, This is what we will use for our db connection.
func NewCrud(driver string, path string) *Crud {
	//db, err := sql.Open("sqlite3", "./db/ircbot.db")
	//Create a new database connection. The DB will we used in our structure as the instance.
	db, err := sql.Open(driver, path)

	//If there is an error when connecting to the database.
	if err != nil {
		//Log errors.
		log.Print(err)
	}

	//returns a new Crud instance we can use to access our Database.
	return &Crud{
		DBDriver:   driver,
		DBPath:     path,
		DBInstance: db,
	}
}

//Insert function for inserting data into specified table. Takes in a table string and many values
func (c *Crud) Insert(table string, values ...string) error {
	//Initiate new transaction.
	tx, err := c.DBInstance.Begin()

	//Error logging for starting transaction.
	if err != nil {
		log.Print(err)
		return err
	}

	//Building the insert string depending on how many values in the parameters
	var irtString string
	irtString = "INSERT INTO " + table + "VALUES(" + formatter.StringBuilder(len(values)) + ")"

	//Initialize insert statement.
	insert, err := tx.Prepare(irtString)

	//Error logging for the statement initialization.
	if err != nil {
		log.Print(err)
		return err
	}

	//Defer closing for the statement.
	defer insert.Close()

	//Execute the statement pushing values to database
	_, err = insert.Exec(values[0])

	if err != nil {
		log.Print(err)
		return err
	}

	//Commit executed statements to the database
	tx.Commit()

	//Log successmessage and return no errors
	log.Printf("Successfully Inserted Into Database: %v\n", values)
	return nil
}

func (c *Crud) insertMessage(table string, userid int, message string) error {
	//Initiate new transaction.
	tx, err := c.DBInstance.Begin()

	//Error logging for starting transaction.
	if err != nil {
		log.Print(err)
		return err
	}

	//Building the insert string depending on how many values in the parameters
	var irtString string
	irtString = "INSERT INTO " + table + "VALUES(" + "?, ?" + ")"

	//Initialize insert statement.
	insert, err := tx.Prepare(irtString)

	//Error logging for the statement initialization.
	if err != nil {
		log.Print(err)
		return err
	}

	//Defer closing for the statement.
	defer insert.Close()

	//Execute the statement pushing values to database
	_, err = insert.Exec(userid, message)

	if err != nil {
		log.Print(err)
		return err
	}

	//Commit executed statements to the database
	tx.Commit()

	//Log successmessage and return no errors
	log.Printf("Successfully Inserted Into Database: %v\n", message)
	return nil
}

//Update function for updating data into specified table.
//Takes in a table string, column to update, what to update it to and id of the target.
//Example call to this function Update("users", "name", "NewName", 4)
func (c *Crud) Update(table string, column string, newVal string, id int) error {
	//Initiate new transaction.
	tx, err := c.DBInstance.Begin()

	//Error logging for starting transaction.
	if err != nil {
		log.Print(err)
		return err
	}

	var updString string
	updString = `UPDATE ` + table + ` SET ` + column + `="` + newVal + `" WHERE id= ?`

	//Initialize update statement.
	update, err := tx.Prepare(updString)

	//Error logging for the statement initialization.
	if err != nil {
		log.Print(err)
		return err
	}

	defer update.Close()

	//Execute the statement pushing values to database
	_, err = update.Exec(id)

	if err != nil {
		log.Print(err)
		return err
	}

	//Commit executed statements to the database
	tx.Commit()

	//Log successmessage and return no errors
	log.Printf("Successfully Updated Database: %v,%v\n", column, newVal)
	return nil
}

func (c *Crud) Delete(table string, where string, id int) error {
	//Initiate new transaction.
	tx, err := c.DBInstance.Begin()

	//Error logging for starting transaction.
	if err != nil {
		log.Print(err)
		return err
	}

	var delString string
	delString = `DELETE FROM ` + table + ` WHERE ` + where + `=?`
	/*if table == GetUsersTable() {
		delString = `DELETE FROM ` + table + ` WHERE ` + where + `=?`
	} else {
		delString = `DELETE FROM ` + table + ` WHERE userid=?`
		log.Print("Table is using userid")
	}*/

	//Initialize delete statement.
	delete, err := tx.Prepare(delString)

	//Error logging for the statement initialization.
	if err != nil {
		log.Print("Error preparing: ", err)
		return err
	}

	defer delete.Close()

	//Execute the statement pushing values to database
	_, err = delete.Exec(id)

	if err != nil {
		log.Print("Error executing: ", err)
		return err
	}

	//Commit executed statements to the database
	tx.Commit()

	//Log successmessage and return no errors
	log.Printf("Successfully Deleted From Database: %s, %d\n", table, id)
	return nil
}

//TODO: Create a DB wrapper
//TODO: https://github.com/joshcam/PHP-MySQLi-Database-Class
func (c *Crud) Select(table string) (map[int]string, error) {
	results := make(map[int]string)
	var query string

	//Execute a query that returns rows
	query = `SELECT ` + `*` + ` FROM ` + table

	rows, err := c.DBInstance.Query(query)

	//Error executing query, log error and return
	if err != nil {
		log.Print(err)
		return results, err
	}

	//Close the rows when function is finished
	defer rows.Close()

	//Loop through the rows and print them out to the console.
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		log.Print(id, name)
		results[id] = name
	}

	//Log successmessage and return no errors
	log.Printf("Successfully Selected From Database: %s \n", table)
	return results, nil
}

func (c *Crud) SelectSpecific(table string, column string, value string) (int, string) {
	var query string

	query = `SELECT * FROM ` + table + ` WHERE ` + column + `='` + value + `'`

	rows, err := c.DBInstance.Query(query)

	if err != nil {
		log.Print(err)
		return 0, ""
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		log.Print(id, name)
		return id, name
	}

	log.Print("End of func")
	return 0, ""
}

func (c *Crud) getMessages(table string, user string) string {
    var query string
    var messages string

    query = `SELECT * FROM ` + table + ` WHERE userid='` + user + `'`
    log.Println(query)
    rows, err := c.DBInstance.Query(query)

	if err != nil {
		log.Print(err)
		return ""
	}

	defer rows.Close()
    log.Println("Starting getting entries: ")
	for rows.Next() {
		var id int
		var userid int
        var message string
        var time string
		rows.Scan(&id, &userid, &message, &time)
        messages += message + "\n"
		log.Println(id, userid, message, time)
	}
    log.Println(messages)
	log.Print("End of func")
    return messages
}

func findByID() {

}

func findByName() {

}

func findAll() {

}
