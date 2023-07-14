package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan bool)
	people := [2] string{"sambok", "nico"}
	for _, person := range people {
		go isSexy(person, c)
	}
	fmt.Println(<- c)
	fmt.Println(<- c)
}

// func sexyCount(person string){
// 	for i:=0;i<10;i++ {
// 		fmt.Println(person, "is sexy", i)
// 		time.Sleep(time.Second)
// 	}
// }

func isSexy(person string, c chan bool){
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c<-true
}