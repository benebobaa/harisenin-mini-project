package request

import (
	"github.com/benebobaa/harisenin-mini-project/domain/entity"
	"github.com/google/uuid"
	"time"
)

type TweetRequestDTO struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (this TweetRequestDTO) ToTweetEntity() entity.Post {
	return entity.Post{
		ID:        uuid.New(),
		Title:     this.Title,
		Content:   this.Content,
		CreatedAt: time.Now(),
	}

}

type CommentRequestDTO struct {
	PostID  string `json:"post_id" validate:"required"`
	Comment string `json:"comment" validate:"required"`
	Rate    int    `json:"rate" validate:"required,gte=1,number"`
}

func (this CommentRequestDTO) ToCommentEntity() entity.Comment {
	return entity.Comment{
		ID:        uuid.New(),
		PostID:    this.PostID,
		Comment:   this.Comment,
		Rate:      this.Rate,
		CreatedAt: time.Now(),
	}
}
