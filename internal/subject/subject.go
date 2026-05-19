package subject

import (
	"path"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
)

type SubjectType string

const (
	SubjectEvent SubjectType = "EVENT"
	SubjectTask  SubjectType = "TASK"
)

var (
	SubDirs map[SubjectType][]string = map[SubjectType][]string{
		SubjectEvent: {
			"01_AGENDA",
			"02_MEDIA",
			"03_NOTES",
			"04_DOCUMENTS",
			"05_OUTCOMES",
		},
		SubjectTask: {
			"01_INPUTS",
			"02_WORKING",
			"03_REVIEW",
			"04_FINAL",
		},
	}
)

// managed by operatree, can add/delete/edit, etc
// searchable, indexable, parsed by describe()
// This is like 01_RAW inside 06_DATA module
type Subject struct {
	Type        SubjectType `yaml:"type"`
	Name        string      `yaml:"name"`
	DirName     string      `yaml:"dirName"`
	SubDirs     []string    `yaml:"subDirs"`
	Date        string      `yaml:"date"`
	Tags        []string    `yaml:"tags"`
	Notes       string      `yaml:"notes"`
	Paricipants []string    `yaml:"paricipants,omitempty"` // omitempty guarantees that field written only for Subject that needs it
	Location    string      `yaml:"location,omitempty"`    // omitempty guarantees that field written only for Subject that needs it
}

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
		sdp := path.Join(s.DirName, v)

		if err := filesystem.CreateDir(sdp); err != nil {
			return err
		}
	}

	return nil
}

func (s *Subject) WriteMetadata() error {

	fn := path.Join(s.DirName, "metadata.yml")
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

	// write metadata file
	if err := s.WriteMetadata(); err != nil {
		return err
	}

	return nil
}
