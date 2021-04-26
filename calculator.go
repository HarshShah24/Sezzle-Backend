package main

import "github.com/gin-gonic/gin"

type REQUEST struct {
	A  int    `json:"a" binding:"required"`
	B  int    `json:"b" binding:"required"`
	OP string `json:"op" binding:"required"`
}

var result int

func doCalulcations(c *gin.Context) {
	var request REQUEST
	c.BindJSON(&request)

	switch request.OP {
	case "+":
		result = request.A + request.B
	case "-":
		result = request.A - request.B
	case "*":
		result = request.A * request.B
	case "/":
		result = request.A / request.B
	}

	println(result)

	c.JSON(201, gin.H{"result": result})
}

func getLastResult(c *gin.Context) {
	c.JSON(200, gin.H{"result": result})
}
