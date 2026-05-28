package config

import (
	"os"
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
