package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type arr []string

type bot interface {
	getGreetings() string
}

type englishBot struct{
	teacherName string
	teacherAge int64
}
type spanishBot struct{}

func main() {
	english := englishBot{"Mrs. Smith", 26}
	spanish := spanishBot{}

	fmt.Printf("%+v\n",english)
	english.changeName("Mrs. Henderson")
	fmt.Printf("%+v\n",english)

	printGreetings(spanish)
	printGreetings(english)

	english.saveFile("english.txt")
	s := spanish.readFile("spanish.txt")

	fmt.Println(s)

	strArr := arr{}
	strArr.printArr(strArr.intoArray(s))
	fmt.Println("Length of the array is",len(strArr.intoArray(s)))
}

func (e *englishBot) changeName(newName string) {
	e.teacherName = newName
}

func (arr) intoArray(s string) []string {
	return strings.Split(s, " ")
}

func(arr) printArr(array []string) {
	fmt.Println("This is an array:",array)
}

func (eb englishBot) saveFile(fileName string) error {
	return ioutil.WriteFile(fileName, []byte(eb.getGreetings()), 0666)
} 

func (spanishBot) readFile(fileName string) string {
	bs, err := ioutil.ReadFile(fileName)
	if err != nil {
		os.Exit(1)
	}

	return string(bs)
}

func printGreetings(b bot) {
	fmt.Println(b.getGreetings())
}

func (englishBot) getGreetings() string {
	return "Hello world"
}

func (spanishBot) getGreetings() string {
	return "Hola"
}