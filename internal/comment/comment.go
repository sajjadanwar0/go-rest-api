package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment by id")
	ErrNotImplemented  = errors.New("not implemented")
)

// Comment - a representation of the comment structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all the methods that our service needs in order to operate
type Store interface {
	GetComment(ctx context.Context, string2 string) (Comment, error)
	PostComment(ctx context.Context, comment Comment) (Comment, error)
	DeleteComment(ctx context.Context, id string) error
	UpdateComment(ctx context.Context, id string, comment Comment) (Comment, error)
}

// Service - is the struct on which all our logic will be built on the top of
type Service struct {
	Store Store
}

// NewService - return a pointer to new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("Retrieving a comment")

	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		fmt.Println(err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, ID string, updatedComment Comment) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, updatedComment)
	if err != nil {
		fmt.Println("error updating comment")
		return Comment{}, err
	}
	return cmt, nil
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return s.Store.DeleteComment(ctx, id)
}

func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	insertedCmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return insertedCmt, nil
}
