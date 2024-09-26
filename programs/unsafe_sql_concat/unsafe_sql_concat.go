package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {

	db, err := sql.Open("sqlite3", "./example.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var studentName string
	fmt.Println("Enter the student's name: ")
	fmt.Scan(&studentName)

	// Vulnerable SQL query construction using string concatenation
	query := "SELECT * FROM students WHERE name = '" + studentName + "';" // Vulnerability: Unsafe concatenation

	fmt.Println("Executing query:", query)

	rows, err := db.Query(query)
	if err != nil {
		log.Fatal("Error executing query:", err)
	}
	defer rows.Close()

	// Print the results
	for rows.Next() {
		var id int
		var name string
		var grade string
		if err := rows.Scan(&id, &name, &grade); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Grade: %s\n", id, name, grade)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
