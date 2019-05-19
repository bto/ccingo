package cc

import (
	"fmt"
	"log"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (nd *node) PrintLlvm() {
	m := ir.NewModule()
	block := m.NewFunc("main", types.I64).NewBlock("")
	block.NewRet(nd.gen(block))
	fmt.Println(m)
}

func (nd *node) gen(block *ir.Block) value.Value {
	if nd.ty == ND_NUM {
		return nd.genNum()
	} else {
		return nd.genOp(block)
	}
}

func (nd *node) genNum() *constant.Int {
	return constant.NewInt(types.I64, int64(nd.val))
}

func (nd *node) genOp(block *ir.Block) value.Value {
	v1 := nd.lhs.gen(block)
	v2 := nd.rhs.gen(block)

	switch nd.ty {
	case '+':
		return block.NewAdd(v1, v2)
	case '-':
		return block.NewSub(v1, v2)
	}

	log.Fatal("invalid node: ", nd)
	return nil
}
