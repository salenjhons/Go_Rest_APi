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

type Comment struct {
	ID     string `json:"id"`
	Slug   string `json:"slug"`
	Body   string `json:"body"`
	Author string `json:"author"`
}

// CommentStore - defines the interface we need our comment storage
// layer to implement
type CommentStore interface {
	GetComment(context.Context, string) (Comment, error)
	PostComment(context.Context, Comment) (Comment, error)
	UpdateComment(context.Context, string, Comment) (Comment, error)
	DeleteComment(context.Context, string) error
	Ping(context.Context) error
}

// Service - the struct for our comment service
type Service struct {
	Store CommentStore
}

// NewService - returns a new comment service
func NewService(store CommentStore) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment - retrieves comments by their ID from the database
func (s *Service) GetComment(ctx context.Context, ID string) (Comment, error) {
	// calls store passing in the context
	cmt, err := s.Store.GetComment(ctx, ID)
	if err != nil {
		fmt.Errorf("an error occured fetching the comment: %s", err.Error())
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

// PostComment - adds a new comment to the database
func (s *Service) PostComment(ctx context.Context, cmt Comment) (Comment, error) {
	cmt, err := s.Store.PostComment(ctx, cmt)
	if err != nil {
		fmt.Errorf("an error occurred adding the comment: %s", err.Error())
	}
	return cmt, nil
}

// UpdateComment - updates a comment by ID with new comment info
func (s *Service) UpdateComment(
	ctx context.Context, ID string, newComment Comment,
) (Comment, error) {
	cmt, err := s.Store.UpdateComment(ctx, ID, newComment)
	if err != nil {
		fmt.Errorf("an error occurred updating the comment: %s", err.Error())
	}
	return cmt, nil
}

// DeleteComment - deletes a comment from the database by ID
func (s *Service) DeleteComment(ctx context.Context, ID string) error {
	return s.Store.DeleteComment(ctx, ID)
}