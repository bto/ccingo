package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func main() {
	var v value.Value

	m := ir.NewModule()
	main := m.NewFunc("main", types.I64)
	block := main.NewBlock("")

	rd := bufio.NewReader(os.Stdin)

	num, c, err := readNum(rd)
	if err != nil && err != io.EOF {
		log.Fatal(err)
	}
	v = constant.NewInt(types.I64, int64(num))
	if err == io.EOF {
		block.NewRet(v)
		fmt.Println(m)
		return
	}

	for {
		if c == byte('+') {
			num, c, err = readNum(rd)
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
			v = block.NewAdd(v, constant.NewInt(types.I64, int64(num)))
		} else if c == byte('-') {
			num, c, err = readNum(rd)
			if err != nil && err != io.EOF {
				log.Fatal(err)
			}
			v = block.NewSub(v, constant.NewInt(types.I64, int64(num)))
		} else {
			log.Fatal("予期しない文字です:", c)
		}

		if err == io.EOF {
			break
		}
	}

	block.NewRet(v)
	fmt.Println(m)
}

func readNum(rd *bufio.Reader) (num int, c byte, err error) {
	var input []byte
	for {
		c, err = rd.ReadByte()
		if err != nil {
			break
		}

		if byte('0') <= c && c <= byte('9') {
			input = append(input, c)
		} else {
			break
		}
	}

	if len(input) == 0 {
		log.Fatal("予期しない文字です:", c)
	}

	num, err1 := strconv.Atoi(string(input))
	if err1 != nil {
		log.Fatal(err1)
	}
	return
}
