package main

import (
	"html/template"
	"fmt"
	"os"
)

type User struct{
	Name string
	Age int
	Number []int
}

const temp = `{{.Name}} age is {{inc .Age}} and "." is {{.}}
{{range $i,$v := .Number}}
{{inc $i}} = {{$v}}{{end}}`
func main(){

	t:=template.New("index.html")
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
	}
	src,err:=t.Funcs(funcMap).Parse(temp)
	if err!=nil{
		fmt.Println("parse err:",err.Error())
	}
	num:=[]int{1,2,3,4,5}
	user:=User{Name:"tyming",Age:22,Number:num}
	if err := src.Execute(os.Stdout,user);err!=nil{
		fmt.Println("execute error:",err.Error())
	}

}