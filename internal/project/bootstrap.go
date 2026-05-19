package project

import (
	"github.com/hanymamdouh82/operatree/internal/filesystem"
)

// Bootstraps a project by creating project struct and call bootstrap modules
// Bootstrap calls different bootstrap functions based on template.
// `bpth` is the abs base path for the project. It souldn't include project name
// To-Do: add templates
func Bootstrap(name string, bpth string) (Project, error) {

	// To-Do: replace with map for different templates
	np := dev(name, bpth)

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

	return np, nil
}
