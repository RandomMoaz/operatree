package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of operatree",
	Long:  `All software has versions. This is Operatree's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("OperaTree v0.0.1 -- HEAD")
	},
}
