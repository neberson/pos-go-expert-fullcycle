package main

import (
	"fmt"

	"github.com/neberson/fcutils/pkg/events"
)

func main() {
	ed := events.NewEventDispatcher()
	fmt.Println(ed)
}
