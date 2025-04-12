package git

import (
	"bytes"
	"compress/zlib"
	"crypto/sha1"
	"encoding/hex"
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

func (g *Git) PrintUsage() {
	fmt.Println("Usage: mygit <command> [<args>]")
	fmt.Println("\nAvailable commands:")
	fmt.Println("  init                      Initialize a new git repository")
	fmt.Println("  cat-file  -p <blob>       Pretty-print the contents of a git object")
	fmt.Println("  hash-object <file> [-w]   Create a git object from a file. Also writes to .git/objects if -w is provided")
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

func (g *Git) HashObject(file string) error {
	content, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Unable to read file: %s", file)
	}

	stats, err := os.Stat(file)
	if err != nil {
		return fmt.Errorf("Unable to fetch stats on file: %s. %s", file, err)
	}

	sha := sha1.New()
	header := fmt.Sprintf("blob %d\x00", stats.Size())
	sha.Write([]byte(header))
	sha.Write(content)
	hashBytes := sha.Sum(nil)
	hash := hex.EncodeToString(hashBytes)

	fmt.Println(hash)

	path := filepath.Join(".git", "objects", hash[:2], hash[2:])
	parent := filepath.Join(".git", "objects", hash[:2])
	err = os.MkdirAll(parent, os.ModeDir|502)
	if err != nil {
		return fmt.Errorf("Error creating directory: %s. %s", path, err)
	}

	data := []byte(header)
	data = append(data, content...)
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	blobObject, err := os.OpenFile(path, os.O_RDONLY|os.O_CREATE, 0666)
	defer blobObject.Close()

	if _, err = w.Write(data); err != nil {
		return fmt.Errorf("Unable to write to zlib writer: %s", err)
	}
	defer w.Close()

	_, err = io.Copy(blobObject, &b)
	if err != nil {
		return fmt.Errorf("Unable to write file: %s", err)
	}
	fmt.Println(b.Bytes())

	return nil
}

func (g *Git) Run() error {
	initCmd := flag.NewFlagSet("init", flag.ExitOnError)

	catFileCmd := flag.NewFlagSet("cat-file", flag.ExitOnError)
	catFilePrettyPrint := catFileCmd.String("p", "", "Pretty print blob object content")

	hashObjectCmd := flag.NewFlagSet("hash-object", flag.ExitOnError)
	hashObjectWriteObject := hashObjectCmd.String("w", "", "Write blob to .git/objects")

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
			return err
		}

	case "hash-object":

		if len(os.Args) < 3 {
			return fmt.Errorf("Expected -w flag with file")
		}
		err := hashObjectCmd.Parse(os.Args[2:])
		if err != nil {
			return fmt.Errorf("hash-object command misuse")
		}
		file := *hashObjectWriteObject

		if err := g.HashObject(file); err != nil {
			return err
		}

	default:
		return fmt.Errorf("Expected subcommands")
	}

	return nil
}
