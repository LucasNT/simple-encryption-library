package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var testeCmd = &cobra.Command{
	Use:   "teste",
	Short: "encrypt Data",
	Long:  "encrypt Data from a file or stdin",
	Run: func(cmd *cobra.Command, args []string) {
		keyPath := viper.GetString("KEY")
		passwordKey := viper.GetString("PASSWORD")
		fmt.Println(passwordKey, keyPath)
	},
}

func init() {
	rootCmd.AddCommand(testeCmd)
}
