package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{service: s}
}

func (h *UserHandler) RegisterRoutes(e *gin.Engine) {
	e.GET("/users", func(c *gin.Context) {
		users := h.service.ListUsers()
		c.JSON(http.StatusOK, gin.H{"users": users})
	})
}
