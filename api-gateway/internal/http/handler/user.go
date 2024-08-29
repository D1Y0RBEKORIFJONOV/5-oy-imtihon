package handler

import (
	_ "apigateway/internal/app/docs"
	"apigateway/internal/entity"
	userusecase "apigateway/internal/usecase/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserServer struct {
	user userusecase.User
}

func NewUserServer(user userusecase.User) *UserServer {
	return &UserServer{
		user: user,
	}
}

// @title Artisan Connect
// @version 1.0
// @description This is a sample server for a restaurant reservation system.
// @host localhost:9006
// @BasePath        /
// @schemes         https
// @securityDefinitions.apiKey ApiKeyAuth
// @in              header
// @name            Authorization

// Register godoc
// @Summary Register
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param body body entity.CreateUserReq true "User registration information"
// @Security ApiKeyAuth
// @Success 201 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/register [post]
func (u *UserServer) Register(c *gin.Context) {
	var req entity.CreateUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := u.user.RegisterUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// VerifyUser godoc
// @Summary VerifyUser
// @Description Confirm the code sent to the email
// @Tags auth
// @Accept json
// @Produce json
// @Param body body entity.VerifyUserReq true "User verification information"
// @Security ApiKeyAuth
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/verify [post]
// @Security BearerAuth
func (u *UserServer) VerifyUser(c *gin.Context) {
	var req entity.VerifyUserReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message, err := u.user.VerifyUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// Login godoc
// @Summary Login
// @Description Authenticate a user and return a token
// @Tags auth
// @Accept json
// @Produce json
// @Param body body entity.LoginReq true "User login information"
// @Security ApiKeyAuth
// @Success 200 {object} entity.LoginRes
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/login [post]
// @Security BearerAuth
func (u *UserServer) Login(c *gin.Context) {
	var req entity.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := u.user.Login(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, token)
}

// UpdateUser godoc
// @Summary UpdateUser
// @Description Authenticate a user and return a token
// @Tags user
// @Accept json
// @Produce json
// @Param body body entity.UpdateUserReq true "User login information"
// @Security ApiKeyAuth
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/update [put]
// @Security BearerAuth
func (u *UserServer) UpdateUser(c *gin.Context) {
	var req entity.UpdateUserReq
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	message, err := u.user.UpdateUser(c.Request.Context(), entity.UpdateUserReq{
		UserID:    id.(string),
		FirstName: req.FirstName,
		LastName:  req.LastName,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// UpdatePassword godoc
// @Summary UpdatePassword
// @Description Authenticate a user and return a token
// @Tags user
// @Accept json
// @Produce json
// @Param body body entity.UpdatePasswordReq true "User login information"
// @Security ApiKeyAuth
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/update/password [put]
// @Security BearerAuth
func (u *UserServer) UpdatePassword(c *gin.Context) {
	var req entity.UpdatePasswordReq
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserID = id.(string)
	message, err := u.user.UpdatePassword(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// UpdateEmail godoc
// @Summary UpdateEmail
// @Description Authenticate a user and return a token
// @Tags user
// @Accept json
// @Produce json
// @Param body body entity.UpdateEmailReq true "User login information"
// @Security ApiKeyAuth
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/update/email [put]
// @Security BearerAuth
func (u *UserServer) UpdateEmail(c *gin.Context) {
	var req entity.UpdateEmailReq
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserID = id.(string)
	message, err := u.user.UpdateEmail(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": message})
}

// DeleteUser godoc
// @Summary DeleteUser
// @Description Authenticate a user and return a token
// @Tags user
// @Accept json
// @Produce json
// @Param body body entity.DeleteUserReq true "User login information"
// @Security ApiKeyAuth
// @Success 200 {object} entity.StatusMessage
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /user/delete  [delete]
// @Security BearerAuth
func (u *UserServer) DeleteUser(c *gin.Context) {
	var req entity.DeleteUserReq
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	req.UserID = id.(string)
	message, err := u.user.DeleteUser(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// GetUser godoc
// @Summary GetUser
// @Description Retrieve user information by field and value
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Success 200 {object} entity.User // Adjust this to match your user object structure
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /user [get]
func (u *UserServer) GetUser(c *gin.Context) {
	email, ok := c.Get("email")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id not found"})
		return
	}
	user, err := u.user.GetUser(c.Request.Context(), email.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// GetAllUser godoc
// @Summary Get all users
// @Description Retrieve user information by field and value with pagination
// @Tags user
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Security BearerAuth
// @Param field query string false "Field to filter by" // Add description if field is required or optional
// @Param value query string false "Value to filter by" // Add description if value is required or optional
// @Param page query int false "Page number for pagination" // Add description if page is required or optional
// @Param limit query int false "Number of items per page" // Add description if limit is required or optional
// @Success 200 {object} entity.GetAllUserRes "List of users" // Adjust this to match your user object structure
// @Failure 400 {object} string "Bad request"
// @Failure 500 {object} string "Internal server error"
// @Router /user/all [get]
func (u *UserServer) GetAllUser(c *gin.Context) {
	page := c.Query("page")
	limit := c.Query("limit")
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		pageInt = 0
	}
	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		limitInt = 0
	}
	GetAllUserReq := entity.GetAllUserReq{
		Field: c.Query("field"),
		Value: c.Query("value"),
		Page:  int64(pageInt),
		Limit: int64(limitInt),
	}
	users, err := u.user.GetAllUsers(c.Request.Context(), &GetAllUserReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}
