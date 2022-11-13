package handler

import (
	"log"
	"net/http"
	"os"

	"github.com/kompiang_mini-project_social-media/internal/service"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/authutils"
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func CreateComment(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		userCtx := authutils.UserFromRequestContext(c)
		if userCtx == nil {
			log.Println("[HANDLER ERROR] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: errors.ErrInternalServer,
			})
		}

		postID := c.FormValue("post_id")
		recommentID := c.FormValue("recomment_id")
		if postID == "" && recommentID == "" {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: []string{"One of post_id or recomment_id shouldnt be empty"},
			})
		}

		var postIDPtr *string = &postID
		var recommentIDPtr *string = &recommentID

		if postID == "" {
			postIDPtr = nil
		} else if recommentID == "" {
			recommentIDPtr = nil
		}

		req := dto.CreateCommentRequest{
			Content:     c.FormValue("content"),
			PostID:      postIDPtr,
			RecommentID: recommentIDPtr,
		}

		image, err := c.FormFile("image")
		if image != nil {
			if err != nil {
				log.Println("[HANDLER ERROR] while get image form file")
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: err,
				})
			}
		}

		video, err := c.FormFile("video")
		if video != nil {
			if err != nil {
				log.Println("[HANDLER ERROR] while get video form file")
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: err,
				})
			}
		}

		if req.Content == "" {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: []string{"Content field can't be empty"},
			})
		}

		var imageFilename *string
		if image != nil {
			imageFilename, err = httputils.HandleFileForm(image)
			if err != nil {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: errors.ErrInternalServer,
				})
			}
			defer os.Remove(*imageFilename)
		}

		var videoFileName *string
		if video != nil {
			videoFileName, err = httputils.HandleFileForm(video)
			if err != nil {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: errors.ErrInternalServer,
				})
			}
			defer os.Remove(*videoFileName)
		}

		comment, err := service.CreateComment(c.Request().Context(), &req, userCtx.Username, imageFilename, videoFileName)
		if err != nil {
			log.Println("[HANDLER ERROR] Couldn't extract user account from context:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Code: http.StatusCreated,
			Data: comment,
		})
	}
}

func GetChildCommentByCommentID(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		parentCommentID := c.Param("parent_comment_id")

		childComment, err := service.GetChildComment(c.Request().Context(), parentCommentID)
		if err != nil {
			log.Println("[HANDLER ERROR] While call get child comment service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: childComment,
		})
	}
}
