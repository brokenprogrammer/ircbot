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

	//Commit executed statements to the database
	tx.Commit()

	//Log successmessage and return no errors
	log.Printf("Successfully Inserted Into Database: %v\n", values)
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
