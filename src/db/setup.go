//Database Migrations
package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

func InitDB(dbdriver string, dbpath string) {
	//Cleaning out the db before startup
	os.Remove(dbpath)

	//Opening a new sql connection to our ircbot.db file using Sqlite
	db, err := sql.Open(dbdriver, dbpath)

	//Error opening DB
	if err != nil {
		//Log the error
		log.Fatal(err)
	}
	//Adding a defer statement so the connection closes when this function closes.(Needs Fix)
	defer db.Close()

	//Migrate tables up to the database
	migrationsUp(db)

	//Removing the tables from the database(Migrating down)
	//migrationsDown(db)

	//Seeding database with dummy data
	databaseSeed(db)
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
	//How to drop tables
	sqlStmt := `DROP TABLE foo;`
	_, err := db.Exec(sqlStmt) //executing above statement

	//error logging
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		//return
	}

	//Dropping the users table
	sqlDropUsers := `DROP TABLE users;`
	_, usrErr := db.Exec(sqlDropUsers)

	//If errors occurr when dropping the users table
	if usrErr != nil {
		log.Printf("%q: %s\n", usrErr, sqlDropUsers)
	}

	//Dropping the blocked table
	sqlDropBlocked := `DROP TABLE blocked;`
	_, blcErr := db.Exec(sqlDropBlocked)

	//If errors occurr when dropping the blocked table
	if blcErr != nil {
		log.Printf("%q: %s\n", blcErr, sqlDropBlocked)
	}

	//Dropping the messages table
	sqlDropMessages := `DROP TABLE messages;`
	_, msgErr := db.Exec(sqlDropMessages)

	//If errors occurr when dropping the messages table
	if msgErr != nil {
		log.Printf("%q: %s\n", msgErr, sqlDropMessages)
	}

	//Finished with dropping database tables
	log.Print("Databases Dropped Successfully. \n")
}

//Function to add some "Dummy" data to the database
func databaseSeed(db *sql.DB) {
	/*
		Tx is a in progress database transaction, it is what db.Beigin returns.
		A transaction must end with tx.Commit() after that all operations will ressult in
		an error (ErrTxDone) meaning the transaction is already finished.

		Starts a new transaction to the database
	*/
	tx, err := db.Begin()

	//Prepare statement for the users table
	users, err := tx.Prepare("INSERT INTO users(name) VALUES(?)")

	//Error preparing statement, log the errors
	if err != nil {
		log.Fatal(err)
	}

	//Close the statement when this functions is done
	defer users.Close()

	//Seeding the users table with users
	_, err = users.Exec("John")
	_, err = users.Exec("BadBob")
	_, err = users.Exec("GoBotOwner")

	//Prepare statement for the blocked table
	blocked, err := tx.Prepare("INSERT INTO blocked(userid) VALUES(?)")

	//Error preparing statement, log the errors
	if err != nil {
		log.Fatal(err)
	}

	//Close the statement when this function is done
	defer blocked.Close()

	//Add user id 2 to the blocked list
	_, err = blocked.Exec(2)

	//Prepare statement for the messages table
	messages, err := tx.Prepare("INSERT INTO messages(userid, message) VALUES(?, ?)")

	//Error preparing statement, log the errors
	if err != nil {
		log.Fatal(err)
	}

	//Close the statement when this function is done
	defer messages.Close()

	//Seeding the messages table with messages
	_, err = messages.Exec(1, "Hello World")
	_, err = messages.Exec(2, "Bad Message")
	_, err = messages.Exec(3, "I the owner of this bot")

	//Commit the transaction to the database
	tx.Commit()

	//Log success message to show that seeding is finished
	log.Print("Database Seeded Successfully.")
}
