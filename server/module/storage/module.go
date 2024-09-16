package storage

import (
	"github.com/MR5356/tos/persistence/database"
	"github.com/google/uuid"
	"time"
)

type Storage struct {
	Title        string `json:"title" yaml:"title"`
	LocationType string `json:"locationType" yaml:"locationType" default:"local"`
	Args         string `json:"args" yaml:"args"`

	database.BaseModel
}

type FileInfo struct {
	ID      uuid.UUID `json:"id"`
	Name    string    `json:"name"`
	Path    string    `json:"path"`
	Size    int64     `json:"size"`
	IsDir   bool      `json:"isDir"`
	ModTime time.Time `json:"modTime"`
}

func (s *Storage) WithID(id uuid.UUID) *Storage {
	s.ID = id
	return s
}

func NewStorageWithID(id uuid.UUID) *Storage {
	return &Storage{
		BaseModel: database.BaseModel{
			ID: id,
		},
	}
}
