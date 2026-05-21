package filesystem

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"gopkg.in/yaml.v3"
)

// if found return with error, any other case return nil
func CheckDirExists(pth string) error {

	_, err := os.Stat(pth)

	// exists, return with error
	if err == nil {
		return fmt.Errorf("dir exists")
	}

	// doesn't exist -> return with nil error
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}

	// default case, cannot identify error -> assume exists and return with error
	return err
}

func CreateDir(pth string) error {

	if err := CheckDirExists(pth); err != nil {
		return err
	}

	if err := os.MkdirAll(pth, 0775); err != nil {
		return err
	}

	return nil
}

// `fullname` is the full file name including abs path
func StructToFile(s any, fullName string) error {

	b, err := yaml.Marshal(s)
	if err != nil {
		return err
	}

	if err := os.WriteFile(fullName, b, 0775); err != nil {
		return err
	}

	return nil
}

// `fullname` is the full file name including abs path
func TextToMDFile(s string, fullName string) error {
	b := []byte(s)

	if err := os.WriteFile(fullName, b, 0775); err != nil {
		return err
	}

	return nil
}

// Reads a file. Although it directly returns os.ReadFile, it is abstracted
// so in future implementations we can add daemon reads, DB reads, etc
// It is also abstract to handle edge-cases of different OS
func ReadFile(mp string) ([]byte, error) {

	return os.ReadFile(mp)
}
