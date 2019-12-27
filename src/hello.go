package main

import "fmt"
import "./pack1"
import "./go_struct"


func learn(){
	fmt.Printf("learn go")
	var a string = "Runoob"
	fmt.Printf(a)
}

func main() {
    fmt.Printf("hello, world\n")
	learn()
	subClass.Pack1()
	go_struct.UseStruct()
}