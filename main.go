package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	ctl "github.com/dimaskiddo/codebase-go-cli/controller"
)

// Root Variable Structure
var r = &cobra.Command{
	Use:   "cli-go",
	Short: "Go CLI Application",
	Long:  "Go CLI Application",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

// Init Function
func init() {
	// Add Child for Root Command
	r.AddCommand(ctl.Version)
	r.AddCommand(ctl.Pwd)
}

// Main Function
func main() {
	err := r.Execute()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
