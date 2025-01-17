package main

import (
	"log"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/git"
)

func main() {
	g := git.NewGit()
	err := g.Run()
	if err != nil {
		g.PrintUsage()
		log.Fatal(err)
	}
}
