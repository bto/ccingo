package cc

import (
	"fmt"
	"log"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
	"github.com/llir/llvm/ir/value"
)

func (nd *node) PrintLlvm() {
	m := ir.NewModule()
	block := m.NewFunc("main", types.I64).NewBlock("")

	v := nd.gen(block)
	if v.Type() != types.I64 {
		v = block.NewZExt(v, types.I64)
	}
	block.NewRet(v)

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
	if v1.Type() != types.I64 {
		v1 = block.NewZExt(v1, types.I64)
	}

	v2 := nd.rhs.gen(block)
	if v2.Type() != types.I64 {
		v2 = block.NewZExt(v2, types.I64)
	}

	switch nd.ty {
	case '+':
		return block.NewAdd(v1, v2)
	case '-':
		return block.NewSub(v1, v2)
	case '*':
		return block.NewMul(v1, v2)
	case '/':
		return block.NewUDiv(v1, v2)
	case ND_EQ:
		return block.NewICmp(enum.IPredEQ, v1, v2)
	case ND_NE:
		return block.NewICmp(enum.IPredNE, v1, v2)
	case '<':
		return block.NewICmp(enum.IPredULT, v1, v2)
	case ND_LE:
		return block.NewICmp(enum.IPredULE, v1, v2)
	}

	log.Fatal("invalid node: ", nd)
	return nil
}
