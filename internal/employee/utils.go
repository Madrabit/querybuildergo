package employee

import (
	"log"
	"os"
	"path/filepath"
)

func FindProjectRoot() string {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatalf("cannot get working dir: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			break // reached root of FS
		}
		dir = parent
	}

	log.Fatal("cannot find go.mod — are you inside a Go module?")
	return "" // unreachable
}
