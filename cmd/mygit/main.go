package main

import (
	"context"
	"log"
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/git"
	"github.com/urfave/cli/v3"
)

// Usage: your_git.sh <command> <arg1> <arg2> ...
func main() {
	g := git.NewGit()

	cmd := &cli.Command{
		Name:  "mygit",
		Usage: "A simple git command line tool",
		Commands: []*cli.Command{
			g.Init(),
			g.CatFile(),
			g.OpenFile(),
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
