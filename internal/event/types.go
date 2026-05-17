package event

import "fmt"

type Event struct {
	Name     string `yaml:"name"`
	Location string `yaml:"location"`
}

// Cannot use *Event since it will not implement the interface
func (e Event) Bootstrap() error {
	fmt.Println("I will bootstrap Event")
	return nil
}
