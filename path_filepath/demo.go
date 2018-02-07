package main

import (
	"fmt"
	"path/filepath"
)

//filepath包，兼容各操作系统的文件路径
func main(){
	//dir返回路径最后一个元素的目录
	fmt.Println(filepath.Dir("a/b/c"))

	//base返回路径最后一个元素
	fmt.Println(filepath.Base("a/b/c"))

	//join连接路径
	fmt.Println(filepath.Join("a/b/c","a"))

	//abs返回绝对路径
	path, _ := filepath.Abs("./1.txt");
	fmt.Println(path);

	//返回以basepath为基准的相对路径
	path, _ = filepath.Rel("./a/b", "./a/b/c/d/e");
	fmt.Println(path);
}
