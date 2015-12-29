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
	irtString = "INSERT INTO " + table + "VALUES(" + stringBuilder(len(values)) + ")"

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

func (c *Crud) Delete(table string, id int) error {
	//Initiate new transaction.
	tx, err := c.DBInstance.Begin()

	//Error logging for starting transaction.
	if err != nil {
		log.Print(err)
		return err
	}

	var delString string
	delString = `DELETE FROM ` + table + ` WHERE id= ?`

	//Initialize delete statement.
	delete, err := tx.Prepare(delString)

	//Error logging for the statement initialization.
	if err != nil {
		log.Print(err)
		return err
	}

	defer delete.Close()

	//Execute the statement pushing values to database
	_, err = delete.Exec(id)

	if err != nil {
		log.Print(err)
		return err
	}

	//Commit executed statements to the database
	tx.Commit()

	//Log successmessage and return no errors
	log.Printf("Successfully Deleted From Database: %s, %d\n", table, id)
	return nil
}

func (c *Crud) Select(table string) error {
	//Execute a query that returns rows
	rows, err := c.DBInstance.Query(`SELECT id, name FROM ` + table)

	//Error executing query, log error and return
	if err != nil {
		log.Print(err)
		return err
	}

	//Close the rows when function is finished
	defer rows.Close()

	//Loop through the rows and print them out to the console.
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		log.Print(id, name)
	}

	//Log successmessage and return no errors
	log.Printf("Successfully Selected From Database: %s \n", table)
	return nil
}

//Function used to build a string depending on a length of values for the database queries
func stringBuilder(length int) string {
	var str string
	str = "?"

	for i := 1; i < length; i++ {
		str += ", ?"
	}

	//returning a string like "?, ?, ?" depending on how many values needed in the prepared statement
	return str
}
