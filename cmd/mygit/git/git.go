package git

import (
	// "bytes"
	"bytes"
	"compress/zlib"
	"context"

	// "encoding/base64"
	"fmt"
	"io"
	"os"

	"github.com/urfave/cli/v3"
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

func (g *Git) createDotFiles() error {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			err := fmt.Errorf("Unable to create directory: %s\n", err)
			return err
		}
	}
	return nil
}

func (g *Git) initHEADFile() error {
	headFileContents := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
		err := fmt.Errorf("Unable to write file: %s\n", err)
		return err
	}
	return nil
}

func (g *Git) Init() *cli.Command {

	return &cli.Command{
		Name:  "init",
		Usage: "Initilizes a git repository",
		Action: func(ctx context.Context, cmd *cli.Command) error {

			if err := g.createDotFiles(); err != nil {
				return err
			}

			if err := g.initHEADFile(); err != nil {
				return err
			}
			return nil
		},
	}
}

func (g *Git) OpenFile() *cli.Command {
	return &cli.Command{
		Name:  "open-file",
		Usage: "Opens file",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args()
			filePath := args.Get(0)
			fileData, _ := os.ReadFile(filePath)
			fmt.Println(string(fileData))
			return nil
		},
	}
}

func (g *Git) CatFile() *cli.Command {
	return &cli.Command{
		Name:  "cat-file",
		Usage: "Provide content or type and size information for repository objects",
		Action: func(ctx context.Context, cmd *cli.Command) error {
			args := cmd.Args()
			if args.Len() < 1 || args.Len() > 1 {
				err := fmt.Errorf("Wrong number of arguments to input. Requires 1 argument to cat-file command")
				return err
			}

			sha := args.Get(0)
			path := fmt.Sprintf(".git/objects/%v/%v", sha[0:2], sha[2:])
			file, _ := os.Open(path)

			r, err := zlib.NewReader(io.Reader(file))
			if err != nil {
				return err
			}

			contentBuffer, err := io.ReadAll(r)
			if err != nil {
				return err
			}

			parts := bytes.Split(contentBuffer, []byte("\x00"))
			newContent := string(parts[1])
			fmt.Println(newContent)

			io.Copy(os.Stdout, r)
			r.Close()
			file.Close()

			return nil
		},
	}
}
