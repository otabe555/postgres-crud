package contactlist

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	model "github.com/otabe555/postgresql/model"
)

type ContactManager struct {
	db *sql.DB
}

func NewContactManager() (model.ContactManagerInterface, error) {
	connection := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable",
		model.User, model.Dbname, model.Password, model.Host, model.Port)
	db, err := sql.Open("postgres", connection)
	cm := &ContactManager{db}
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return cm, nil
}

func (cm ContactManager) AddContact(user *model.Contact) error {
	sqlStatement := `INSERT INTO contactlist (name, phone, gender, email) VALUES($1, $2, $3, $4) RETURNING id`
	err := cm.db.QueryRow(sqlStatement, user.Name, user.Phone, user.Gender, user.Email).Scan(&user.ID)
	if err != nil {
		return err
	}
	return nil
}

func (cm ContactManager) UpdateContact(id int, user *model.Contact) error {
	sqlStatement := `UPDATE contactlist SET name=$1, phone=$2, gender=$3, email=$4 WHERE id=$5`
	res, err := cm.db.Exec(sqlStatement, user.Name, user.Phone, user.Gender, user.Email, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (cm ContactManager) DeleteContact(id int) error {
	sqlStatement := `DELETE FROM contactlist WHERE id=$1`
	res, err := cm.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (cm ContactManager) GetAllContacts() ([]model.Contact, error) {
	var users []model.Contact
	sqlStatement := `SELECT * FROM contactlist`
	rows, err := cm.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var user model.Contact
		err = rows.Scan(&user.ID, &user.Name, &user.Phone, &user.Gender, &user.Email, &user.CreateAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, err
}

func (cm ContactManager) GetContact(id int) (*model.Contact, error) {
	var user model.Contact
	sqlStatement := `SELECT * FROM contactlist WHERE id=$1`
	row := cm.db.QueryRow(sqlStatement, id)
	err := row.Scan(&user.ID, &user.Name, &user.Phone, &user.Gender, &user.Email, &user.CreateAt)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &user, err
}
