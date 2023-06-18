package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("myfile.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	io.Copy(os.Stdout, f)
}