//测试wait.Until() 的用途
package main

import (
	"fmt"
	"k8s.io/apimachinery/pkg/util/wait"
	"time"
)

type stop struct {
}

func main() {
	stopCh := make(chan struct{})
	//初始化一个计数器
	i := 0
	go wait.Until(func() {
		fmt.Printf("----%d----\n", i)

		i++
	}, time.Second*10, stopCh)

}
