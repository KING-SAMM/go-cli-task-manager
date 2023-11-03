package main

import (
	"fmt"
	"time"
)

type TaskGroup struct {
	name string
	items map[string]bool
	// completed bool
}

// Create new TaskGroup
func newTaskGroup(name string) TaskGroup {
	tg := TaskGroup{
		name: name,
		items: map[string]bool{time.Now().Format(time.RFC1123)+ ": \t" +"Study algoritms": false, time.Now().Format(time.RFC1123)+ ": \t" +"Finsh my book": false},
		// completed: false,
	}

	return tg
}

// Format TaskGroup
func (tg TaskGroup) format() string {
	fs := "TASK GROUP: \n"
	fs += "==========\n"
	fs += "Time Created  \t\t\tTask  \t\t\tCompleted \n"
	fs += "------------  \t\t\t----- \t\t\t---------- \n"

	for k, v := range tg.items {
		fs += fmt.Sprintf("%-12v  ----\t%v \n", k+"\t", v)
	}

	return fs
}
