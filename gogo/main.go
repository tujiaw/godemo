package main

import (
	"log"
	"os"
	_ "./matchers"
	"./search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
