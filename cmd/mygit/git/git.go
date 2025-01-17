package git

import (
	"bytes"
	"compress/zlib"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type Git struct{}

func NewGit() *Git {
	return &Git{}
}

func (g *Git) createDotFiles() error {
	for _, dir := range []string{".git", ".git/objects", ".git/refs"} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("Unable to create directory: %s", err)
		}
	}
	return nil
}

func (g *Git) initHEADFile() error {
	headFileContents := []byte("ref: refs/heads/master\n")
	if err := os.WriteFile(".git/HEAD", headFileContents, 0644); err != nil {
		return fmt.Errorf("Unable to write file: %s", err)
	}
	return nil
}

func (g *Git) Init() error {
	if err := g.createDotFiles(); err != nil {
		return err
	}
	return g.initHEADFile()
}

func (g *Git) CatFile(sha string) error {
	path := filepath.Join(".git", "objects", sha[:2], sha[2:])
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	r, err := zlib.NewReader(file)
	if err != nil {
		return err
	}
	defer r.Close()

	contentBuffer, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	parts := bytes.Split(contentBuffer, []byte("\x00"))
	newContent := string(parts[1])
	fmt.Print(newContent)

	return nil
}

func (g *Git) Run() error {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)
	catFileCmd := flag.NewFlagSet("cat-file", flag.ExitOnError)
	catFilePrettyPrint := catFileCmd.String("p", "", "Pretty print blob object content")

	if len(os.Args) < 2 {
		return fmt.Errorf("Expected 'init' or 'cat-file' subcommands")
	}

	switch os.Args[1] {
	case "init":
		initCmd.Parse(os.Args[2:])
		if err := g.Init(); err != nil {
			return err
		}
		fmt.Println("Initialized git repository")

	case "cat-file":
		if len(os.Args) < 3 {
			return fmt.Errorf("Expected -p flag with blob object")
		}
		catFileCmd.Parse(os.Args[2:])
		sha := *catFilePrettyPrint
		if err := g.CatFile(sha); err != nil {
			fmt.Println(catFilePrettyPrint)
			return err
		}

	default:
		return fmt.Errorf("Expected 'init' or 'cat-file' subcommands")
	}

	return nil
}
