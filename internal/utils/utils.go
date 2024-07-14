package utils

import (
	"fmt"
	"os"
	"syscall"

	"golang.org/x/term"
)

func ReadPassword() (string, error) {
	fmt.Fprintln(os.Stderr, "Password: ")
	bytePwd, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", fmt.Errorf("Failed to read password, %w", err)
	}
	pass := string(bytePwd)
	return pass, nil
}

func GetDefaultKeyPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	path := homeDir + "/.local/share/simpleEncryptFile"
	if err := os.MkdirAll(path, 0750); err != nil {
		return "", fmt.Errorf("Failed to create the .localFolder %w", err)
	}
	return path, nil

}
