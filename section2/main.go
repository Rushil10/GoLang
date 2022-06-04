package main

import (
	"fmt"
	"strings"
)

func main() {
	var m1 int
	m1=2
	var (
		m3=4
		m2=5
	)
	fmt.Println(m1+m2+m3)

	//typeCasting
	var m4 int64
	var m5 int32
	fmt.Println(int64(m5) + m4)

	m6:=2
	m7:=9
	fmt.Println(m6+m7)

	//Mutable Strings
	name:="My name"
	word:="Rushil"
	name="My name is Rushil"
	fmt.Println(strings.Contains(name,word))
	todo()
	ptrs()
	swap(&m1, &m2)
	fmt.Println(m1,m2)
	strs()
}

func todo() {
	//var arr[] int
	arr:=[] int {1,2,3,4}
	arr=append(arr,5)
	fmt.Println(arr)
}

func ptrs() {
	p:=4
	pt1 := &p
	fmt.Println(*pt1)
}

func swap(m1 , m2 *int){
	var temp int
	temp=*m1
	*m1=*m2
	*m2=temp
}

//Data Encapsulation

type Car struct {
	Name string
	Age int
	ModelNo int
}

func (c Car) Print() {
	fmt.Println(c)
}

func (c Car) Drive() {
	fmt.Println("Driving ...")
}

func (c Car) getName() string {
	return c.Name
}

func strs() {
	//var c1 Car
	c:=Car{
		Name:"Wagonr",
		Age:5,
		ModelNo:2801,
	}
	c.Print()
	c.Drive()
	fmt.Println(c.getName())
}