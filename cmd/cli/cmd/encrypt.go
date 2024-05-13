package cmd

import (
	"LucasNT/simple-encryption-library/pkg"
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var encryptCmd = &cobra.Command{
	Use:   "encrypt [flags] <FILE>",
	Short: "encrypt Data",
	Long:  "encrypt Data from a file or stdin",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a file input, or - for stdin")
		}
		if args[0] != "-" {
			stats, err := os.Stat(args[0])
			if err != nil {
				return fmt.Errorf("Input has a problemna %w", err)
			}
			if stats.IsDir() {
				return errors.New("Input is a directory")
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		pubKey, err := pkg.LoadPubKey(keyPath + ".pub")
		if err != nil {
			log.Fatal(err)
		}

		var input io.ReadCloser
		if args[0] == "-" {
			input = os.Stdin
		} else {
			input, err = os.Open(args[0])
			if err != nil {
				log.Fatal(err)
			}
			defer input.Close()
		}
		encryptedData, err := pkg.EncryptData(input, pubKey)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print(string(encryptedData))
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)
}
