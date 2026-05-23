package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

// Resolves -d flag as a base dir. Never used with handlers requirs project dir
// If -d is not provided it fallback to cfg.StandardDir
// If cfg.StandardDir is not configured it exits with error
// It supports resolving "-d ." to working dir abs path
func resolveBaseDir(cmd *cobra.Command, args []string) {
	// -d is used, set actDir to value if destDir and return
	if destDir != "" {

		// the case of current working dir
		if destDir == "." {
			abs, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}

			actDir = abs
			return
		}

		actDir = destDir
		return
	}

	// use cfg
	if cfg.StandardDir == "" {
		log.Fatal("-d is not provided and standardDir is not set by operatree init")
	}

	actDir = cfg.StandardDir
}

// This is to be used with any command that requires a project.
// If -d is present, use it (resolve "." if it is the value)
// If -d is not provided it fallback to cfg.DefaultProject
// If -d is not provided and cfg.DefaultProject is set, error
func resolveProjectDir(cmd *cobra.Command, args []string) {

	// 1. explicit -d flag -> user is responsible for correctness of -d value
	if destDir != "" {

		// the case of current working dir
		if destDir == "." {
			abs, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}

			actDir = abs
			return
		}

		actDir = destDir
		return
	}

	if cfg.Default.AbsPath == "" {
		log.Fatal("-d is not provided and default is not set by operatree init")
	}

	// finally set the variable that all handlers depend on
	actDir = cfg.Default.AbsPath
}
