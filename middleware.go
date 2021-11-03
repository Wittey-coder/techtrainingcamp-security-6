package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)
type Response struct {
	Code int
	Message string
	Data interface{}
}

func RequestInfo() gin.HandlerFunc{
	return func(context *gin.Context) {
		path := context.FullPath()
		method := context.Request.Method
		fmt.Println("Request: " + method)
		fmt.Println("Path: " + path)
		context.Next()
		fmt.Println("Status: ", context.Writer.Status())
	}

}