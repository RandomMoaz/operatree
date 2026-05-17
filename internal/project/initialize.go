package project

import (
	"fmt"
	"os"

	"github.com/hanymamdouh82/operatree/internal/event"
)

// bootstraps a project by creating directory structure for primary entities
// pth is the parent path where project will be created
// In future features, pth will be stored into DB
func Bootstrap(pth string, name string) (Project, error) {
	// Validate path is not existing
	// Create dir
	// Create subdirs -> each subdir uses Project.<receiver_function()>

	p := Project{
		name:    name,
		baseDir: pth,
	}

	if err := os.Mkdir(p.ProjectDir(), os.ModeAppend); err != nil {
		return p, err
	}

	// Bootstrap events
	evt := event.Event{}
	p.Units = append(p.Units, evt)

	// iterate and invoke every bootstrap function
	// we collect bootstrapping results and we don't interrupt the process
	bes := []error{}
	for _, v := range p.Units {
		if err := v.Bootstrap(); err != nil {
			bes = append(bes, err)
		}
	}

	// dump errors
	for _, v := range bes {
		fmt.Printf("%s\n", v.Error())
	}

	return p, nil
}
