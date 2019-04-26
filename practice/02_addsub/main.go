package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	ErrInvalidCharacter = errors.New("予期しない文字です")
)

func readNum(rd *bufio.Reader) (num []byte, c byte, err error) {
	for {
		c, err = rd.ReadByte()
		if err != nil {
			break
		}

		if byte('0') <= c && c <= byte('9') {
			num = append(num, c)
		} else {
			break
		}
	}
	if len(num) == 0 {
		err = ErrInvalidCharacter
	}

	return
}

func main() {
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	rd := bufio.NewReader(os.Stdin)

	num, c, err := readNum(rd)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	fmt.Println("  mov rax,", string(num))
	if err == io.EOF {
		fmt.Println("  ret")
		return
	}

	for {
		var op string
		if c == byte('+') {
			op = "add"
		} else if c == byte('-') {
			op = "sub"
		} else {
			log.Fatal("予期しない文字です:", c)
		}

		num, c, err = readNum(rd)
		if err != nil && err != io.EOF {
			log.Fatal(err)
		}
		fmt.Println(" ", op, "rax,", string(num))
		if err == io.EOF {
			break
		}
	}

	fmt.Println("  ret")
}
