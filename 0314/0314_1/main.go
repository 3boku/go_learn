package main

import "fmt"

func repeatMe(words ...string) {
	fmt.Println(words)
}


func main() {
	repeatMe("3boku", "woong", "msg", "good")
}