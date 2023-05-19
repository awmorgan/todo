package main

import (
	"encoding/json"
	"os"
	"time"
)

type Task struct {
	Description string
	DueDate     time.Time
	Priority    int
	Completed   bool
}

func NewTask(description string, dueDate time.Time, priority int) *Task {
	return &Task{
		Description: description,
		DueDate:     dueDate,
		Priority:    priority,
		Completed:   false,
	}
}

type TodoList struct {
	Tasks []*Task
}

func NewTodoList() *TodoList {
	return &TodoList{
		Tasks: []*Task{},
	}
}

func (t *TodoList) AddTask(task *Task) {
	t.Tasks = append(t.Tasks, task)
}

func (t *TodoList) RemoveTask(taskNumber int) bool {
	if taskNumber <= 0 || taskNumber > len(t.Tasks) {
		return false
	}

	index := taskNumber - 1 // Adjust for 0-based index
	t.Tasks = append(t.Tasks[:index], t.Tasks[index+1:]...)
	return true
}

func (t *TodoList) CompleteTask(taskNumber int) bool {
	if taskNumber <= 0 || taskNumber > len(t.Tasks) {
		return false
	}

	t.Tasks[taskNumber-1].Completed = true
	return true
}

func (t *TodoList) PrioritizeTask(taskNumber, priority int) bool {
	if taskNumber <= 0 || taskNumber > len(t.Tasks) {
		return false
	}

	if priority <= 0 || priority > len(t.Tasks) {
		return false
	}

	task := t.Tasks[taskNumber-1]
	t.Tasks = append(t.Tasks[:taskNumber-1], t.Tasks[taskNumber:]...)
	t.Tasks = append(t.Tasks[:priority-1], append([]*Task{task}, t.Tasks[priority-1:]...)...)

	return true
}

func (t *TodoList) SaveTasks(filename string) error {
	// Convert the todo list to JSON
	bytes, err := json.Marshal(t.Tasks)
	if err != nil {
		return err
	}

	// Write the JSON to the file
	return os.WriteFile(filename, bytes, 0644)
}

func (t *TodoList) LoadTasks(filename string) error {
	// Open the file
	f, err := os.Open(filename)
	// Check for error
	if err != nil {
		return err
	}
	// defer file closing
	defer f.Close()

	// Create a slice to hold tasks
	var tasks []*Task
	// Decode the JSON from the file into tasks
	err = json.NewDecoder(f).Decode(&tasks)
	// Check for error
	if err != nil {
		return err
	}

	// Set t.Tasks to tasks
	t.Tasks = tasks

	// Return nil error
	return nil
}

func (t *TodoList) ClearTasks() {
	// Set t.Tasks to an empty slice
	t.Tasks = []*Task{}
}

func (t *TodoList) GetTaskList() []*Task {
	// Simply return t.Tasks
	return t.Tasks
}

func (t *TodoList) GetTaskCount() int {
	// Return the length of t.Tasks
	return len(t.Tasks)
}
