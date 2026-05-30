package project

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/hanymamdouh82/operatree/internal/template"
	"github.com/hanymamdouh82/operatree/pkg/module"
)

func TestFactory(t *testing.T) {

	want := Project{
		Name:     "dummy-proj",
		Template: "dummy",
		absDir:   "/tmp/operatree-demo/dummy-proj",
		Modules: []module.Module{
			{
				Name:    "00_ADMIN",
				Type:    "ADMIN",
				Modules: []module.Module{},
				AbsPath: "/tmp/operatree-demo/dummy-proj/00_ADMIN",
				SubDirs: []string{
					"sub1",
				},
			},
			{
				Name:    "02_PROJECT_MANAGEMENT",
				Type:    "PROJECT_MANAGEMENT",
				AbsPath: "/tmp/operatree-demo/dummy-proj/02_PROJECT_MANAGEMENT",
				Modules: []module.Module{
					{
						Name:    "05_TASKS",
						Type:    "TASKS",
						AbsPath: "/tmp/operatree-demo/dummy-proj/02_PROJECT_MANAGEMENT/05_TASKS",
						Modules: []module.Module{},
						SubDirs: []string{
							"sub2",
						},
					},
				},
			},
		},
	}

	// load module
	gotTemplate, gotErr := template.Load("template_dummy.yml")
	if gotErr != nil {
		t.Errorf("error: %s", gotErr.Error())
	}

	// parse template and construct project struct
	gotProject, err := Factory(gotTemplate, "dummy-proj")
	if err != nil {
		t.Errorf("error: %s", err.Error())
	}

	gotProject.absDir = ""

	// hydrate
	hydratePath("/tmp/operatree-demo/dummy-proj", &gotProject)

	// Compare using EquateEmpty which treats nil slices/maps the same as empty ones
	if diff := cmp.Diff(want, gotProject, cmpopts.EquateEmpty(), cmp.AllowUnexported(Project{})); diff != "" {
		t.Errorf("Factory() mismatch (-want +got):\n%s", diff)
	}
}
