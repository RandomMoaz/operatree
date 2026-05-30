package runner

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/hanymamdouh82/operatree/pkg/config"
)

func EditFile(filePath string) error {

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("cannot load operatree config")
	}

	editor := cfg.Editor
	if editor == "" {
		editor = os.Getenv("EDITOR")
	}

	if editor == "" {
		return fmt.Errorf("cannot find default editor")
	}

	if err := run(editor, []string{filePath}); err != nil {
		return err
	}

	return nil
}

func OpenFileManager(filePath string) error {

	cfg, err := config.Load()
	if err != nil {
		return fmt.Errorf("cannot load operatree config")
	}

	fileManager := cfg.FileManager
	if fileManager == "" {
		return fmt.Errorf("cannot find default file manager")
	}

	if err := run(fileManager, []string{filePath}); err != nil {
		return err
	}

	return nil
}

// Run docker compose for file and services
// if service array is empty, it will run the all services in the compose file
// stack is governed by caller function
func run(prog string, args []string) error {

	// Get abs path for Docker binary
	binary, err := exec.LookPath(prog)
	if err != nil {
		return err
	}

	cmd := exec.Command(binary, args...)
	cmd.Stdout = os.Stdout // stream directly
	cmd.Stderr = os.Stderr // capture docker's actual output
	cmd.Stdin = os.Stdin   // needed if docker prompts for anything

	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}
