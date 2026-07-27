package project

import (
	"errors"
	"io/fs"
	"log"
	"path/filepath"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
	"gopkg.in/yaml.v3"
)

// Sync walks the full project tree and updates each subject in memory
// from its metadata file on disk, then writes the project metadata file
// only if something actually changed on disk.
func Sync(p *Project) error {
	dirty := false
	for i := range p.Modules {
		changed, err := syncModule(&p.Modules[i])
		if err != nil {
			return err
		}
		dirty = dirty || changed
	}

	if !dirty {
		return nil
	}

	return p.WriteMetadata()
}

func syncModule(m *module.Module) (bool, error) {
	dirty := false

	// Sync subjects at this level
	for j := range m.Subjects {
		s := m.Subjects[j]
		b, err := filesystem.ReadFile(filepath.Join(s.DirName, subject.METADATA_FILE))
		if err != nil {
			if errors.Is(err, fs.ErrNotExist) {
				// subject file genuinely missing — skip, don't abort
				log.Printf("missing yml for subject %s\n", s.DirName)
			} else {
				// real read error (permissions, I/O) — skip but surface it
				log.Printf("cannot read yml for subject %s: %v\n", s.DirName, err)
			}
			continue
		}

		var diskMeta subject.Subject
		if err := yaml.Unmarshal(b, &diskMeta); err != nil {
			// malformed yaml — skip, don't abort
			log.Printf("malformed yml for subject %s: %v\n", s.DirName, err)
			continue
		}

		
		diskMeta.DirName = s.DirName
		diskMeta.Files = s.Files

		m.Subjects[j] = diskMeta
		dirty = true
	}

	
	for i := range m.Modules {
		changed, err := syncModule(&m.Modules[i])
		if err != nil {
			return dirty, err
		}
		dirty = dirty || changed
	}

	return dirty, nil
}
