package main

import "fmt"

//inside this file, we do not have a main function
//remember: main function only appears once in the program

var points = []int{20, 90, 100, 45, 70}

func sayHello(n string) {
	fmt.Println("Hello", n)
}

func showScore() {
	fmt.Println("you scored this many points:", score) //score variable is shared from main.go
}
