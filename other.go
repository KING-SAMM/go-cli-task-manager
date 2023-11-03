package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func addTask(task string) {
	file, err := os.OpenFile("tasks/tasks.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	_, err = fmt.Fprintln(file, task)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func listTasks() {
	file, err := os.Open("tasks/tasks.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 1; scanner.Scan(); i++ {
		fmt.Printf("%d %s \n", i, scanner.Text())
	}
}

func removeTask(taskNum int) {
	// Open the tasks.txt file for reading
	file, err := os.Open("tasks/tasks.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Read tasks into a slice
	var tasks []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		tasks = append(tasks, scanner.Text())
	}

	// Check if taskNum is out of range
	if taskNum < 1 || taskNum > len(tasks) {
		fmt.Println("Invalid task number. Please provide a valid task number.")
		return
	}

	// Remove the specified task
	removedTask := tasks[taskNum-1]
	tasks = append(tasks[:taskNum-1], tasks[taskNum:]...)

	// Open the tasks.txt file for writing
	file, err = os.Create("tasks/tasks.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Write the updated tasks back to the file
	for _, task := range tasks {
		fmt.Fprintln(file, task)
	}

	fmt.Printf("Task '%s' removed successfully!\n", removedTask)
}


func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a command (add, list, or remove).")
		return
	}

	command := os.Args[1]

	switch command {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide a task to add.")
			return
		}
		task := os.Args[2]
		addTask(task)
		fmt.Println("Task added successfully!")

	case "list":
		listTasks()

	case "remove":
		if len(os.Args) < 3 {
			fmt.Println("Please provide the task number to remove.")
			return
		}
		taskNum := os.Args[2]
		// Convert taskNum to an integer and call removeTask
		taskNumInt, err := strconv.Atoi(taskNum)
		if err != nil {
			panic(err.Error())
		}
		// Implement this part yourself
		removeTask(taskNumInt)

	default:
		fmt.Println("Invalid command. Please use 'add', 'list', or 'remove'.")
	}
}
