package main

import "github.com/gin-gonic/gin"
import "newsfeeder/httpd/handler"

func main() {
	r := gin.Default()

	r.GET("/ping", handler.PingGet)

	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
