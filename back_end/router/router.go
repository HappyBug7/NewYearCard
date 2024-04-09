package router

import (
	"Back_end/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	r.Use(middleware.Error)

	r.POST("/api", ctr.Card.AddCard)
	r.GET("/api", ctr.Card.GetCard)
	r.POST("/encrypt", ctr.Card.GetEncryptedID)
	r.POST("/decrypt", ctr.Card.GetDecryptedID)
	r.GET("/api/hello", ctr.Card.HelloWorld)
}
