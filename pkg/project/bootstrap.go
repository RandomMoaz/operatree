package project

import (
	"fmt"
	"path/filepath"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/internal/template"
	"github.com/hanymamdouh82/operatree/pkg/config"
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

	// get template file name from provided template name using tmpltMap
	tn, ok := template.Templates[t]
	if !ok {
		return Project{}, fmt.Errorf("undefined template")
	}

	// load template file
	pt, err := template.Load(tn)
	if err != nil {
		return Project{}, fmt.Errorf("undefined template file")
	}

	// run factory -> converts template to Project struct
	np, err := Factory(pt, name)
	if err != nil {
		return np, err
	}

	// path hydration:
	// Walk project, subjects, modules, nested modules and injects AbsPath, DirName
	// It is crucial to hydrate at runtime to comply with relative-path requirememnts
	ppth := filepath.Join(bpth, name)
	hydratePath(ppth, &np)

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
