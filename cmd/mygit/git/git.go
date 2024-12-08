package git

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

//TODO: create object types
// type blob struct {
//
// }

type Git struct {
	fileType string
}

func NewGit() *Git {
	return &Git{}
}

func Action(cCtx *cli.Context) error {

	g := NewGit()

	if cCtx.Args().Len() < 1 {
		err := fmt.Errorf("usage: mygit <command> [<args>...]\n")
		return err
	}

	//
	// if len(os.Args) < 2 {
	// 	fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
	// 	os.Exit(1)
	// }
	//
	switch command := os.Args[1]; command {
	case "init":
		if err := g.initRepo(); err != nil {
			err := fmt.Errorf("Error initializing git directory: %s\n", err)
			return err
		}

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

	default:
		err := fmt.Errorf("Unknown command %s\n", command)
		g.printUsage()
		return err
	}

	return nil
}

func (g *Git) initRepo() error {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			err := fmt.Errorf("Unable to create directory: %s\n", err)
			return err
		}
	}

	headFileContents := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
		err := fmt.Errorf("Unable to write file: %s\n", err)
		return err
	}

	fmt.Println("Initialized git directory")
	return nil

}

func (g *Git) printUsage() {
	fmt.Fprintf(os.Stderr, "usage: mygit <command> [<args>...]\n")
}
