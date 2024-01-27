package response

import (
	"github.com/benebobaa/harisenin-mini-project/domain/entity"
	"github.com/google/uuid"
	"time"
)

type TweetResponse struct {
	ID          uuid.UUID        `json:"post_id"`
	Title       string           ` json:"title"`
	Content     string           `json:"content"`
	Comment     []entity.Comment `json:"comment"`
	TotalPoints int              `json:"total_points"`
	CreatedAt   time.Time        `json:"-"`
}

func NewTweetResponse(tweets []*entity.Post) []TweetResponse {
	var tweetResponses []TweetResponse
	var totalPoints int
	for _, tweet := range tweets {

		for _, comment := range tweet.Comment {
			totalPoints += comment.Rate
		}
		tweetResponses = append(tweetResponses, TweetResponse{
			ID:          tweet.ID,
			Title:       tweet.Title,
			Content:     tweet.Content,
			Comment:     tweet.Comment,
			TotalPoints: totalPoints,
			CreatedAt:   tweet.CreatedAt,
		})

		totalPoints = 0
	}
	return tweetResponses
}
