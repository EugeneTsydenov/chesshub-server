package main

import (
	"log"

	"github.com/EugeneTsydenov/chesshub-server/infrastructure/env"
)

func main() {
	err := env.Load()
	if err != nil {
		log.Fatal("main.go: failed load environments")
	}
}
