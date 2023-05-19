package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Here we assume that the TodoList struct and its methods are defined in todo_list.go
var todoList *TodoList = NewTodoList()

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("To-Do List Application")

	for {
		fmt.Print("\nEnter command: ")
		command, _ := reader.ReadString('\n')
		runCommand(command)
	}
}

func runCommand(command string) {
	command = strings.TrimSpace(command)
	args := strings.Split(command, " ")

	if len(args) == 0 {
		fmt.Println("No command provided.")
		return
	}

	switch args[0] {
	case "add":
		if len(args) < 2 {
			fmt.Println("Please provide a task description.")
			return
		}
		// You might want to handle parsing dates and priorities here as well
		task := NewTask(args[1], time.Now(), 1)
		todoList.AddTask(task)
		fmt.Println("Task added.")
	case "remove":
		// args[1] should be parsed to an int
		if len(args) < 2 {
			fmt.Println("Please provide a task number.")
			return
		}
		taskNumber, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("Please provide a valid task number.")
			return
		}
		todoList.RemoveTask(taskNumber)
		fmt.Println("Task removed.")
	case "list":
		for i, task := range todoList.GetTaskList() {
			fmt.Printf("%d: %s\n", i+1, task.Description)
		}
	case "help":
		if len(args) > 1 {
			switch args[1] {
			case "add":
				fmt.Println("add: Add a new task to the list. Usage: add [task description]")
			case "remove":
				fmt.Println("remove: Remove a task from the list. Usage: remove [task number]")
			case "list":
				fmt.Println("list: List all tasks.")
			default:
				fmt.Println("Unknown command for help.")
			}
		} else {
			fmt.Println("Available commands: add, remove, list, help")
		}
	case "quit", "exit", "q":
		fmt.Println("Exiting program.")
		os.Exit(0)
	default:
		fmt.Println("Unknown command.")
	}
}
