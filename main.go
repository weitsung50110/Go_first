package main

import (
	"fmt"

	"github.com/gin-gonic/gin" // 導入 Gin 框架
)

func main() {
	// 建立一個 Gin 的預設路由器
	router := gin.Default()

	// 設定靜態檔案伺服器，指定靜態檔案的存放路徑為 "./static"
	router.Static("./static", "./static")

	// 設定一個 GET 請求路由，當訪問根目錄時，回傳 index.html 檔案
	router.GET("/", func(c *gin.Context) {
		c.File("templates/index.html")
	})

	// 設定一個 POST 請求路由，處理來自 /submit 路徑的 POST 請求
	router.POST("/submit", func(c *gin.Context) {
		// 從 POST 請求中獲取名為 inputText 和 inputName 的表單數據
		inputText := c.PostForm("inputText")
		inputName := c.PostForm("inputName")
		fmt.Println("inputName : ", inputName) // 將 inputName 印出到終端機上

		// 將資料封裝為 JSON 格式，使用 gin 框架的 gin.H(map) 型別
		responseData := gin.H{
			"inputText": inputText,
			"inputName": inputName,
		}

		// 將 JSON 格式的資料作為回應返回給前端
		c.JSON(200, responseData)
	})

	// 啟動伺服器，監聽端口 8080
	router.Run(":8080")
}
