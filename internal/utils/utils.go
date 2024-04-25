package utils

import (
	"fmt"
	"syscall"

	"golang.org/x/term"
)

func ReadPassword() (string, error) {
	fmt.Println("Password: ")
	bytePwd, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		return "", fmt.Errorf("Failed to read password, %w", err)
	}
	pass := string(bytePwd)
	return pass, nil
}
