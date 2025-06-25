package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

const dataFile = "todo.json"

func loadTasks() ([]Task, error) {
	file, err := os.Open(dataFile)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil
		}
		return nil, err
	}
	defer file.Close()

	var tasks []Task
	if err := json.NewDecoder(file).Decode(&tasks); err != nil {
		return nil, err
	}
	return tasks, nil
}

func saveTasks(tasks []Task) error {
	file, err := os.Create(dataFile)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(tasks)
}

func nextID(tasks []Task) int {
	maxID := 0
	for _, t := range tasks {
		if t.ID > maxID {
			maxID = t.ID
		}
	}
	return maxID + 1
}

func AddTask(title string) {
	//panic("unimplemented")
	tasks, err := loadTasks()

	if err != nil {
		panic(err)
		return
	}

	newTask := Task{
		ID:    nextID(tasks),
		Title: title,
		Done:  false,
	}

	tasks = append(tasks, newTask)

	if err := saveTasks(tasks); err != nil {
		fmt.Println("Error saving task:", err)
		return
	}

	fmt.Printf("Added task Completed!: [%d] %s\n", newTask.ID, newTask.Title)
}

func ListTasks() {
	//panic("unimplemented")
	tasks, err := loadTasks()
	if err != nil {
		panic(err)
		return
	}
	if len(tasks) == 0 {
		fmt.Print("No tasks found")
	}
	for index, task := range tasks {
		fmt.Println("%d Task:%s", index, task.Title)
	}
}

func CompleteTask(id int) {
	panic("unimplemented")
}

func DeleteTask(id int) {
	panic("unimplemented")
}
