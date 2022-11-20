package models

import (
	"github.com/gin-gonic/gin"
	"time"
)

type EndPoints map[string]func(c *gin.Context)

type RemotePlanInput struct {
	Name      string `json:"name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Workspace string `json:"workspace"`
	RunId     string `json:"runid"`
	Token     string `json:"token" binding:"required"`
}

type User struct {
	EndPoint    string    `json:"endpoint"`
	RequestType string    `json:"type"`
	Name        string    `json:"name"`
	Email       string    `json:"email"`
	Time        time.Time `json:"time"`
}
