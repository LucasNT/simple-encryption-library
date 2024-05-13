package cmd

import (
	"LucasNT/simple-encryption-library/internal/utils"
	"fmt"
	"log"
	"os"

	"filippo.io/age"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   os.Args[0],
		Short: "Encrypt and Decrypt Data lib",
		Long:  "Teste Longo",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if s, exists := os.LookupEnv("KEY_PASSWORD"); exists {
				passwordKey = s
			}
			if isReadPassword {
				var err error
				passwordKey, err = utils.ReadPassword()
				if err != nil {
					log.Fatal(err)
				}
			}
		},
	}
	keyPath        string
	key            *age.X25519Identity
	isReadPassword bool
	passwordKey    string = ""
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
	rootCmd.PersistentFlags().StringVar(&keyPath, "key", path, "Path to the key")
	rootCmd.PersistentFlags().BoolVarP(&isReadPassword, "password", "p", false, "Programn should ask for password")
}

func initialize() {
}
