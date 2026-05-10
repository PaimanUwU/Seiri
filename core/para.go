package core

import (
	"os"
	"path/filepath"
	"strings"
)

type Category string

const (
	Inbox     Category = "00_Inbox"
	Projects  Category = "10_Projects"
	Areas     Category = "20_Areas"
	Resources Category = "30_Resources"
	Archive   Category = "40_Archive"
)

type Note struct {
	Name     string
	Path     string
	Category Category
}

// ScanFolder looks into a specific PARA directory and returns the notes
func ScanFolder(basePath string, cat Category) ([]Note, error) {
	var notes []Note
	path := filepath.Join(basePath, string(cat))

	files, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".md") {
			notes = append(notes, Note{
				Name:     file.Name(),
				Path:     filepath.Join(path, file.Name()),
				Category: cat,
			})
		}
	}
	return notes, nil
}
