package DB

import (
	"github.com/jinzhu/gorm"
)

type NoteListDB struct {
	gorm.Model
	UserID string
	NoteListStatus bool
	NoteListLevel  string
	NoteListTitle  string
	Timestamp 	   string
	ModTime        string
	NoteListShare  bool
	NoteListData   string
	BackImage      string
}

