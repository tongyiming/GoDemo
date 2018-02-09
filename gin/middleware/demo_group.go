package main

import (
"github.com/gin-gonic/gin"
	"fmt"
)
func middleware3(c *gin.Context){
	fmt.Println("middleware1 begin")
	c.Next()
	fmt.Println("middleware1 end")
}

func middleware4(c *gin.Context){
	fmt.Println("middleware2 begin")
	c.Next()
	fmt.Println("middleware2 end")
}

func handler1(c *gin.Context){
	fmt.Println("test......")
}

func main(){
	/*
		类似与JAVA中的拦截器
		测试结果:
		middleware1 begin
		test......
		middleware2 begin
		middleware2 end
		middleware1 end
	 */
	router:=gin.Default()
	//加载中间件router.Use()
	g:=router.Group("/g",middleware3)

	g.GET("/test",handler1,middleware4)
	//默认8080
	router.Run()
}