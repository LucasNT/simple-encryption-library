package main

import (
	"log"

	"github.com/LucasNT/simple-encryption-library/cmd/simple-encrypt-file/cmd"
)

func init() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
}
