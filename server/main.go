package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	// "strconv"
	"fmt"
)





func main() {

	router := gin.Default()
	router.LoadHTMLGlob("./templates/*.html")
	router.Static("/img", "./img/")


	// router.GET("/research/:route/:page/", func(c *gin.Context) {
	// 	var page_temp int
	// 	route := c.Param("route")
	// 	page_temp, _ = strconv.Atoi(c.Param("page"))
	// 	page := fmt.Sprintf("%04d.jpg", page_temp)
	// 	fmt.Println(route, page)
		
	// 	c.HTML(http.StatusOK, "research.html", gin.H{
	// 		"route": "living_room",
	// 		"page": "0000.jpg",
	// 	})
	// })

	var page int = 0

	router.GET("/research/", func(c *gin.Context) {

		page = 0
		
		c.HTML(http.StatusOK, "research.html", gin.H{
			"route": "living_room",
			"page": "0000.jpg",
		})
	})

	router.POST("/post/", func(c *gin.Context) {
		s := c.PostForm("str")
		fmt.Println(s)
		page+=1
		c.HTML(http.StatusOK, "research.html", gin.H{
			"route": "living_room",
			"page": fmt.Sprintf("%04d.jpg", page),
		})
	})

	router.Run(":3000") // 0.0.0.0:8080 でサーバーを立てます。

}