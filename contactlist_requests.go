package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	model "github.com/otabe555/postgresql/model"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "madkingxxx"
	password = "otabek123"
	dbname   = "test"
)

func CreateConnection() *sql.DB {
	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable", user, dbname, password, host, port)
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func AddContact(user *model.Contact) int {
	db := CreateConnection()
	defer db.Close()
	id, sqlStatement := -1, `INSERT INTO contactlist (name, phone, gender, email) VALUES($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(sqlStatement, user.Name, user.Phone, user.Gender, user.Email).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the request %v", err)
	}
	fmt.Printf("Inserted a record %v", id)
	return id
}

func UpdateContact(id int, user *model.Contact) int {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `UPDATE contactlist SET name=$1, phone=$2, gender=$3, email=$4 WHERE id=$5`
	res, err := db.Exec(sqlStatement, user.Name, user.Phone, user.Gender, user.Email, id)
	if err != nil {
		log.Fatal("Unable to execute the requst ", err)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatalf("Error while checking the affected rows %v\n", err)
	}
	fmt.Printf("Total count of affected rows %v\n", rowsAffected)
	return int(rowsAffected)
}

func DeleteContact(id int) int {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM contactlist WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)
	if err != nil {
		log.Fatal("Unable to execute the query")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Fatal("Error while checking the affected rows", err)
	}
	fmt.Print("The total rows were affected are ", rowsAffected)
	return int(rowsAffected)
}

func GetAllContacts() ([]model.Contact, error) {
	db := CreateConnection()
	defer db.Close()
	var users []model.Contact
	sqlStatement := `SELECT * FROM contactlist`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal("Unable to execute the request ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user model.Contact
		err = rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Gender, &user.Email, &user.CreateAt)
		if err != nil {
			log.Fatal("Unable to scan the row ", err)
		}
		users = append(users, user)
	}
	return users, err
}

func GetContact(id int) (model.Contact, error) {
	db := CreateConnection()
	defer db.Close()
	var user model.Contact
	sqlStatement := `SELECT * FROM contactlist WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&user.ID, &user.Name, &user.Phone, &user.Gender, &user.Email, &user.CreateAt)
	if err == sql.ErrNoRows {
		fmt.Print("err no rows were returned")
	} else if err != nil {
		log.Fatal("Unable to scan the row ", err)
	}
	return user, err
}
