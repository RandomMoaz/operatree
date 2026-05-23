package project

import (
	"fmt"

	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/hanymamdouh82/operatree/internal/filesystem"
)

// Bootstraps a project by creating project struct and call bootstrap modules
// Bootstrap calls different bootstrap functions based on template.
// `bpth` is the abs base path for the project. It souldn't include project name
// `t` template name
func Bootstrap(name string, bpth string, t string) (Project, error) {

	if name == "" {
		return Project{}, fmt.Errorf("project name is missing")
	}

	if bpth == "" {
		return Project{}, fmt.Errorf("project bath is missing, either -d is missing value or init is not used")
	}

	// get template factory from templates map
	tf, ok := Templates[t]
	if !ok {
		return Project{}, fmt.Errorf("undefined template")
	}
	np := tf(name, bpth)
	np.Template = t

	// create project dir
	if err := filesystem.CreateDir(np.ProjectDir()); err != nil {
		return np, err
	}

	// bootstrap modules
	// We collect errors without preventing creation of next module
	var merrs []error
	for _, m := range np.Modules {
		if err := m.Bootstrap(); err != nil {
			merrs = append(merrs, err)
		}
	}

	// write project metadata
	if err := np.WriteMetadata(); err != nil {
		return np, err
	}

	// Register in config
	if err := config.AddProject(name, np.ProjectDir(), t); err != nil {
		return np, err
	}

	return np, nil
}
