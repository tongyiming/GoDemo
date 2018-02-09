package main

import (
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type User struct{
	Username string `form:"username"`
	Age int `form:"age" binding:"test"`
}

func test(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	//使用val,ok:=field.Interface().(int)取字段值
	if val,ok:=field.Interface().(int);ok{
		fmt.Println("------------",val)
		if val ==0{return false}
	}
	return true
}


func getUser(c *gin.Context) {
	var u User
	if err := c.ShouldBindWith(&u, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "User dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}


func main(){

	//测试:curl "http://localhost:8080/user?username=tyming&age=10"
	//测试:curl "http://localhost:8080/user?username=tyming&age=0"
	route := gin.Default()
	//注册验证函数
	binding.Validator.RegisterValidation("test", test)
	route.GET("/user", getUser)
	route.Run(":8080")

}