package main

import (
	"fmt"
	"log"
)

func main() {
	var v int
	_, err := fmt.Scan(&v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")
	fmt.Printf("  mov rax, %d\n", v)
	fmt.Println("  ret")
}
