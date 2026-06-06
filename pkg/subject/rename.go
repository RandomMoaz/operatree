package subject

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/internal/metadata"
)

// renames a subject. It provides a simple interactive CLI or silent mode to rename the subject.
// Renames the subject dir and recreates the METADATA.yml file
// Project receiver function is responsible for updating project metadata not subject.
// `nn` is the new name
func renameSubject(s *Subject, nn string) error {
	// check new name to either show interactive CLI or use provided name

	// interactive cli
	if nn == "" {
		inn, err := interactiveRename(s.Name)
		if err != nil {
			return err
		}

		nn = inn
	}

	// normalize name
	nn = metadata.FormatName(nn)

	// silent doesn't require a separate function, the provided name is the new name to be used

	// identify base dir
	bd, found := strings.CutSuffix(s.DirName, s.Name)
	if !found {
		return fmt.Errorf("failed to extract dir name from subject name")
	}

	newDirName := filepath.Join(filepath.Clean(bd), nn)

	// call filesystem to rename
	if err := filesystem.RenameDir(s.DirName, newDirName); err != nil {
		return nil
	}

	// update subject struct for new name and updated dir
	s.DirName = newDirName
	s.Name = nn

	// call write metadata to overwrite the metadata with new name
	if err := s.WriteMetadata(); err != nil {
		return err
	}

	return nil
}

func interactiveRename(oldName string) (string, error) {
	var nn string = oldName

	// Standard fields — all types
	err := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("New Name").
				Value(&nn),
		),
	).Run()
	if err != nil {
		return nn, err
	}

	return nn, nil
}
