package subject

import (
	"fmt"
	"path/filepath"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/internal/runner"
)

// A method to create module directory
func (s *Subject) MkDir() error {

	if err := filesystem.CreateDir(s.DirName); err != nil {
		return err
	}

	return nil
}

// A method to create module sub directories
func (s *Subject) MkSubDirs() error {

	for _, v := range s.SubDirs {
		sdp := filepath.Join(s.DirName, v)

		if err := filesystem.CreateDir(sdp); err != nil {
			return err
		}
	}

	return nil
}

// Creates empty files on disk
func (s *Subject) WriteFiles() error {
	for _, f := range s.Files {
		sdp := filepath.Join(s.DirName, f)
		t := fmt.Sprintf("# %s\n", s.Name)
		if err := filesystem.TextToMDFile(t, sdp); err != nil {
			return err
		}
	}

	return nil
}

// Writes metadata.yml file for the subject at subject dir
func (s *Subject) WriteMetadata() error {

	fn := filepath.Join(s.DirName, METADATA_FILE)
	if err := filesystem.StructToFile(s, fn); err != nil {
		return err
	}

	return nil
}

// Writes subject to disk
func (s *Subject) WriteToDisk() error {

	// make dir
	if err := s.MkDir(); err != nil {
		return err
	}

	// make subdirs
	if err := s.MkSubDirs(); err != nil {
		return err
	}

	// Create emtpy subject files
	if err := s.WriteFiles(); err != nil {
		return err
	}

	// write metadata file
	if err := s.WriteMetadata(); err != nil {
		return err
	}

	return nil
}

// Prints subject on screen in a prettey style
// Not all properties are displayed, such as subDirs not included
// This function is intended to be used for subject briefing only
func (s *Subject) Describe() {
	describe(s)
}

func (s *Subject) EditMetadata() error {

	fn := filepath.Join(s.DirName, METADATA_FILE)
	if err := runner.EditFile(fn); err != nil {
		return err
	}

	return nil
}
