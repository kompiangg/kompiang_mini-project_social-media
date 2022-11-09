package handler

import (
	"log"
	"net/http"

	"github.com/kompiang_mini-project_social-media/internal/service"
	"github.com/kompiang_mini-project_social-media/pkg/dto"
	"github.com/kompiang_mini-project_social-media/pkg/errors"
	"github.com/kompiang_mini-project_social-media/pkg/utils/authutils"
	"github.com/kompiang_mini-project_social-media/pkg/utils/httputils"
	"github.com/labstack/echo/v4"
)

func CreatePost(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		// var req dto.CreatePostRequest
		// err := c.Bind(&req)
		// if err != nil {
		// 	return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
		// 		Err:    errors.ErrBadRequest,
		// 		Detail: []string{"Content type must be application/json"},
		// 	})
		// }
		userCtx := authutils.UserFromRequestContext(c)

		if userCtx == nil {
			log.Println("[HANDLER ERROR] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: errors.ErrInternalServer,
			})
		}

		repostID := c.FormValue("repost_id")
		req := dto.CreatePostRequest{
			Content:  c.FormValue("content"),
			RepostID: &repostID,
		}

		if *req.RepostID == "" {
			req.RepostID = nil
		}

		image, err := c.FormFile("image")
		if image != nil {
			if err != nil {
				log.Println("[HANDLER ERROR] Couldn't extract user account from context")
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: err,
				})
			}
		}

		video, err := c.FormFile("video")
		if video != nil {
			if err != nil {
				log.Println("[HANDLER ERROR] Couldn't extract user account from context")
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
		}

		var videoFileName *string
		if video != nil {
			videoFileName, err = httputils.HandleFileForm(video)
			if err != nil {
				return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
					Err: errors.ErrInternalServer,
				})
			}
		}

		post, err := service.CreatePost(c.Request().Context(), &req, userCtx.Username, imageFilename, videoFileName)
		if err != nil {
			log.Println("[HANDLER ERROR] While calling the create post service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Code: http.StatusCreated,
			Data: post,
		})
	}
}

func DeletePost(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		var req dto.DeletePostRequest
		err := c.Bind(&req)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: []string{"Content type must be application/json"},
			})
		}

		if req.PostID == "" {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: []string{"post_id field can't be empty"},
			})
		}

		userCtx := authutils.UserFromRequestContext(c)
		if userCtx == nil {
			log.Println("[HANDLER ERROR] Couldn't extract user account from context")
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		err = service.DeletePost(c.Request().Context(), &req)
		if err != nil {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: "Post Deleted!",
		})
	}
}

func GetPostById(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		if id == "" {
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err:    errors.ErrBadRequest,
				Detail: []string{"ID parameter need to be field"},
			})
		}

		post, err := service.GetPostByID(c.Request().Context(), id)
		if err != nil {
			log.Println("[HANDLER ERROR] While call get post by id service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: post,
		})
	}
}

func GetPostByUsername(service service.Service) echo.HandlerFunc {
	return func(c echo.Context) error {
		username := c.Param("username")

		posts, err := service.GetAllPostByUsername(c.Request().Context(), username)
		if err != nil {
			log.Println("[HANDLER ERROR] While call get post by username service:", err.Error())
			return httputils.WriteErrorResponse(c, httputils.ErrorResponseParams{
				Err: err,
			})
		}

		return httputils.WriteResponse(c, httputils.SuccessResponseParams{
			Data: posts,
		})
	}
}
