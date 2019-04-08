package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	var num []byte
	var c byte
	var err error
	zero := byte('0')
	nine := byte('9')
	plus := byte('+')
	minus := byte('-')
	rd := bufio.NewReader(os.Stdin)

	for {
		c, err = rd.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		if zero <= c && c <= nine {
			num = append(num, c)
		} else {
			break
		}
	}
	if len(num) == 0 {
		log.Fatal("予期しない文字です:", c)
	}
	fmt.Println("  add rax,", string(num))

	for {
		if c == plus {
			op := "add"
		} else if c == minus {
			op := "sub"
		} else {
			log.Fatal("予期しない文字です:", c)
		}
	}

	/*
	if len(os.Args) != 2 {
		log.Fatal("引数の個数が正しくありません")
	}

	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	var v int
	s := os.Args[1]

	n, err := fmt.Sscanf(os.Args[1], "%d%s", &v, &s)
	if err != nil {
		fmt.Println("code", err.Error())
		log.Fatal(err)
	}
	// fmt.Printf("  mov rax, %d\n", v)
	fmt.Println("n is", n)
	fmt.Println("err is", err)
	fmt.Println("v is", v)
	fmt.Println("s is", s)

	fmt.Println("  ret")
	*/
}
