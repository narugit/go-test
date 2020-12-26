package main

import "fmt"

import (
	mm "go-test/mymod/io"
	"github.com/nsf/termbox-go"
)


func main() {
	fmt.Printf("Hello World\n")
	mm.DoSome()
  if err := termbox.Init(); err != nil {
        fmt.Printf("error")
  }
  defer termbox.Close()
}
