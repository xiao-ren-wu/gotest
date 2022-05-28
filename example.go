package gotest

import "fmt"

func SayHello() {
	fmt.Println("hello World")
}

func SayGoodBye() {
	fmt.Println("hello")
	fmt.Println("good bye")
}

func SayName() {
	nameMap := map[int]string{
		1: "Tom",
		2: "Jim",
		3: "Kitty",
		4: "Erik",
	}

	for id, name := range nameMap {
		fmt.Printf("%d::%s\n", id, name)
	}
}
