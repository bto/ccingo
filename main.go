package main

import (
	"fmt"
	"log"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func main() {
	var v int
	_, err := fmt.Scan(&v)
	if err != nil {
		log.Fatal(err)
	}

	m := ir.NewModule()

	main := m.NewFunc("main", types.I8)
	entry := main.NewBlock("")
	entry.NewRet(constant.NewInt(types.I8, int64(v)))

	fmt.Println(m)
}
