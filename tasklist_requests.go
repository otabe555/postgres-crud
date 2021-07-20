package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	model "github.com/otabe555/postgresql/model"
)

func AddTask(task *model.Task) int {
	db := CreateConnection()
	defer db.Close()
	id, sqlStatement := -1, `INSERT INTO tasklist (assignee, title, deadline, done) VALUES($1, $2, $3, $4) RETURNING id`
	err := db.QueryRow(sqlStatement, task.Assignee, task.Title, task.Deadline, task.Done).Scan(&id)
	if err != nil {
		log.Fatalf("Unable to execute the request %v", err)
	}
	fmt.Printf("Inserted a record %v", id)
	return id
}

func UpdateTask(id int, user *model.Task) int {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `UPDATE tasklist SET assignee=$1, title=$2, deadline=$3, done=$4 WHERE id=$5`
	res, err := db.Exec(sqlStatement, user.Assignee, user.Title, user.Deadline, user.Done, id)
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

func DeleteTask(id int) int {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `DELETE FROM tasklist WHERE id=$1`
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

func GetAllTasks() ([]model.Task, error) {
	db := CreateConnection()
	defer db.Close()
	var tasks []model.Task
	sqlStatement := `SELECT * FROM tasklist`
	rows, err := db.Query(sqlStatement)
	if err != nil {
		log.Fatal("Unable to execute the request ", err)
	}
	defer rows.Close()
	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.Id, &task.Assignee, &task.Title, &task.Deadline, &task.Done)
		if err != nil {
			log.Fatal("Unable to scan the row ", err)
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}

func GetTask(id int) (model.Task, error) {
	db := CreateConnection()
	defer db.Close()
	var task model.Task
	sqlStatement := `SELECT * FROM tasklist WHERE id=$1`
	row := db.QueryRow(sqlStatement, id)
	err := row.Scan(&task.Id, &task.Assignee, &task.Title, &task.Deadline, &task.Done)
	if err == sql.ErrNoRows {
		fmt.Print("err no rows were returned")
	} else if err != nil {
		log.Fatal("Unable to scan the row ", err)
	}
	return task, err
}

func MakeTaskDone(id int) int {
	db := CreateConnection()
	defer db.Close()
	sqlStatement := `UPDATE tasklist SET done=true WHERE id=$1`
	res, err := db.Exec(sqlStatement, id)
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
