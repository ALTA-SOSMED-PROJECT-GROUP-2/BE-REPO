package handler

import (
	"BE-REPO/features/comment/model"
	"BE-REPO/features/comment/service"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

type CommentHandler struct {
	cs *service.CommentService
}

func NewCommentHandler(cs *service.CommentService) *CommentHandler {
	return &CommentHandler{cs: cs}
}

func (ch *CommentHandler) CreateComment(c echo.Context) error {
	comment := new(model.Comment)
	if err := c.Bind(comment); err != nil {
		return err
	}
	if err := ch.cs.Create(comment); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, comment)
}

func (ch *CommentHandler) GetComment(c echo.Context) error {
	id := c.Param("id")
	commentID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	comment, err := ch.cs.Get(uint(commentID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "comment not found"})
	}
	return c.JSON(http.StatusOK, comment)
}

func (ch *CommentHandler) GetAllComments(c echo.Context) error {
	postingID := c.Param("posting_id")
	postingIDInt, err := strconv.Atoi(postingID)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid posting id"})
	}
	comments, err := ch.cs.GetAll(uint(postingIDInt))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "comments not found"})
	}
	return c.JSON(http.StatusOK, comments)
}

func (ch *CommentHandler) DeleteComment(c echo.Context) error {
	id := c.Param("id")
	commentID, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid id"})
	}
	err = ch.cs.Delete(uint(commentID))
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "comment not found"})
	}
	return c.JSON(http.StatusOK, map[string]string{"message": "comment deleted successfully"})
}
