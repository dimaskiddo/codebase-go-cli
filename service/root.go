package service

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CLI Root Variable Structure
var rootCLI = &cobra.Command{
	Use:   "cli-go",
	Short: "Go CLI Application",
	Long:  "A Simple Go CLI Application",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// ExecRoot Function to Execute CLI Root
func ExecRoot() {
	err := rootCLI.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
