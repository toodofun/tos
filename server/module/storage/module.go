package storage

import (
	"github.com/MR5356/tos/persistence/database"
	"github.com/google/uuid"
)

type Storage struct {
	LocationType string `json:"locationType" yaml:"locationType" default:"local"`
	Args         string `json:"args" yaml:"args"`

	database.BaseModel
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
