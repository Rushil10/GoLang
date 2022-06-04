package main 

import (
	"fmt"
	"time"
)

func main() {
	c:=make(chan int)
	go func() {
		c <- 1
	}()
	val:= <- c
	fmt.Println(val)
	fmt.Println(c)
	go func() {
		c <- 2
	}()
	time.Sleep(time.Second * 2)
	val= <- c
	fmt.Println(val)
	fmt.Println(c)
}