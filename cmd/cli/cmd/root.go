package cmd

import (
	"fmt"
	"os"

	"filippo.io/age"
	"github.com/spf13/cobra"
)

var (
	keyPath string
	rootCmd = &cobra.Command{
		Use:   os.Args[0],
		Short: "Teste",
		Long:  "Teste Longo",
	}
	key *age.X25519Identity
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	path := homeDir + "/.local/encryptKey/key"
	rootCmd.PersistentFlags().StringVar(&keyPath, "key", path, "Path to the key, default is $HOME/.local/usr/share/key")
}

func initialize() {
}
