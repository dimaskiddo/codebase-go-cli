package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// CLI Version Variable Structure
var versionCLI = &cobra.Command{
	Use:   "version",
	Short: "Get Current CLI Version",
	Long:  "Get Current CLI Version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go CLI Version 1.0")
	},
}
