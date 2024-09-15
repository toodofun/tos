package database

import (
	"github.com/MR5356/aurora/pkg/config"
	"github.com/MR5356/aurora/pkg/util/structutil"
	"testing"
)

var _ = config.New(config.WithDatabase("sqlite", ":memory:"))

type User struct {
	ID   string `gorm:"primary_key"`
	Name string

	BaseModel
}

func TestBaseMapper(t *testing.T) {
	NewDatabase(config.Current(config.WithDatabase("sqlite", ":memory:")))
	err := GetDB().AutoMigrate(&User{})
	if err != nil {
		t.Fatalf("Failed to initialize database: %v", err)
	}

	mapper := NewMapper(GetDB(), &User{})

	u1 := &User{
		ID:   "1",
		Name: "test",
	}

	err = mapper.Insert(u1)
	if err != nil {
		t.Fatalf("Failed to insert user: %v", err)
	}

	err = mapper.Insert(u1)
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	u, err := mapper.Detail(&User{ID: u1.ID})
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if u.Name != u1.Name {
		t.Fatalf("Expected %s, got %s", u1.Name, u.Name)
	}

	u, err = mapper.Detail(&User{ID: "2"})
	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	u1.Name = "test2"
	err = mapper.Update(&User{ID: u1.ID}, structutil.Struct2Map(u1))
	if err != nil {
		t.Fatalf("Failed to update user: %v", err)
	}

	u, err = mapper.Detail(&User{ID: u1.ID})
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}

	if u.Name != u1.Name {
		t.Fatalf("Expected %s, got %s", u1.Name, u.Name)
	}

	us, err := mapper.List(&User{})
	if err != nil {
		t.Fatalf("Failed to list users: %v", err)
	}

	if len(us) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(us))
	}

	us, err = mapper.List(&User{}, "updated_at desc")
	if err != nil {
		t.Fatalf("Failed to list users: %v", err)
	}

	if len(us) != 1 {
		t.Fatalf("Expected 1 user, got %d", len(us))
	}

	c, err := mapper.Count(&User{})
	if err != nil {
		t.Fatalf("Failed to count users: %v", err)
	}

	if c != 1 {
		t.Fatalf("Expected 1 user, got %d", c)
	}

	p, err := mapper.Page(&User{}, 1, 1)
	if err != nil {
		t.Fatalf("Failed to page users: %v", err)
	}

	if p.Total != 1 {
		t.Fatalf("Expected 1 user, got %d", p.Total)
	}

	err = mapper.Delete(u1)
	if err != nil {
		t.Fatalf("Failed to delete user: %v", err)
	}
}
