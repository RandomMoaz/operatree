package template

import "fmt"

func ListTemplates() {
	for _, v := range Templates {
		t, err := Load(v)
		if err != nil {
			// we skip undefined or not found template
			continue
		}

		fmt.Printf("%-20s - %s\n", t.Name, t.Description)
	}
}
