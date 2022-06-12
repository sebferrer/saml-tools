package httpserver

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Serve() {
	router := gin.Default()

	router.POST("/api/decrypt", Decrypt)
	router.StaticFS("/ui", http.Dir("../../ui"))

	router.Run(":8080")
}
