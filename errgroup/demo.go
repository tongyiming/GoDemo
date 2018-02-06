package main

import (
	"github.com/golang/sync/errgroup"
	"fmt"
	"github.com/pkg/errors"
)

var g errgroup.Group

func main(){


	g.Go(func() error{
		return errors.New("1234567890")
	})

	for i:=0;i<5;i++{
		i:=i
		g.Go(func() error{
			if i%2==0{
				return errors.New("error!")
			}else{
				return nil
			}

		})
	}


	//返回第一个非空的error
	if err:= g.Wait();err == nil{
		fmt.Println("finished")
	}else{
		fmt.Println("error is:",err.Error())
	}

}
