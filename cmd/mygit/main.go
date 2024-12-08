package main

import (
	"fmt"
	"log"
	"os"

	"github.com/codecrafters-io/git-starter-go/cmd/mygit/git"
	"github.com/urfave/cli/v2"
)

// Usage: your_git.sh <command> <arg1> <arg2> ...
func main() {
	app := &cli.App{
		Name:  "mygit ",
		Usage: "[-v | --version] [-h | --help] [-C <path>] [-c <name>=<value>]\n[--exec-path[=<path>]] [--html-path] [--man-path] [--info-path]\n[-p | --paginate | -P | --no-pager] [--no-replace-objects] [--bare]\n[--git-dir=<path>] [--work-tree=<path>] [--namespace=<name>]\n[--config-env=<name>=<envvar>] <command> [<args>]",
		Action: func(cCtx *cli.Context) error {
			if err := git.Action(cCtx); err != nil {
				fmt.Fprintf(os.Stderr, err.Error())
				os.Exit(1)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
