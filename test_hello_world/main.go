package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Satu", "Dua", "Tiga")
	PublicHello()
	fmt.Print(getHello2())
}

func PublicHello() {
	fmt.Println("Public Hello World")
}

func getHello2() string {
	return "getHello2"
}
