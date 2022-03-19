package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not Implemented")
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
// interact with the store or repository layer
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
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

// GetComment - returns a comment by its ID
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

// PostComment - creates a new comment
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}

// UpdateComment - updates a comment
func (s *Service) UpdateComment(
	ctx context.Context,
	ID string,
	updatedCmt Comment,
) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, updatedCmt)
	if err != nil {
		fmt.Println("error updating comment")
		return Comment{}, err
	}
	return cmt, nil
}

// DeleteComment - deletes a comment by its id
func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}
