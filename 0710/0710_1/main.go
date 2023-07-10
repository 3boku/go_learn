package main

import (
	"fmt"
)

func main() {
	names := [5]string{"a", "b", "c", "d", "e"};
	fmt.Println(names)
	//arrays

	names2 := []string{"a", "b", "c", "d", "e"};
	names2 = append(names2, "f")
	//slace할땐 append가 값을 리턴해주는 형식이라 앞에 슬라이스를 꼭 써야한다
	fmt.Println(names2)
}