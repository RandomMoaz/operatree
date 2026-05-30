package project

import (
	"log"
	"path/filepath"

	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/pkg/module"
	"github.com/hanymamdouh82/operatree/pkg/subject"
	"gopkg.in/yaml.v3"
)

// Sync walks the full project tree and updates each subject in memory
// from its metadata file on disk, then writes the project metadata file.
func Sync(p *Project) error {
	for i := range p.Modules {
		if err := syncModule(&p.Modules[i]); err != nil {
			return err
		}
	}

	return p.WriteMetadata()
}

func syncModule(m *module.Module) error {
	// Sync subjects at this level
	for j, s := range m.Subjects {
		b, err := filesystem.ReadFile(filepath.Join(s.DirName, subject.METADATA_FILE))
		if err != nil {
			// subject file missing or unreadable — skip, don't abort
			log.Printf("missing yml for subject %s\n", s.DirName)
			continue
		}

		var diskMeta subject.Subject
		if err := yaml.Unmarshal(b, &diskMeta); err != nil {
			// malformed yaml — skip, don't abort
			log.Printf("malformed yml for subject %s\n", s.DirName)
			continue
		}

		m.Subjects[j] = diskMeta
	}

	// Recurse into sub-modules
	for i := range m.Modules {
		if err := syncModule(&m.Modules[i]); err != nil {
			return err
		}
	}

	return nil
}
