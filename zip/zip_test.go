package zip

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestExampleArchiveFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test_zip")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	outFilePath := filepath.Join(tmpDir, "foo.zip")

	progress := func(archivePath string) {
		fmt.Println(archivePath)
	}

	err = ArchiveFile("testdata/foo/", outFilePath, progress)
	if err != nil {
		t.Fatal(err)
	}

	// Output:
	// foo/bar
	// foo/baz/aaa
}

func TestExampleUnarchiveFile(t *testing.T) {
	tmpDir, err := os.MkdirTemp("", "test_zip")
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	progress := func(archivePath string) {
		fmt.Println(archivePath)
	}

	err = UnarchiveFile("testdata/foo.zip", tmpDir, progress)
	if err != nil {
		t.Fatal(err)
	}

	// Output:
	// foo/bar
	// foo/baz/aaa
}
