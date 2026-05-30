package module

import (
	"path/filepath"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
)

// A method to create module directory
func (m *Module) MkDir() error {

	if err := filesystem.CreateDir(m.AbsPath); err != nil {
		return err
	}

	return nil
}

// A method to create module sub directories
func (m *Module) MkSubDirs() error {

	for _, v := range m.SubDirs {
		sdp := filepath.Join(m.AbsPath, v)

		if err := filesystem.CreateDir(sdp); err != nil {
			return err
		}
	}

	return nil
}

// Creates module dirs, subdirs, nested modules, metadata templates, etc
// This is to be used during project bootstrapping, or new module bootstrapping
func (m *Module) Bootstrap() error {
	// Create module directory
	if err := m.MkDir(); err != nil {
		return err
	}

	// Recursive bootstrapping for submodules
	for _, v := range m.Modules {
		v.Bootstrap()
	}

	// create module subdirs
	if err := m.MkSubDirs(); err != nil {
		return err
	}

	return nil
}
