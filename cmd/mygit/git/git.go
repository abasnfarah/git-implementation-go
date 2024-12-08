package git

import (
	"context"
	"fmt"
	"os"

	"github.com/urfave/cli/v3"
)

func Action(ctx context.Context, cmd *cli.Command) error {
	if cmd.Args().Len() < 1 {
		fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
		return nil
	}

	return nil

	//
	// if len(os.Args) < 2 {
	// 	fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
	// 	os.Exit(1)
	// }
	//
	// switch command := os.Args[1]; command {
	// case "init":
	// 	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
	// 		if err := os.MkdirAll(dir, 0755); err != nil {
	// 			fmt.Fprintf(os.Stderr, "Error creating directory: %s\n", err)
	// 		}
	// 	}
	//
	// 	headFileContents := []byte("ref: refs/heads/master\n")
	// 	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
	// 		fmt.Fprintf(os.Stderr, "Error writing file: %s\n", err)
	// 	}
	//
	// 	fmt.Println("Initialized git directory")
	//
	// default:
	// 	fmt.Fprintf(os.Stderr, "Unknown command %s\n", command)
	// 	os.Exit(1)
	// }

}
