package event

import "fmt"

// Creates new event in 01_EVENTS dir and creates sub dirs for event
func New(name string, unit UnitEvents) {
	// create event dir
	// create subdirs

	pth := unit.UnitDir()
	fmt.Printf("will create into: %s\n", pth)
}
