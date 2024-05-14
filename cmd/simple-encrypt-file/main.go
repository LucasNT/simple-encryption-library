package main

import (
	"log"

	"github.com/LucasNT/simple-encryption-library/cmd/cli/cmd"
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