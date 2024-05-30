package cmd

import (
	"fmt"
	"os"

	pkg "github.com/LucasNT/simple-encryption-library/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate Data",
	Long:  "Generate Data from a file or stdin",
	Run: func(cmd *cobra.Command, args []string) {
		keyPath := viper.GetString("KEY")
		passwordKey := viper.GetString("PASSWORD")
		if err := pkg.GenerateKeys(keyPath, passwordKey); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
}
