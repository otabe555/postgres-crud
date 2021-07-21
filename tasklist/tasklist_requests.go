package tasklist

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	model "github.com/otabe555/postgresql/model"
)

type TaskManager struct {
	db *sql.DB
}

func NewTaskManager() (model.TaskManagerInterface, error) {
	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable",
		model.User, model.Dbname, model.Password, model.Host, model.Port)
	db, err := sql.Open("postgres", connection)
	tm := &TaskManager{db}
	if err != nil {
		return nil, err
	}
	err = tm.db.Ping()
	if err != nil {
		return nil, err
	}
	return tm, nil
}

func (tm TaskManager) AddTask(task *model.Task) error {
	sqlStatement := `INSERT INTO tasklist (assignee, title, deadline, done) VALUES($1, $2, $3, $4) RETURNING id`
	err := tm.db.QueryRow(sqlStatement, task.Assignee, task.Title, task.Deadline, task.Done).Scan(&task.Id)
	if err != nil {
		return err
	}
	return nil
}

func (tm TaskManager) UpdateTask(id int, user *model.Task) error {
	sqlStatement := `UPDATE tasklist SET assignee=$1, title=$2, deadline=$3, done=$4 WHERE id=$5`
	res, err := tm.db.Exec(sqlStatement, user.Assignee, user.Title, user.Deadline, user.Done, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (tm TaskManager) DeleteTask(id int) error {
	sqlStatement := `DELETE FROM tasklist WHERE id=$1`
	res, err := tm.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (tm TaskManager) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	sqlStatement := `SELECT * FROM tasklist`
	rows, err := tm.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var task model.Task
		err = rows.Scan(&task.Id, &task.Assignee, &task.Title, &task.Deadline, &task.Done)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, err
}

func (tm TaskManager) GetTask(id int) (*model.Task, error) {
	var task model.Task
	sqlStatement := `SELECT * FROM tasklist WHERE id=$1`
	row := tm.db.QueryRow(sqlStatement, id)
	err := row.Scan(&task.Id, &task.Assignee, &task.Title, &task.Deadline, &task.Done)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &task, err
}

func (tm TaskManager) MakeTaskDone(id int) error {
	defer tm.db.Close()
	sqlStatement := `UPDATE tasklist SET done=true WHERE id=$1`
	res, err := tm.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}
