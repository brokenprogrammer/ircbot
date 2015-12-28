//Database Migrations
package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func InitDB() {
	//Cleaning out the db before startup
	os.Remove("./db/ircbot.db")

	//Opening a new sql connection to our ircbot.db file using Sqlite
	db, err := sql.Open("sqlite3", "./db/ircbot.db")

	//Error opening DB
	if err != nil {
		//Log the error
		log.Fatal(err)
	}
	//Adding a defer statement so the connection closes when this function closes.(Needs Fix)
	defer db.Close()

	//Migrate tables up to the database
	migrationsUp(db)
	/*
		//Prepare statement
		stmt, err := db.Prepare("INSERT INTO tbl1(one, two) values(?,?)")
		if err != nil {
			log.Fatal(err)
		}

		//Execute the above statement
		stmt.Exec("HelloAg", "2")

		if err != nil {
			log.Fatal(err)
		}*/

}

//Function that runs migrations creating tables to the database
func migrationsUp(db *sql.DB) {
	//How to create tables
	sqlStmt := `create table if not exists foo (id integer not null primary key, name text);`
	_, err := db.Exec(sqlStmt) //executing above statement

	//error logging
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		//return
	}

	/*
		Users Table
		This table will remember all the users who type in the chat.
		They will be stored here and then we can use it to connect messages to them and block
		them from using the bots commands.
	*/
	sqlUserTable := `create table if not exists users (id INTEGER PRIMARY KEY AUTOINCREMENT, name TEXT NOT NULL);`
	//Execute the above string.
	_, usrErr := db.Exec(sqlUserTable)

	//Logging errors that occurr when adding the users table
	if usrErr != nil {
		log.Printf("%q: %s\n", usrErr, sqlUserTable)
	}

	/*
		Blocked Users Table
		This table will be used to store blocked users, if a user is blocked the bot will not
		listen to commands typed from this user.
	*/
	sqlBlockedTable := `create table if not exists blocked (id INTEGER PRIMARY KEY AUTOINCREMENT, userid INTEGER NOT NULL);`
	//Execute the above string.
	_, blcErr := db.Exec(sqlBlockedTable)

	//Logging errors that occurr when adding the block table
	if blcErr != nil {
		log.Printf("%q: %s\n", blcErr, sqlBlockedTable)
	}

	/*
		Messages Table
		This table will store all messages written by users so we later can use a command so the bot
		creates a new text file with all the typed messages in it as well as a timestamp.
	*/
	sqlMessageTable := `create table if not exists messages (id INTEGER PRIMARY KEY AUTOINCREMENT, userid INTEGER NOT NULL, message TEXT NOT NULL, time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL);`
	//Execute the above string.
	_, msgErr := db.Exec(sqlMessageTable)

	//Logging errors that occurr when adding the message table
	if msgErr != nil {
		log.Printf("%q: %s\n", msgErr, sqlMessageTable)
	}

	//Finished with migrating to database
	log.Print("Databases Migrated Successfully. \n")
}

//Function that runs the migrations removing tables from the database
func migrationsDown(db *sql.DB) {

}

//Function to add some "Dummy" data to the database
func databaseSeed(db *sql.DB) {

}
