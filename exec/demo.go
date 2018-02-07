package main

import (
	"os/exec"
	"fmt"
)

func main(){
	//LookPath在环境变量中查找可执行文件
	f,err:=exec.LookPath("openssl")
	if err !=nil{
		fmt.Println("lookpath error:",err.Error())
	}

	fmt.Println("lookpath:",f)

	//command返回cmd结构来执行带有相关参数的命令
	//run开始指定命令并且等待他执行结束=>start()+wait()
	cmd:=exec.Command("mkdir","-p","a/b/c")
	err = cmd.Run()
	if err !=nil{
		fmt.Println("run error:",err.Error())
	}


}
