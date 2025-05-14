package router

import (
	"net/http"
	"server/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 静态文件托管（正式部署用）
	r.Static("/assets", "../client/dist/assets")
	r.LoadHTMLFiles("../client/dist/index.html")

	// 主页
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	api := r.Group("/api")
	{
		api.GET("/questions", controller.GetQuestions)
		api.POST("/questions", controller.AddQuestion)
		api.PUT("/questions/:id", controller.UpdateQuestion)
		api.DELETE("/questions/:id", controller.DeleteOneQuestion) // 单个删除
		api.POST("/questions/delete", controller.DeleteQuestions)  // 批量删除

		api.POST("/questions/ai_generate", controller.GenerateByAI)
		api.GET("/readme", controller.GetReadme)
	}

	return r
}
