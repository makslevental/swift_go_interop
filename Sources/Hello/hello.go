package main

import "fmt"
import "C"

//export hello
func hello() {
   fmt.Printf("Hello from Go!")
}

func main() {}