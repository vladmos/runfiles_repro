package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestRunfilesAvailabilty(t *testing.T) {
	testdata := filepath.Join(filepath.FromSlash(os.Getenv("TEST_SRCDIR")), os.Getenv("TEST_WORKSPACE"))

	files := []string{"foo.txt", "baz.txt", "bar/bar.txt", "bar/foo.txt"}
	for _, name := range files {
		path := filepath.Join(testdata, filepath.FromSlash(name))
		if _, err := os.ReadFile(path); err != nil {
			t.Fatalf("%q doesn't exist: %s", err)
		}
	}
}
