package service

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Version Variable Structure
var Version = &cobra.Command{
	Use:   "version",
	Short: "Show application version info",
	Long:  "Show application version info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Go CLI Version 1.0")
	},
}
