package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	baseDir string // abs base dir where project is located. Doesn't include project name
	prjDir  string // abs path of project including its name
	verbose bool   // verbose flag
)

var rootCmd = &cobra.Command{
	Use:   "operatree",
	Short: "OperaTree project operating system",
	Long:  "OperaTree is your project operating system built on your filesystem",
	Run: func(cmd *cobra.Command, args []string) {
		// do stuff here
		// typically this is the entry point, anything to be executed before the command
		fmt.Println("Welcome OperaTree...A project operating system built on your filesystem...!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
