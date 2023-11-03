package main

import (
	"fmt"
)

func main() {
	myTaskGroup := newTaskGroup("November Tasks")

	fmt.Println(myTaskGroup.format())
}