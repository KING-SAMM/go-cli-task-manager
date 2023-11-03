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
		items: map[string]bool{time.Now().Format(time.RFC1123)+ ": " +"Study algoritms ": false, time.Now().Format(time.RFC1123)+ ": " +"Finsh my book ": false},
		// completed: false,
	}

	return tg
}

// Format TaskGroup
func (tg TaskGroup) format() string {
	fs := "TASK GROUP: \n"
	fs += "==========\n"
	fs += "Time Created                   Task                 Completed \n"
	fs += "------------                   -----                ---------- \n"

	for k, v := range tg.items {
		fs += fmt.Sprintf("%-12v ---- %v \n", k+":", v)
	}

	return fs
}
