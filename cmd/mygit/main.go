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
	cmd := &cli.Command{
		Name:  "mygit ",
		Usage: "[-v | --version] [-h | --help] [-C <path>] [-c <name>=<value>]\n[--exec-path[=<path>]] [--html-path] [--man-path] [--info-path]\n[-p | --paginate | -P | --no-pager] [--no-replace-objects] [--bare]\n[--git-dir=<path>] [--work-tree=<path>] [--namespace=<name>]\n[--config-env=<name>=<envvar>] <command> [<args>]",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			git.Action(ctx, cmd)
			return nil
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
