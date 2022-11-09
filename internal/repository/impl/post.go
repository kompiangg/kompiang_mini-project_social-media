package impl

import (
	"context"
	goerrors "errors"
	"log"

	"github.com/kompiang_mini-project_social-media/internal/models"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"gorm.io/gorm"
)

func (r repository) CreatePost(ctx context.Context, req *dto.CreatePostRequest, username string, id string) (*dto.Post, error) {
	post := models.Post{
		ID:          id,
		PublishedBy: username,
		RepostID:    req.RepostID,
		Content:     req.Content,
		PictureURL:  req.PictureURL,
		VideoURL:    req.VideoURL,
	}

	res := r.db.WithContext(ctx).Create(&post)
	err := res.Error
	if err != nil {
		log.Println("[CreatePost Repository Error] While calling the query:", err.Error())
		return nil, err
	}

	return &dto.Post{
		ID:          id,
		PublishedBy: username,
		Content:     post.Content,
		RepostID:    post.RepostID,
		PictureURL:  post.PictureURL,
		VideoURL:    post.VideoURL,
		Comments:    nil,
		CreatedAt:   post.CreatedAt,
	}, nil
}

func (r repository) DeletePost(ctx context.Context, req *dto.DeletePostRequest) error {
	post := models.Post{
		ID: req.PostID,
	}

	res := r.db.WithContext(ctx).First(&post)
	err := res.Error
	if err != nil {
		if goerrors.Is(err, gorm.ErrRecordNotFound) {
			return errors.ErrRecordNotFound
		}
		log.Println("[DeletePost Repository Error] While calling the query:", err.Error())
		return err
	}

	res = r.db.WithContext(ctx).Delete(&post)
	err = res.Error
	if err != nil {
		if goerrors.Is(err, gorm.ErrRecordNotFound) {
			return errors.ErrRecordNotFound
		}
		log.Println("[DeletePost Repository Error] While calling the query:", err.Error())
		return err
	}
	return nil
}

func (r repository) GetPostByID(ctx context.Context, id string) (*dto.Post, error) {
	post := models.Post{
		ID: id,
	}

	res := r.db.WithContext(ctx).Preload("Comments").First(&post)
	err := res.Error
	if err != nil {
		if goerrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrRecordNotFound
		}
		log.Println("[GetPostByID Repository Error] While calling the query:", err.Error())
		return nil, err
	}

	var commentsResponse dto.CommentsResponse
	for _, postComment := range post.Comments {
		var childCommentsID []string
		var comments []models.Comment

		res := r.db.WithContext(ctx).Preload("ChildComments").Where("comment_id = ?", postComment.ID).Find(&comments)
		err := res.Error
		if err != nil {
			log.Println("[GetChildComment Repository Error] While calling the query:", err.Error())
			return nil, err
		}
		for _, childComment := range comments {
			childCommentsID = append(childCommentsID, childComment.ID)
		}

		comment := dto.CommentResponse{
			ID:            postComment.ID,
			PostID:        postComment.PostID,
			CommentedBy:   postComment.CommentBy,
			Content:       postComment.Content,
			RecommentID:   nil,
			PictureURL:    &postComment.PictureURL,
			VideoURL:      &postComment.VideoURL,
			CreatedAt:     postComment.CreatedAt,
			ChildComments: childCommentsID,
		}
		commentsResponse = append(commentsResponse, comment)
	}

	return &dto.Post{
		ID:          post.ID,
		PublishedBy: post.PublishedBy,
		Content:     post.Content,
		RepostID:    post.RepostID,
		PictureURL:  post.PictureURL,
		VideoURL:    post.VideoURL,
		Comments:    &commentsResponse,
		CreatedAt:   post.CreatedAt,
	}, nil
}

func (r repository) GetAllPostByUsername(ctx context.Context, username string) (*dto.Posts, error) {
	var posts []models.Post

	res := r.db.WithContext(ctx).Where("published_by = ?", username).First(&models.Post{})
	err := res.Error
	if err != nil {
		if goerrors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.ErrRecordNotFound
		}
		log.Println("[DeletePost Repository Error] While calling the query:", err.Error())
		return nil, err
	}

	res = r.db.WithContext(ctx).Preload("Comments").Where("published_by = ?", username).Find(&posts)
	err = res.Error
	if err != nil {
		log.Println("[GetAllPostByUsername Repository Error] While calling the query:", err.Error())
		return nil, err
	}

	var postsDto dto.Posts

	for _, post := range posts {
		var comments dto.CommentsResponse
		for _, postComment := range post.Comments {
			var childCommentsID []string
			for _, childComment := range postComment.ChildComments {
				childCommentsID = append(childCommentsID, childComment.ID)
			}

			comment := dto.CommentResponse{
				ID:            postComment.ID,
				PostID:        postComment.PostID,
				CommentedBy:   postComment.CommentBy,
				Content:       postComment.Content,
				RecommentID:   nil,
				PictureURL:    &postComment.PictureURL,
				VideoURL:      &postComment.VideoURL,
				CreatedAt:     postComment.CreatedAt,
				ChildComments: childCommentsID,
			}
			comments = append(comments, comment)
		}
		postsDto = append(postsDto, dto.Post{
			ID:          post.ID,
			PublishedBy: post.PublishedBy,
			Content:     post.Content,
			RepostID:    post.RepostID,
			PictureURL:  post.PictureURL,
			VideoURL:    post.VideoURL,
			Comments:    &comments,
			CreatedAt:   post.CreatedAt,
		})
	}

	return &postsDto, nil
}
