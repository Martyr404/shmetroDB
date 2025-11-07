package server

import (
	"encoding/json"
	"io"
	"net/http"
	"shmetroDB/middleware"
	"shmetroDB/orm"

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
	router.Use(middleware.CORSMiddleware())

	api := router.Group("/api")

	api.POST("/calculate", func(c *gin.Context) {
		// 读取请求体
		rawBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"StateCode": "4000",
				"Msg":       "ioReader Read Request Body Error",
				"Data":      "",
			})
			return
		}

		// 解析JSON格式
		var req RequestBody
		if err := json.Unmarshal(rawBody, &req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"StateCode": "4001",
				"Msg":       "invalid carriage_number",
				"Data":      "",
			})
			return
		}

		// 验证参数是否为空
		if req.Carriage_number == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"StateCode": "4002",
				"Msg":       "invalid carriage_number",
				"Data":      "",
			})
			return
		}
		// 验证车厢号是否合法
		if !orm.ValidateCarriageNumbers(req.Carriage_number) {
			c.JSON(http.StatusBadRequest, gin.H{
				"StateCode": "4002",
				"Msg":       "invalid carriage_number",
				"Data":      "",
			})
			return
		}

		//计算车号
		trainInfos, Err := orm.ParseCarriageNumber(req.Carriage_number)
		if Err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"StateCode": "5000",
				"Msg":       "parse carriage number goes wrong",
				//方便debug用，实际上线用""替换Err
				"Data": Err,
			})
		} else {
			//trainInfo 以下不应该为空
			//遍历查询车号详情
			json_data := make([]orm.TrainInfo, 0)
			for index := range trainInfos {
				if !trainInfos[index].IsEmpty {
					Err := orm.QueryInfo(trainInfos[index].TrainId, &trainInfos[index])
					if Err != nil {
						c.JSON(http.StatusInternalServerError, gin.H{
							"StateCode": "5001",
							"Msg":       "Query carriage number goes wrong",
							//方便debug用，实际上线用""替换Err
							"Data": Err,
						})
						break
					}
					if !trainInfos[index].IsEmpty {
						json_data = append(json_data, trainInfos[index])
					}
				}
			}
			c.JSON(http.StatusAccepted, gin.H{
				"StateCode": "2000",
				"Msg":       "Query carriage number:" + req.Carriage_number + " successfully",
				"Data":      json_data,
			})
		}

	})
	router.Run(":9987") // 启动服务
}
