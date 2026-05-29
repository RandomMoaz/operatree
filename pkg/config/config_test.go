package config

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"testing"
)

func TestConfigDirXDG(t *testing.T) {
	os.Setenv("XDG_CONFIG_HOME", "/home/user/myconfig")
	want := "/home/user/myconfig/operatree"

	gotDir, gotErr := configDir()
	if gotDir != want {
		t.Errorf("config dir = %s; want %s", gotDir, want)
	}

	if gotErr != nil {
		t.Errorf("error: %s", gotErr.Error())
	}
}

func TestStandardPath(t *testing.T) {

	p := `f:\tmp\operatree\project workspace`

	pc := path.Clean(p)
	pfpc := filepath.Clean(p)
	pj := path.Join(p)

	fmt.Printf("path.Clean: %s\n", pc)
	fmt.Printf("filepath.Clean: %s\n", pfpc)
	fmt.Printf("path.Join: %s\n", pj)

}
