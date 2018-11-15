package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

/*
REST请求是很直观的，因为REST是基于HTTP协议的一个补充，他的每一次请求都是一个HTTP请求，
然后根据不同的method来处理不同的逻辑，很多Web开发者都熟悉HTTP协议
，所以学习REST是一件比较容易的事情。所以我们在8.3小节将详细的讲解如何在Go语言中来实现REST方式。
*/
