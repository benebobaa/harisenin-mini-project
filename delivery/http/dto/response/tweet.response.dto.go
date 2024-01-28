package response

import (
	"fmt"
	"github.com/benebobaa/harisenin-mini-project/domain/entity"
	"github.com/google/uuid"
	"time"
)

type TweetResponse struct {
	ID          uuid.UUID        `json:"post_id"`
	User        entity.User      `json:"user"`
	Title       string           ` json:"title"`
	Content     string           `json:"content"`
	Comment     []entity.Comment `json:"comment"`
	TotalPoints int              `json:"total_points"`
	CreatedAt   time.Time        `json:"-"`
}

func NewTweetResponses(tweets []*entity.Post) []TweetResponse {
	var tweetResponses []TweetResponse
	var totalPoints int
	for _, tweet := range tweets {
		fmt.Println("user tweet", tweet.User.Username)

		for _, comment := range tweet.Comment {
			totalPoints += comment.Rate

			fmt.Println("user comment", comment.User.Username)
		}

		tweetResponses = append(tweetResponses, TweetResponse{
			ID:          tweet.ID,
			User:        tweet.User,
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

type ResponseErrorDTO struct {
	StatusCode int `json:"status_code"`
	Error      any `json:"error"`
}
