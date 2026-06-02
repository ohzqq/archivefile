package zip

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestExampleZip(t *testing.T) {
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

	err = Zip("testdata/foo/", outFilePath, progress)
	if err != nil {
		t.Fatal(err)
	}

	// Output:
	// foo/bar
	// foo/baz/aaa
}

func ExampleZip() {
	tmpDir, err := os.MkdirTemp("", "test_zip")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	outFilePath := filepath.Join(tmpDir, "foo.zip")

	progress := func(archivePath string) {
		fmt.Println(archivePath)
	}

	err = Zip("testdata/foo/", outFilePath, progress)
	if err != nil {
		panic(err)
	}

	// Output:
	// foo/bar
	// foo/baz/aaa
}

func TestExampleUnzip(t *testing.T) {
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

	err = Unzip("testdata/foo.zip", tmpDir, progress)
	if err != nil {
		t.Fatal(err)
	}

	// Output:
	// foo/bar
	// foo/baz/aaa
}

func ExampleUnzip() {
	tmpDir, err := os.MkdirTemp("", "test_zip")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = os.RemoveAll(tmpDir)
	}()

	progress := func(archivePath string) {
		fmt.Println(archivePath)
	}

	err = Unzip("testdata/foo.zip", tmpDir, progress)
	if err != nil {
		panic(err)
	}

	// Output:
	// foo/bar
	// foo/baz/aaa
}
