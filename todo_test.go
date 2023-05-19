package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"
)

func TestNewTask(t *testing.T) {
	description := "Task 1"
	dueDate := time.Now()
	priority := 1

	task := NewTask(description, dueDate, priority)

	if task.Description != description {
		t.Errorf("Expected task description to be %s, got %s", description, task.Description)
	}

	if !task.DueDate.Equal(dueDate) {
		t.Errorf("Expected due date to be %v, got %v", dueDate, task.DueDate)
	}

	if task.Priority != priority {
		t.Errorf("Expected priority to be %d, got %d", priority, task.Priority)
	}

	if task.Completed {
		t.Error("Expected task to be uncompleted on creation")
	}
}

func TestNewTodoList(t *testing.T) {
	todoList := NewTodoList()

	if todoList == nil {
		t.Error("Expected new TodoList to be non-nil")
	}

	if len(todoList.Tasks) != 0 {
		t.Errorf("Expected new TodoList to have no tasks, found %d", len(todoList.Tasks))
	}
}

func TestAddTask(t *testing.T) {
	todoList := NewTodoList()
	task := NewTask("Task 1", time.Now(), 1)

	todoList.AddTask(task)

	if len(todoList.Tasks) != 1 {
		t.Errorf("Expected TodoList to have 1 task, found %d", len(todoList.Tasks))
	}

	if todoList.Tasks[0] != task {
		t.Error("Expected the added task to be the same as the task in the TodoList")
	}
}

func TestRemoveTask(t *testing.T) {
	todoList := NewTodoList()
	task1 := NewTask("Task 1", time.Now(), 1)
	task2 := NewTask("Task 2", time.Now(), 2)

	todoList.AddTask(task1)
	todoList.AddTask(task2)

	removed := todoList.RemoveTask(1)

	if !removed {
		t.Error("Expected RemoveTask to return true, got false")
	}

	if len(todoList.Tasks) != 1 {
		t.Errorf("Expected TodoList to have 1 task, found %d", len(todoList.Tasks))
	}

	if todoList.Tasks[0] != task2 {
		t.Error("Expected the remaining task to be Task 2")
	}

	// Attempt to remove a non-existent task
	removed = todoList.RemoveTask(10)
	if removed {
		t.Error("Expected RemoveTask to return false for non-existent task, got true")
	}

	// Attempt to remove a task from an empty list
	emptyTodoList := NewTodoList()
	removed = emptyTodoList.RemoveTask(1)
	if removed {
		t.Error("Expected RemoveTask to return false for empty list, got true")
	}
}

func TestCompleteTask(t *testing.T) {
	todoList := NewTodoList()
	task1 := NewTask("Task 1", time.Now(), 1)
	task2 := NewTask("Task 2", time.Now(), 2)

	todoList.AddTask(task1)
	todoList.AddTask(task2)

	completed := todoList.CompleteTask(1)

	if !completed {
		t.Error("Expected CompleteTask to return true, got false")
	}

	if !todoList.Tasks[0].Completed {
		t.Error("Expected the first task to be completed")
	}

	// Attempt to complete a non-existent task
	completed = todoList.CompleteTask(10)
	if completed {
		t.Error("Expected CompleteTask to return false for non-existent task, got true")
	}

	// Attempt to complete a task from an empty list
	emptyTodoList := NewTodoList()
	completed = emptyTodoList.CompleteTask(1)
	if completed {
		t.Error("Expected CompleteTask to return false for empty list, got true")
	}
}

func TestPrioritizeTask(t *testing.T) {
	todoList := NewTodoList()
	task1 := NewTask("Task 1", time.Now(), 1)
	task2 := NewTask("Task 2", time.Now(), 2)

	todoList.AddTask(task1)
	todoList.AddTask(task2)

	prioritized := todoList.PrioritizeTask(2, 1)

	if !prioritized {
		t.Error("Expected PrioritizeTask to return true, got false")
	}

	if todoList.Tasks[0] != task2 || todoList.Tasks[1] != task1 {
		t.Error("Expected the tasks to be reordered")
	}

	// Attempt to prioritize a non-existent task
	prioritized = todoList.PrioritizeTask(10, 1)
	if prioritized {
		t.Error("Expected PrioritizeTask to return false for non-existent task, got true")
	}

	// Attempt to prioritize a task to a non-existent position
	prioritized = todoList.PrioritizeTask(1, 10)
	if prioritized {
		t.Error("Expected PrioritizeTask to return false for non-existent position, got true")
	}

	// Attempt to prioritize a task in an empty list
	emptyTodoList := NewTodoList()
	prioritized = emptyTodoList.PrioritizeTask(1, 1)
	if prioritized {
		t.Error("Expected PrioritizeTask to return false for empty list, got true")
	}
}

func TestSaveTasks(t *testing.T) {
	todoList := NewTodoList()
	task1 := NewTask("Task 1", time.Now(), 1)
	task2 := NewTask("Task 2", time.Now(), 2)

	todoList.AddTask(task1)
	todoList.AddTask(task2)

	err := todoList.SaveTasks("/tmp/todoList.json")

	if err != nil {
		t.Errorf("Expected SaveTasks to complete without error, got: %v", err)
	}

	// Load the file and check its contents
	bytes, err := os.ReadFile("/tmp/todoList.json")
	if err != nil {
		t.Errorf("Expected to read file without error, got: %v", err)
	}

	var tasks []*Task
	err = json.Unmarshal(bytes, &tasks)
	if err != nil {
		t.Errorf("Expected to unmarshal JSON without error, got: %v", err)
	}

	if len(tasks) != 2 {
		t.Errorf("Expected to load 2 tasks from file, found %d", len(tasks))
	}
}
