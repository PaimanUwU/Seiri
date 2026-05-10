package app

import (
	"github.com/yourusername/seiri/core"
)

type Model struct {
	NotesDir string
	ActiveCategory int // 0: Inbox, 1: Projects, etc.
	Cursor int // Which note is highlighted in the active column
	
	// Lists of notes for each column
	Inbox     []core.Note
	Projects  []core.Note
	Areas     []core.Note
	Resources []core.Note
	
	Loaded bool
	Err error
}

func NewModel(notesDir string) Model {
	return Model{
		NotesDir: notesDir,
		ActiveCategory: 0,
		Cursor: 0,
	}
}
