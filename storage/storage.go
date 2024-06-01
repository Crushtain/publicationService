package storage

import "github.com/Crushtain/publicationService.git/graph/model"

type Storage interface {
	CreatePost(post *model.Post) error
	CreateComment(comment *model.Comment) error
	GetPost(id string) (*model.Post, error)
	CommentsByPostID(postID string) ([]*model.Comment, error)
}
