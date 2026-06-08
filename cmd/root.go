package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/spf13/cobra"
)

const dFlagHelp_project = "project directory — uses default project if unset, or provide '.' for current dir, a relative path, or an absolute path"
const dFlagHelp_baseDir = "base directory — uses standardDir from config if unset, or provide '.' for current dir, a relative or absolute path"

var (
	verbose       bool          // verbose flag
	cfg           config.Config // loaded config
	destDir       string        // initial -d value enetered by user
	actDir        string        // The actual dir will be used, all commands should use this variable only
	newName, uuid string
)

var rootCmd = &cobra.Command{
	Use:   "operatree",
	Short: "A project operating system built on your filesystem",
	Long: `OperaTree brings structure, searchability, and metadata to your projects
using plain files and directories — no database, no lock-in, no proprietary formats.

Everything is human-readable, Git-friendly, and pipes naturally into standard UNIX tools.
Your data outlives any software.

Get started:
  operatree init          # set up configuration
  operatree create        # create a new project
  operatree show          # inspect configuration and tracked projects

Use 'operatree [command] --help' for more information about a command.`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {
	c, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}
	cfg = c
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
