//Database Migrations
package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	//"os"
)

func InitDB() {
	//os.Remove("./ircbot.db")

	//Opening a new sql connection to our ircbot.db file using Sqlite
	db, err := sql.Open("sqlite3", "./db/ircbot.db")

	//Error opening DB
	if err != nil {
		//Log the error
		log.Fatal(err)
	}
	//Adding a defer statement so the connection closes when this function closes.(Needs Fix)
	defer db.Close()

	//How to create tables
	sqlStmt := `create table if not exists foo (id integer not null primary key, name text);`
	_, err = db.Exec(sqlStmt) //executing above statement

	//error logging
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
		//return
	}

	//Prepare statement
	stmt, err := db.Prepare("INSERT INTO tbl1(one, two) values(?,?)")
	if err != nil {
		log.Fatal(err)
	}

	//Execute the above statement
	stmt.Exec("HelloAg", "2")

	if err != nil {
		log.Fatal(err)
	}

}
