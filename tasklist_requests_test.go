package main

import (
	"testing"

	"github.com/otabe555/postgresql/model"
)

func TestAddTask(t *testing.T) {
	if AddTask(&model.Task{
		Assignee: "Kim Kim",
		Title:    "Do the homework",
		Deadline: "2021-07-21",
		Done:     false,
	}) == -1 {
		t.Error("Failed to Add user")
	}
}

func TestUpdateTask(t *testing.T) {
	if UpdateTask(1, &model.Task{
		Assignee: "Kim Kim",
		Title:    "Do the homework",
		Deadline: "2021-07-21",
		Done:     true,
	}) != 1 {
		t.Error("Failed to Update task")
	}
}

func TestDeleteTask(t *testing.T) {
	if DeleteTask(5) != 1 {
		t.Error("Failed to Delete the task")
	}
}

func TestGetTask(t *testing.T) {
	_, err := GetTask(1)
	if err != nil {
		t.Error("Failed to Get task", err)
	}
}
func TestGetAllTasks(t *testing.T) {
	_, err := GetAllTasks()
	if err != nil {
		t.Error("Failed to GetAll task list ", err)
	}
}
func TestMakeTaskDone(t *testing.T) {
	if MakeTaskDone(1) != 1 {
		t.Error("Failed to complete this task ")
	}
}
