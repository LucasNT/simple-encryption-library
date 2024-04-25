package main

import (
	"LucasNT/simpleEncryptFile/cmd/cli/cmd"
	"log"
)

func init() {
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
}
