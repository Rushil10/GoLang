package main 

import (
	"fmt"
	"sync"
)

func main() {
	wg:=&sync.WaitGroup{}
	wg.Add(2)
	go func() {
		fmt.Println("Hello")
		wg.Done()
	}()
	go func() {
		fmt.Println("World")
		wg.Done()
	}()
	fmt.Println("Printed First")
	wg.Wait()
	fmt.Println("Exit")
}