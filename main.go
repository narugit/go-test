package main

import "fmt"

import (
	mm "github.com/narugit/go-test/struct/mymod"
)

type Hog mm.Hoge

func main() {
	fmt.Printf("Hello World\n")
	Hog.fuga = 1
	fmt.Printf("%d", Hog.fuga)
}
