package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/hanymamdouh82/operatree/pkg/config"
	"github.com/spf13/cobra"
)

const dFlagHelp_project = "project directory (default: default project, '.' for current dir, or an absolute path)"
const dFlagHelp_baseDir = "base directory (default: config standardDir, '.' for current dir, or an absolute path)"

var (
	verbose bool          // verbose flag
	cfg     config.Config // loaded config
	destDir string        // initial -d value enetered by user
	actDir  string        // The actual dir will be used, all commands should use this variable only
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
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
