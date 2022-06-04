package main

import "fmt"

func main() {
	f:=true
	flag := &f
	if flag==nil {
		fmt.Println("Nil value in flag")
	} else if *flag {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}

	for i := 0;i<10;i++ {
		fmt.Println(i)
	}

	arr:=[]string{"Rushil","Shah"}
	for i,value := range arr {
		fmt.Println(i,value)
	}

	mymap:= make(map[string]interface{})
	mymap["name"]="Rushil"
	mymap["age"]=21
	for k,v := range mymap {
		fmt.Printf("Key %s Value %v",k,v)
		fmt.Println()
	}
}