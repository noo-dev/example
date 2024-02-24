package handler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/noo-dev/example/internal/category/service"
)

type CategoryHandler struct {
	service *service.CategoryService
}

func NewCategoryHandler(service *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{service: service}
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	Categorys, err := h.service.GetAllCategories()
	if err != nil {
		c.JSON(500, "something went wrong")
	}

	c.JSON(200, gin.H{
		"data": Categorys,
	})
}

func (h *CategoryHandler) GetOneCategory(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(400, "invalid id")
	}
	Categorys, err := h.service.GetCategoryById(id)
	if err != nil {
		c.JSON(500, "something went wrong")
	}

	c.JSON(200, gin.H{
		"data": Categorys,
	})
}
