package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

const fileName = "task.json"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func loadTasks() ([]Task, error) {
	var tasks []Task
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return tasks, nil
	}
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &tasks)
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, 0644)
}

func getNextID(tasks []Task) int {
	maxID := 0
	for _, task := range tasks {
		if task.ID > maxID {
			maxID = task.ID
		}
	}
	return maxID + 1
}

func addTask(desc string) {
	tasks, _ := loadTasks()
	now := time.Now()
	task := Task{
		ID:          getNextID(tasks),
		Description: desc,
		Status:      "todo",
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	tasks = append(tasks, task)
	saveTasks(tasks)
	fmt.Printf("Task added successfully (ID: %d)\n", task.ID)
}

func updateTask(id int, desc string) {
	tasks, _ := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = desc
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Println("Task updated successfully.")
			return
		}
	}
	fmt.Println("Task not found.")
}

func deleteTask(id int) {
	tasks, _ := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks(tasks)
			fmt.Println("Task deleted successfully.")
			return
		}
	}
	fmt.Println("Task not found.")
}

func markStatus(id int, status string) {
	tasks, _ := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			saveTasks(tasks)
			fmt.Printf("Task marked as %s.\n", status)
			return
		}
	}
	fmt.Println("Task not found.")
}

func listTasks(filter string) {
	tasks, _ := loadTasks()
	for _, task := range tasks {
		if filter == "" || task.Status == filter {
			fmt.Printf("ID: %d | %s | %s\n", task.ID, task.Description, task.Status)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command.")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add \"Task description\"")
			return
		}
		addTask(os.Args[2])

	case "update":
		if len(os.Args) < 4 {
			fmt.Println("Usage: task-cli update <id> \"New description\"")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		updateTask(id, os.Args[3])

	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		deleteTask(id)

	case "mark-in-progress":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		markStatus(id, "in-progress")

	case "mark-done":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli mark-done <id>")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		markStatus(id, "done")

	case "list":
		filter := ""
		if len(os.Args) > 2 {
			filter = os.Args[2]
		}
		listTasks(filter)

	default:
		fmt.Println("Unknown command.")
	}
}
