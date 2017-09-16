package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main(){
	gin.SetMode(gin.DebugMode)
	router := gin.Default()

	router.Use(Middleware)
	router.GET("/simple/server/get", GetHandler)

	http.ListenAndServe(":8888", router)
}

func Middleware(c *gin.Context){
	fmt.Print("this is a middleWare!")
}

func GetHandler(c *gin.Context){
	value,exists := c.GetQuery("param")
	if (!exists){
		value = "the param is not exsits"
	}
	c.Data(http.StatusOK, "text/plain", []byte(fmt.Sprintf("get success %s\n", value)))
	return
}

func PostHandle(c *gin.Context){
	type JsonFormat struct {
		Id int `json:"id"`
		Name string `json:"name"`
	}
	jsonDecmo := JsonFormat{1,"sunzhibo"}
	c.JSON(http.StatusOK, jsonDecmo)
}