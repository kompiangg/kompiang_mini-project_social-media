package impl

import (
	"context"
	"log"

	"github.com/kompiang_mini-project_social-media/internal/models"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
)

func (r repository) CreateComment(ctx context.Context, req *dto.CreateCommentRequest, username string, id string) (*dto.CommentResponse, error) {
	comment := models.Comment{
		ID:         id,
		CommentBy:  username,
		Content:    req.Content,
		PostID:     req.PostID,
		CommentID:  req.RecommentID,
		PictureURL: req.PictureURL,
		VideoURL:   req.VideoURL,
	}

	res := r.db.WithContext(ctx).Create(&comment)
	err := res.Error
	if err != nil {
		log.Println("[CreateComment Repository Error] While calling the query:", err.Error())
		return nil, err
	}

	return &dto.CommentResponse{
		ID:            comment.ID,
		CommentedBy:   comment.CommentBy,
		Content:       comment.Content,
		PostID:        comment.PostID,
		RecommentID:   comment.CommentID,
		PictureURL:    &comment.PictureURL,
		VideoURL:      &comment.VideoURL,
		ChildComments: []string{},
		CreatedAt:     comment.CreatedAt,
	}, nil
}

func (r repository) GetChildComment(ctx context.Context, parentID string) (*dto.CommentsResponse, error) {
	var comments []models.Comment

	res := r.db.WithContext(ctx).Preload("ChildComments").Where("comment_id = ?", parentID).Find(&comments)
	err := res.Error
	if err != nil {
		log.Println("[GetChildComment Repository Error] While calling the query:", err.Error())
		return nil, err
	}

	var resComments dto.CommentsResponse
	for _, comment := range comments {
		var childID []string
		for _, child := range comment.ChildComments {
			childID = append(childID, child.ID)
		}

		resComment := dto.CommentResponse{
			ID:            comment.ID,
			CommentedBy:   comment.CommentBy,
			Content:       comment.Content,
			PostID:        nil,
			RecommentID:   comment.CommentID,
			PictureURL:    &comment.PictureURL,
			VideoURL:      &comment.VideoURL,
			ChildComments: childID,
			CreatedAt:     comment.CreatedAt,
		}
		resComments = append(resComments, resComment)
	}

	return &resComments, nil
}
