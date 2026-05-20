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

func init() {
	// rootCmd.PersistentFlags().StringVarP(&prjDir, "dest", "d", ".", "project directory")
	rootCmd.PersistentFlags().StringVarP(&prjDir, "dest", "d", "/mnt/extra/onfly/testprj", "project directory")
}

var rootCmd = &cobra.Command{
	Use:   "operatree",
	Short: "OperaTree project operating system",
	Long:  "OperaTree is your project operating system built on your filesystem",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome OperaTree...A project operating system built on your filesystem...!")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
