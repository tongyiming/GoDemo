package main

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func middleware1(c *gin.Context){
	fmt.Println("middleware1 begin")
	c.Next()
	fmt.Println("middleware1 end")
}

func middleware2(c *gin.Context){
	fmt.Println("middleware2 begin")
	c.Next()
	fmt.Println("middleware2 end")
}

func handler(c *gin.Context){
	fmt.Println("test......")
}
func main(){
	/*
		类似与JAVA中的拦截器
		测试结果:
		middleware1 begin
		middleware2 begin
		test......
		middleware2 end
		middleware1 end
	 */
	router:=gin.Default()

	router.GET("/",middleware1,middleware2,handler)
	//默认8080
	router.Run()
}