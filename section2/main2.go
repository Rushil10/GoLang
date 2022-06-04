package main

import "fmt"

type Car interface {
	Drive()
	Stop()
}

func newCar(arg string) Car {
	return &Lambo{arg}
}

type Lambo struct {
	LamboModel string
}

type Chevy struct {
	ChevyModel string
}

func (l *Lambo) Drive() {
	fmt.Println("Driving Lamborgini..")
	fmt.Println(l.LamboModel)
}

func (l *Lambo) Stop() {
	fmt.Println("Lambo Stopped !")
}

func (c *Chevy) Drive() {
	fmt.Println("Driving Chevy")
	fmt.Println(c.ChevyModel)
}

func Anything(anything interface{}){
	fmt.Println(anything)
}

func main() {
	l:=newCar("Collarado")
	c:=Chevy{"C007"}
	l.Drive()
	c.Drive()
	l.Stop()
	Anything("Hello")
	Anything(1.22)
	Anything(struct{}{})
	mymap:=make(map[string]interface{})
	mymap["Name"]="Rushil"
	mymap["Age"]=21
	fmt.Println(mymap)
}