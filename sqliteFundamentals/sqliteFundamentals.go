package sqlitefundamentals

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Book struct {
	id     int
	name   string
	author string
}

func Run() {
	db, err := sql.Open("sqlite3", "./books.db")
	log.Println(db)
	if err != nil {
		log.Println(err)
	}

	//create table
	statement, err := db.Prepare("CREATE TABLE IF NOT EXISTS books (id INTEGER PRIMARY KEY, isbn INTEGER UNIQUE, author VARCHAR(64), name VARCHAR(64) null)")
	if err != nil {
		log.Printf("Error creating table : %v", err)
	} else {
		log.Println("Succesfully created table books!")
	}

	statement.Exec()

	//create
	statement, _ = db.Prepare("INSERT OR IGNORE INTO books (name , author , isbn) VALUES (?,?,?)")
	statement.Exec("A Tale of Two Cities", "Charles Dickens", 140430547)
	log.Println("Inserted the book into the database!")

	//read
	rows, _ := db.Query("SELECT id, name, author FROM books")
	var tempBook Book

	// rows.Next returns true or false if there is a next value in the table
	for rows.Next() {
		//rows.Scan takes what is the rows and puts it in the pointers provided
		rows.Scan(&tempBook.id, &tempBook.name, &tempBook.author)
		log.Printf("ID:%d, Book:%s, Author:%s\n", tempBook.id, tempBook.name, tempBook.author)
	}

	//update
	statement, _ = db.Prepare("update books set name=? where id=?")
	statement.Exec("The Tale of Two Cities", 1)
	log.Println("Successfully updated the book in database!")

	//delete
	statement, _ = db.Prepare("delete from books where id=?")
	statement.Exec(1)
	log.Println("Successfully deleted the book in database!")
}
