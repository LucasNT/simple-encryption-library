package main

import (
	"LucasNT/simpleEncryptFile/cmd/cli/cmd"
	"log"

	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

func main() {
}
