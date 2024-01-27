package entity

import (
	"github.com/google/uuid"
	"time"
)

type Quote struct {
	ID   int
	Text string
}

type Post struct {
	ID         uuid.UUID `gorm:"column:id" json:"post_id"`
	Title      string    `gorm:"column:title" json:"title"`
	Content    string    `gorm:"column:content" json:"content"`
	TotalPoint int       `gorm:"column:total_point" json:"total_point"`
	Comment    []Comment `gorm:"Foreignkey:post_id;association_foreignkey:ID;" json:"comment"`
	CreatedAt  time.Time `gorm:"column:created_at" json:"-"`
}

type Comment struct {
	ID        uuid.UUID `gorm:"column:id" json:"comment_id"`
	PostID    string    `gorm:"column:post_id" json:"artikel_id"`
	Rate      int       `gorm:"column:rate" json:"rate"`
	Comment   string    `gorm:"column:comment" json:"comment"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
}
