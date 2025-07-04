package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

func main() {
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer db.Close()

	var version string
	err = db.QueryRow("SELECT SQLITE_VERSION()").Scan(&version)
	if err != nil {
		fmt.Println("Version:", err)
		return
	}

	fmt.Println("SQLite3 version:", version)
	os.Remove("test.db")
}
