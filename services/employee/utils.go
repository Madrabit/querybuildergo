package employee

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"
)

func StringOrEmpty(s sql.NullString) string {
	if s.Valid {
		return s.String
	}
	return ""
}

func TimeOrEmpty(t sql.NullTime, layout string) string {
	if t.Valid {
		return t.Time.Format(layout)
	}
	return ""
}

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
