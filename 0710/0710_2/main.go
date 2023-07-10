package main

import "fmt"

// func main(){
// 	sambok := map[string]string{"name": "sambok", "age": "16"}
// 	fmt.Println(sambok)

// 	for _, value := range sambok{
// 		fmt.Println(value)
// 	}
// }

type person struct{
	name string
	age int
	favFood []string
}
func main(){
	favFood := []string{
		"kimchi",
		"ramen",
		"udon",
	}
	sambok := person{
		"sambok",
		16,
		favFood,
	}
	//위와 아래는 똑같은 구조체 이다.
	sambok2 := person{
		name:"sambok",
		age:16,
		favFood:favFood,
	}

	fmt.Println(
		sambok.name,
		sambok2,
	)
}