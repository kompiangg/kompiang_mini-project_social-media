package dto

import "time"

type CreateCommentRequest struct {
	Content     string  `json:"content"`
	PostID      *string `json:"post_id"`
	RecommentID *string `json:"recomment_id"`
	PictureURL  string  `json:"picture_url"`
	VideoURL    string  `json:"video_url"`
}

type CommentResponse struct {
	ID            string    `json:"comment_id"`
	CommentedBy   string    `json:"commented_by"`
	Content       string    `json:"content"`
	PostID        *string   `json:"post_id"`
	RecommentID   *string   `json:"recomment_id"`
	PictureURL    *string   `json:"picture_url"`
	VideoURL      *string   `json:"video_url"`
	ChildComments []string  `json:"child_comment"`
	CreatedAt     time.Time `json:"created_at"`
}

type CommentsResponse []CommentResponse
