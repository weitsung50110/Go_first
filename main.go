package main

import (
	"github.com/gin-gonic/gin" //導入GIN
)

func main() {
	router := gin.Default()

	// 伺服器回傳 index.html
	router.GET("/", func(c *gin.Context) {
		c.File("templates/index.html")
	})

	// 從 POST 請求中獲取 inputText
	router.POST("/submit", func(c *gin.Context) {
		inputText := c.PostForm("inputText")
		// 將 inputText 回傳給前端
		c.String(200, inputText)
	})

	// 設定靜態檔案伺服器
	router.Static("/static", "./static")

	// 啟動伺服器，監聽端口 8080
	router.Run(":8080")
}
