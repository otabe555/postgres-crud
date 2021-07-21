package model

type Contact struct {
	ID int
	Name,
	Phone,
	Gender,
	Email,
	CreateAt string
}

type ContactManagerInterface interface {
	AddContact(user *Contact) error
	UpdateContact(id int, user *Contact) error
	DeleteContact(id int) error
	GetAllContacts() ([]Contact, error)
	GetContact(id int) (*Contact, error)
}

type Task struct {
	Id int
	Assignee,
	Title,
	Deadline string
	Done bool
}

type TaskManagerInterface interface {
	AddTask(task *Task) error
	UpdateTask(id int, task *Task) error
	DeleteTask(id int) error
	GetAllTasks() ([]Task, error)
	GetTask(id int) (*Task, error)
	MakeTaskDone(id int) error
}

const (
	Host     = "localhost"
	Port     = 5432
	User     = "madkingxxx"
	Password = "otabek123"
	Dbname   = "test"
)
