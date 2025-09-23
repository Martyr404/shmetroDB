package server

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	id string
}

// 定义JSON请求体结构（与客户端发送的JSON字段对应）
type RequestBody struct {
	Carriage_number string `json:"carriage_number"` // 注意JSON标签要与客户端一致
}

func (s Server) Init() {
	router := gin.Default()

	// 推荐使用POST方法处理带请求体的请求
	router.POST("/calculate", func(c *gin.Context) {
		// 读取请求体
		rawBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "read body error"})
			return
		}

		// 解析JSON格式
		var req RequestBody
		if err := json.Unmarshal(rawBody, &req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"msg":    "parse error (expect JSON format: {\"carriage_number\": \"xxx\"})",
				"detail": err.Error(), // 增加错误详情便于调试
			})
			return
		}

		// 验证参数
		if req.Carriage_number == "" {
			c.JSON(http.StatusBadRequest, gin.H{"msg": "missing carriage_number"})
			return
		}

		// 处理成功
		c.JSON(http.StatusOK, gin.H{
			"msg": "you are searching number " + req.Carriage_number,
		})
	})

	router.Run(":9987") // 启动服务
}
