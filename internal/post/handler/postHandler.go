package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/noo-dev/example/internal/post/service"
)

type PostHandler struct {
	service *service.PostService
}

func NewPostHandler(service *service.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) GetAllPosts(c *gin.Context) {
	posts, err := h.service.GetAllPosts()
	if err != nil {
		c.JSON(500, "something went wrong")
	}

	c.JSON(200, gin.H{
		"data": posts,
	})
}

func (h *PostHandler) GetOnePost(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, "invalid id")
	}
	posts, err := h.service.GetPostById(id)
	if err != nil {
		c.JSON(500, "something went wrong")
	}

	c.JSON(200, gin.H{
		"data": posts,
	})
}
