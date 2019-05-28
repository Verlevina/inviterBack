package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func httpHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Access-Control-Allow-Origin", "*");
}

func Startup() {
	r := gin.Default()
	http.HandleFunc("/", httpHandler);

	r.GET("/ping", pong)
	// r.GET("/user", getUser)
	// r.POST("/user", postUser)
	r.GET("API/getEventTemplate", getEventTemplate)
	r.Run(":8080")
}


