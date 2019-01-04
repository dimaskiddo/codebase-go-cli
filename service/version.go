package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CLI Version Variable Structure
var versionCLI = &cobra.Command{
	Use:   "version",
	Short: "Show the Go CLI version",
	Long:  "Show the Go CLI version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go CLI Version 1.0")
	},
}

func init() {
	// Add CLI Version as Child of CLI Root Command
	rootCLI.AddCommand(versionCLI)
}
