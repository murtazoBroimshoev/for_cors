package handlers

import (
	"Murtazo/app/controlers"

	"github.com/gin-gonic/gin"
)

func Handlers() {
	r := gin.Default()
	r.POST("sigup",controlers.Register)
}