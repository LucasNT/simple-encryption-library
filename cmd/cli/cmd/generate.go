package cmd

import (
	"LucasNT/simpleEncryptFile/pkg"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Data",
	Long:  "Generate Data from a file or stdin",
	Run: func(cmd *cobra.Command, args []string) {
		if err := pkg.GenerateKeys(keyPath, passwordKey); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
