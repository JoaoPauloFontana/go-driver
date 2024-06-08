package files

import (
	"errors"
	"time"
)

var (
	ErrNameRequired  = errors.New("name is required and can't be blank")
	ErrTypeRequired  = errors.New("type is required and can't be blank")
	ErrPathRequired  = errors.New("path is required and can't be blank")
	ErrOwnerRequired = errors.New("owner_id is required")
)

type File struct {
	ID         int64     `json:"id"`
	FolderID   int64     `json:"-"`
	OwnerID    int64     `json:"owner_id"`
	Name       string    `json:"name"`
	Type       string    `json:"type"`
	Path       string    `json:"-"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
}

func (f *File) Validate() error {
	if f.Name == "" {
		return ErrNameRequired
	}

	if f.Type == "" {
		return ErrNameRequired
	}

	if f.Path == "" {
		return ErrNameRequired
	}

	if f.OwnerID == 0 {
		return ErrNameRequired
	}

	return nil
}
