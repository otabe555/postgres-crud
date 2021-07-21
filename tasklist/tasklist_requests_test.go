package tasklist

import (
	"testing"

	"github.com/otabe555/postgresql/model"
)

var (
	tm         model.TaskManagerInterface
	id         int
	user       = &model.Task{Assignee: "Kim Kim", Title: "Do the homework", Deadline: "2021-07-21", Done: false}
	userforupd = &model.Task{Assignee: "Kim Kim", Title: "Do the homework", Deadline: "2021-07-21", Done: true}
)

func TestNewTaskManager(t *testing.T) {
	var err error
	tm, err = NewTaskManager()
	if err != nil {
		t.Error("Failed to connect to the db")
	}
}
func TestTaskManager_AddTask(t *testing.T) {
	if tm.AddTask(user) != nil {
		t.Error("Failed to Add user")
	}
	id = user.Id
}

func TestTaskManager_UpdateTask(t *testing.T) {
	if tm.UpdateTask(1, userforupd) != nil {
		t.Error("Failed to Update task")
	}
}

func TestTaskManager_DeleteTask(t *testing.T) {
	TestTaskManager_AddTask(t)
	if tm.DeleteTask(id) != nil {
		t.Error("Failed to Delete the task")
	}
}

func TestTaskManager_GetTask(t *testing.T) {
	_, err := tm.GetTask(1)
	if err != nil {
		t.Error("Failed to Get task", err)
	}
}
func TestTaskManager_GetAllTasks(t *testing.T) {
	_, err := tm.GetAllTasks()
	if err != nil {
		t.Error("Failed to GetAll task list ", err)
	}
}
func TestTaskManager_MakeTaskDone(t *testing.T) {
	if tm.MakeTaskDone(1) != nil {
		t.Error("Failed to complete this task ")
	}
}
