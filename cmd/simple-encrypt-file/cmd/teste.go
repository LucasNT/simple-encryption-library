package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testeCmd = &cobra.Command{
	Use:   "teste",
	Short: "encrypt Data",
	Long:  "encrypt Data from a file or stdin",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(passwordKey)
	},
}

func init() {
	rootCmd.AddCommand(testeCmd)
}
