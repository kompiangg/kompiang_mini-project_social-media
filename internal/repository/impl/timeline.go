package impl

import (
	"context"
	"log"

	"github.com/kompiang_mini-project_social-media/internal/models"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
)

func (r repository) GetTimeline(ctx context.Context, userCtx *dto.UserContext) (*dto.Timeline, error) {
	var posts []models.Post

	res := r.db.WithContext(ctx).
		Table("posts").
		Select("posts.*").
		Joins("inner join user_relations on user_relations.follow_by = posts.published_by").
		Find(&posts)
	err := res.Error
	if err != nil {
		log.Println("[GetTimeline Repository Error] While calling the query:", err.Error())
		return nil, err
	}

	var timeline dto.Timeline

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
		timeline = append(timeline, dto.Post{
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

	return &timeline, nil
}

func (r repository) CreateGeneralNotifications(ctx context.Context, req *dto.AdminNotificationRequest) error {
	var username *string
	if req.Username == "" {
		username = nil
	} else {
		username = &req.Username
	}

	notification := models.Notification{
		TargetNotification: username,
		Content:            req.Content,
	}

	res := r.db.WithContext(ctx).Create(&notification)
	err := res.Error
	if err != nil {
		log.Println("[CreateGeneralNotifications Repository Error] While calling the query:", err.Error())
		return err
	}

	return nil
}

func (r repository) GetNotificationsByUsername(ctx context.Context, username string) (*dto.NotificationsResponse, error) {
	var notifications []models.Notification

	res := r.db.WithContext(ctx).Find(&notifications).Where("target_notification = ? OR target_notification IS NULL", username)
	err := res.Error
	if err != nil {
		log.Println("[GetNotificationsByUsername Repository Error] While calling the query:", err.Error())
		return nil, err
	}

	var notificationsDto dto.NotificationsResponse

	for _, notification := range notifications {
		notificationDto := dto.NotificationResponse{
			Content:            notification.Content,
			TargetNotification: notification.TargetNotification,
			CreatedAt:          notification.CreatedAt,
		}
		notificationsDto = append(notificationsDto, notificationDto)
	}

	return &notificationsDto, nil
}
