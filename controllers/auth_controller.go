package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type Auth struct {
}

func (a *Auth) Login(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{"data": "{}"})
}