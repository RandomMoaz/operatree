package project

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// Loads a project by reading project metadata file and sets project structrue
// Path represents project root path
func Load(pth string) (Project, error) {

	b, err := os.ReadFile(filepath.Join(pth, METADATA_FILE))
	if err != nil {
		return Project{}, err
	}

	// unmarshal into loader struct, this is because Unit is an interface not a struct
	// then we can convert units as per build logic
	var p Project
	if err := yaml.Unmarshal(b, &p); err != nil {
		return Project{}, err
	}

	hydratePath(pth, &p)

	return p, err
}
