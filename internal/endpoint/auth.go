package endpoint

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/lavatee/tracker_backend/internal/model"
)

type SignUpInput struct {
	TelegramUsername string `json:"telegram_username"`
	FirstName        string `json:"first_name"`
	LastName         string `json:"last_name"`
	Grade            int    `json:"grade"`
	ClassLetter      string `json:"class_letter"`
	Password         string `json:"password"`
	ByReferral       string `json:"by_referral"`
}

func (e *Endpoint) SignUp(c *gin.Context) {
	var input SignUpInput
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if input.ByReferral == "" {
		input.ByReferral = "none"
	}
	newUser := model.User{
		TelegramUsername: input.TelegramUsername,
		FirstName:        input.FirstName,
		LastName:         input.LastName,
		Grade:            input.Grade,
		ClassLetter:      input.ClassLetter,
		PasswordHash:     input.Password,
		ByReferral:       input.ByReferral,
	}
	id, err := e.services.Users.SignUp(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})
}

type SignInInput struct {
	TelegramUsername string `json:"telegram_username"`
	Password         string `json:"password"`
}

func (e *Endpoint) SignIn(c *gin.Context) {
	var input SignInInput
	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, refreshToken, err := e.services.Users.SignIn(input.TelegramUsername, input.Password)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

type RefreshInput struct {
	RefreshToken string `json:"refresh_token"`
}

func (e *Endpoint) Refresh(c *gin.Context) {
	var input RefreshInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accessToken, refreshToken, err := e.services.Users.Refresh(input.RefreshToken)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{"access_token": accessToken, "refresh_token": refreshToken})
}

func (e *Endpoint) GetOneUser(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user, err := e.services.Users.GetOneUser(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (e *Endpoint) GetUserReferrals(c *gin.Context) {
	userId, err := e.GetUserId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	referralUsers, err := e.services.Users.GetUserReferrals(userId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": referralUsers,
	})
}
