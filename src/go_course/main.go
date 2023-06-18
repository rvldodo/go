package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName, lastName, gender string
	age int
}

type area interface {
	AreaValue() float64
}

type value struct {
	height float64
	width float64
}

func (v value) AreaValue() float64 {
	return v.height * v.width
}


func (p Person) greet() string {
	return "Hello my name is " + p.firstName + " " + p.lastName + " and I'm " + strconv.Itoa(p.age) + " years old"
}

func main() {
	// person := Person{"Rivaldo", "Lawalata", "m", 21}

	// var v value
	// v = value{5, 2}
	// fmt.Println("Area value: ", v.AreaValue())

	// fmt.Println(person.greet())
	arr := map[int] string{1: "ritza.kerz18@gmail.com"}
	for i, v := range arr {
		fmt.Printf("%d - ID: %s\n", i, v)
	}
}
