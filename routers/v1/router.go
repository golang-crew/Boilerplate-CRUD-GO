package router

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-crew/Bolierplate-CRUD-Gingonic/handler"
)

func ApplyRoutes(r *gin.RouterGroup) {
	test := r.Group("/test")
	{
		test.GET("/", func(c *gin.Context) {
			c.String(200, "pong")
		})
	}
	memos := r.Group("/memos")
	{
		memos.GET("/", handler.GetMemoList)
		memos.POST("/", handler.CreateMemo)
		memos.GET("/:memoID", handler.GetMemo)
		memos.DELETE("/:memoID", handler.DeleteMemo)
		memos.PUT("/:memoID", handler.UpdateMemo)
	}
}
