package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	id        int
	username  string
	password  string
	createdAt time.Time
}

var lastUserId int = 0

func createTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS users (
		id INT AUTO_INCREMENT,
		username TEXT NOT NULL,
		password TEXT NOT NULL,
		created_at DATETIME,
		PRIMARY KEY(id)
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("The users table has been created successfully")
	}
}

func insertUser(db *sql.DB) {
	query := `INSERT INTO users (username, password, created_at) VALUES (?, ?, ?)`

	username := "adam"
	password := "pass"
	createdAt := time.Now()

	res, err := db.Exec(query, username, password, createdAt)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("A new record has been inserted into users table")
	}

	userId, err := res.LastInsertId()
	log.Println("New record's ID:", userId)
}

func getUserById(db *sql.DB, userId int) {
	var (
		id        int
		username  string
		password  string
		createdAt time.Time
	)

	query := `SELECT * FROM users WHERE id = ?`
	err := db.QueryRow(query, 1).Scan(&id, &username, &password, &createdAt)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("Successfully read fields for id:", id)

		// Display the fields
		fmt.Println("User ID:\t", id)
		fmt.Println("Username:\t", username)
		fmt.Println("Password:\t", password)
		fmt.Println("The date:\t", createdAt)
	}
}

func getAllUsers(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM users")
	defer rows.Close()

	var users []user
	if err != nil {
		log.Println(err)
	} else {
		for rows.Next() {
			var u user
			err := rows.Scan(&u.id, &u.username, &u.password, &u.createdAt)
			if err != nil {
				log.Println("Error getting user record")
			} else {
				users = append(users, u)
			}
		}
	}

	fmt.Println("User records:")
	for index, user := range users {
		fmt.Println("Index:", index)
		fmt.Println("User ID :", user.id)
		fmt.Println("Username:", user.username)
		fmt.Println("Password:", user.password)
		fmt.Println("The Date:", user.createdAt)
		fmt.Println()
		lastUserId = user.id
	}
}

func deleteLastUser(db *sql.DB) {
	if lastUserId == 0 {
		getAllUsers(db)
	}

	log.Println("Deleting user id:", lastUserId)
	_, err := db.Exec("DELETE FROM users WHERE id = ?", lastUserId)
	if err != nil {
		log.Println(err)
	} else {
		log.Println("User successfully deleted")
	}
}

func main() {
	// Open a connection to database
	db, err := sql.Open("mysql", "main:main@(127.0.0.1:3306)/main?parseTime=true")
	if err != nil {
		panic(fmt.Sprintf("Error connecting to database: %s", err))
	} else {
		fmt.Println("Database connection is successful")
	}

	fmt.Println(db.Ping())
	createTable(db)
	insertUser(db)
	getUserById(db, 1)
	getAllUsers(db)
	deleteLastUser(db)
}
