package main

import (
	"LucasNT/simple-encryption-library/cmd/cli/cmd"
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
