package main

import "fmt"

func main() {
    go fmt.Println("горутина")
    fmt.Println("main")
}