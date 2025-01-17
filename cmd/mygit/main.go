package main

import (
	"fmt"
	"log"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/git"
)

func printUsage() {
	fmt.Println("Usage: mygit <command> [<args>]")
	fmt.Println("\nAvailable commands:")
	fmt.Println("  init                   Initialize a new git repository")
	fmt.Println("  cat-file -p <blob>      Pretty-print the contents of a git object")
}

func main() {
	g := git.NewGit()
	err := g.Run()
	if err != nil {
		printUsage()
		log.Fatal(err)
	}
}
