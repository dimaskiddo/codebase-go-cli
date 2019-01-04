package service

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CLI Root Variable Structure
var rootCLI = &cobra.command{
	Use:   "Your Command",
	Short: "Command Short Description",
	Long:  "Command Full Description",
}

// ExecRoot Function to Execute CLI Root
func ExecRoot() {
	err := rootCLI.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
