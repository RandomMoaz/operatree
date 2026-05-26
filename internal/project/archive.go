package project

import (
	"fmt"
	"os"
	"path"
	"slices"

	"github.com/hanymamdouh82/operatree/internal/activitylog"
	"github.com/hanymamdouh82/operatree/internal/filesystem"
	"github.com/hanymamdouh82/operatree/internal/module"
	"github.com/hanymamdouh82/operatree/internal/subject"
)

// Archive moves subject from its module to project's 99_ARCHIVE module.
//
// it confirm project includes 99_ARCHIVE module before moving.
// It uses filesystem to move subject dir into archive module dir, then
// removes the subject from the module subjects and updates project metadata.
func Archive(p *Project, s subject.Subject) error {

	// confirm project includes 99_ARCHIVE module
	ma, err := p.ModuleExists("99_ARCHIVE")
	if err != nil {
		return err
	}

	an := path.Join(ma.AbsPath, ARCHIVED_DEST, s.Name)
	if err := filesystem.Archive(s.DirName, an); err != nil {
		return err
	}

	for i := range p.Modules {
		if err := updateModule(&p.Modules[i], s); err != nil {
			return err
		}
	}

	// update project metadata
	if err := p.WriteMetadata(); err != nil {
		return err
	}

	if err :=activitylog.Log(
		p.ProjectDir(),
		activitylog.ActionArchive,
		string(s.Type),
		s.Name,
		
	); err !=nil{
		fmt.Fprintf(os.Stderr,"Warning :could not write activity log :%v\n",err)
	}

	return nil
}


func updateModule(m *module.Module, s subject.Subject) error {
	sidx := slices.IndexFunc(m.Subjects, func(ms subject.Subject) bool {
		return ms.DirName == s.DirName
	})

	if sidx != -1 {
		m.Subjects = slices.Delete(m.Subjects, sidx, sidx+1)
	}

	for i := range m.Modules {
		if err := updateModule(&m.Modules[i], s); err != nil {
			return err
		}
	}

	return nil
}
