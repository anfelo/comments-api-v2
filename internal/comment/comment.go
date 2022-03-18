package comment

import (
	"context"
	"fmt"
)

// Comment - a representation of the commnet
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - the main interface that describes how to
// interact with the store
type Store interface {
	GetComment(context.Context, string) (Comment, error)
}

// Service - is the struct on which all our logic
// will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a pointer to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment - return a comment by its ID
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, err
	}
	return cmt, nil
}
