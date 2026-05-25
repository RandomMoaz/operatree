package template

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestLoad(t *testing.T) {
	want := OTTemplate{
		Name:        "dummy",
		Description: "dummy description",
		Modules: []ModuleTemplate{
			{
				Description: "admin description",
				Name:        "admin",
				Type:        "admin",
				Modules:     []ModuleTemplate{},
				SubDirs: []string{
					"sub1",
				},
			},
			{
				Description: "project_management description",
				Name:        "project_management",
				Type:        "project_management",
				Modules: []ModuleTemplate{
					{
						Description: "tasks description",
						Name:        "tasks",
						Type:        "tasks",
						SubDirs: []string{
							"sub2",
						},
					},
				},
			},
		},
	}

	got, err := Load("template_dummy.yml")
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	if !reflect.DeepEqual(got, want) {
		jGot, _ := json.MarshalIndent(got, "", "  ")
		jWant, _ := json.MarshalIndent(want, "", "  ")
		t.Errorf("got:\n%s\n, want:\n%s\n", jGot, jWant)
	}
}
