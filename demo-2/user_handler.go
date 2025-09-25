package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	// swagger 相关
	_ "github.com/xuepp/gindemo/demo-2/docs" // 这一行要导入你生成的 docs 包
)

type UserHandler struct {
	service *UserService
}

func NewUserHandler(s *UserService) *UserHandler {
	return &UserHandler{service: s}
}

// RegisterRoutes 注册路由
func (h *UserHandler) RegisterRoutes(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/users", h.CreateUser)
	r.GET("/users", h.ListUsers)
}

// CreateUser godoc
// @Summary 创建用户
// @Param user body User true "用户信息"
// @Success 200 {object} User
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

// ListUsers godoc
// @Summary 用户列表
// @Success 200 {array} User
// @Router /users [get]
func (h *UserHandler) ListUsers(c *gin.Context) {
	users, err := h.service.ListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
