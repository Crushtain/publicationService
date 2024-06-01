package storage

import "github.com/Crushtain/publicationService.git/graph/model"

type InMemory struct {
	posts    []*model.Post
	comments []*model.Comment
}

func NewInMemory() *InMemory {
	return &InMemory{posts: []*model.Post{}}
}

func (m *InMemory) CreatePost(post *model.Post) error {
	m.posts = append(m.posts, post)
	return nil
}

func (m *InMemory) CreateComment(comment *model.Comment) error {
	m.comments = append(m.comments, comment)
	return nil
}

func (m *InMemory) GetPost(id string) (*model.Post, error) {
	var outPost *model.Post

	for _, post := range m.posts {
		if post.ID == id {
			outPost = post
		}
	}
	if outPost == nil {
		return &model.Post{}, nil
	}
	var comments []*model.Comment
	for _, comment := range m.comments {
		if comment.PostID == id {
			comments = append(comments, comment)
		}
	}
	outPost.Comments = comments

	return outPost, nil
}

func (m *InMemory) CommentsByPostID(postID string) ([]*model.Comment, error) {
	var comments []*model.Comment

	for _, c := range m.comments {
		if c.PostID == postID {
			comments = append(comments, c)
		}
	}
	if comments == nil {
		return []*model.Comment{}, nil
	}
	return comments, nil
}
