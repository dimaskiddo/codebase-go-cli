package service

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CLI Working Directory Variable Structure
var pwdCLI = &cobra.Command{
	Use:   "pwd",
	Short: "Get Current Working Directory",
	Long:  "Get Current Working Directory",
	Run: func(cmd *cobra.Command, args []string) {
		pwd, err := os.Getwd()
		if err != nil {
			fmt.Println("Error to Get Current Working Directory!")
			os.Exit(1)
		}

		fmt.Println(pwd)
	},
}
