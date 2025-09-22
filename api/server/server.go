package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func server(port string) {
	router := gin.Default()
	router.GET("/calculate", func(c *gin.Context) {
		body := c.Request.Body
		var raw []byte
		_, err := body.Read(raw)
		if err != nil {
			fmt.Printf("reading request body err:%s\n", err.Error())
		}

	})
	router.Run(":" + port)
}
