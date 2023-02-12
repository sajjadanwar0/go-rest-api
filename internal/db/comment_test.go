//go:build integration
// +build integration

package db

import (
	"context"
	"github.com/sajjadanwar0/go-rest-api/internal/comment"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestComment(t *testing.T) {
	t.Run("test create comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "slug",
			Author: "me",
			Body:   "body",
		})
		assert.NoError(t, err)
		newCmt, err := db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		assert.Equal(t, "slug", newCmt.Slug)
	})

	t.Run("test delete comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmt, err := db.PostComment(context.Background(), comment.Comment{
			Slug:   "new-slug",
			Author: "new-me",
			Body:   "new-body",
		})
		assert.NoError(t, err)
		err = db.DeleteComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
		_, err = db.GetComment(context.Background(), cmt.ID)
		assert.NoError(t, err)
	})
}
