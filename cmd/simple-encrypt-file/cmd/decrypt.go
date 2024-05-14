package cmd

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"

	pkg "github.com/LucasNT/simple-encryption-library/pkg"
	"github.com/spf13/cobra"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt [flags] <FILE>",
	Short: "Decrypt Data",
	Long:  "Decrypt Data from a file or stdin",
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
		key, err := pkg.LoadKey(keyPath, passwordKey)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
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

		out, err := pkg.DencryptData(input, key)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
