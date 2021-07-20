package model

type Contact struct {
	ID int
	Name,
	Phone,
	Gender,
	Email,
	CreateAt string
}

type Task struct {
	Id int
	Assignee,
	Title,
	Deadline string
	Done bool
}
