package dto

import "time"

type CreatePostRequest struct {
	RepostID   *string `json:"repost_id"`
	Content    string  `json:"content"`
	PictureURL *string `json:"picture_url"`
	VideoURL   *string `json:"video_url"`
}

type DeletePostRequest struct {
	PostID string `json:"post_id"`
}

type Post struct {
	ID          string            `json:"id"`
	PublishedBy string            `json:"published_by"`
	Content     string            `json:"content"`
	RepostID    *string           `json:"repost_id"`
	PictureURL  *string           `json:"picture_url"`
	VideoURL    *string           `json:"video_url"`
	Comments    *CommentsResponse `json:"comments"`
	CreatedAt   time.Time         `json:"created_at"`
}

type Posts []Post
