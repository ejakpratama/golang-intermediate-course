package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("Satu", "Dua", "Tiga")
	PublicHello()
	fmt.Print(GetHello2())
}

func PublicHello() {
	fmt.Println("Public Hello World")
}

func GetHello2() string {
	return "hello world 1"
}
