package cmd

import (
	"errors"
	"log"
	"os"

	"filippo.io/age"
	"github.com/LucasNT/simple-encryption-library/internal/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	rootCmd = &cobra.Command{
		Use:   os.Args[0],
		Short: "Encrypt and Decrypt Data lib",
		Long:  "Teste Longo",
	}
	key            *age.X25519Identity
	isReadPassword bool
)

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initialize)
	rootCmd.PersistentFlags().String("key", "", "Path to the key")
	rootCmd.PersistentFlags().BoolVarP(&isReadPassword, "password", "p", false, "Programn should ask for password")
	viper.BindPFlag("KEY", rootCmd.PersistentFlags().Lookup("key"))
	defaultKeyPath, err := utils.GetDefaultKeyPath()
	if err != nil {
		log.Panic(err)
	}
	viper.SetDefault("KEY", defaultKeyPath)
}

func initialize() {
	viper.AddConfigPath(".")
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		if !errors.As(err, &viper.ConfigFileNotFoundError{}) {
			log.Panic(err)
		}
	}
	if isReadPassword {
		passwordKey, err := utils.ReadPassword()
		if err != nil {
			log.Fatal(err)
		}
		viper.Set("PASSWORD", passwordKey)
	}
}
