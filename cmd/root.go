package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/hanymamdouh82/operatree/internal/config"
	"github.com/hanymamdouh82/operatree/internal/project"
	"github.com/spf13/cobra"
)

var (
	baseDir string // abs base dir where project is located. Doesn't include project name
	prjDir  string // abs path of project including its name
	verbose bool   // verbose flag
	cfg     config.Config
)

var rootCmd = &cobra.Command{
	Use:   "operatree",
	Short: "OperaTree project operating system",
	Long:  "OperaTree is your project operating system built on your filesystem",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome OperaTree...A project operating system built on your filesystem...!")
	},
}

func init() {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	cfg = c

	rootCmd.PersistentFlags().StringVarP(&prjDir, "dest", "d", "", "project directory")
	// rootCmd.PersistentPreRunE = resolveProjectDir
	rootCmd.PersistentPreRun = resolveProjectDir
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func resolveProjectDir(cmd *cobra.Command, args []string) {
	// commands that don't need a project dir
	noProjectCmds := map[string]bool{
		"init":      true,
		"bootstrap": true,
		"version":   true,
		"help":      true,
		"explain":   true,
		"default":   true,
	}
	if noProjectCmds[cmd.Name()] {
		return
	}

	// 1. explicit -d flag
	if prjDir != "" {
		return
	}

	// 2. current dir has project metadata
	if isProjectDir(".") {
		prjDir = "."
		return
	}

	// 3. config default
	cfg, err := config.Load()
	if err != nil {
		return
	}
	if cfg.Default.AbsPath != "" {
		prjDir = cfg.Default.AbsPath
		return
	}

	// 4. nothing — friendly error
	// return fmt.Errorf("no project found. Use -d to specify one, or run 'operatree default' to set a default")
	fmt.Errorf("no project found. Use -d to specify one, or run 'operatree default' to set a default")
}

func isProjectDir(path string) bool {
	_, err := os.Stat(filepath.Join(path, project.METADATA_FILE))
	return err == nil
}
