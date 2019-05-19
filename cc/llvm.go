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

func (nds nodes) PrintLlvm() {
	var v value.Value
	vars := make(map[string]*ir.InstAlloca)

	m := ir.NewModule()
	block := m.NewFunc("main", types.I64).NewBlock("")

	for _, nd := range nds {
		v = nd.gen(block, vars)
		if v.Type() != types.I64 {
			v = block.NewZExt(v, types.I64)
		}

		if nd.ty == ND_RETURN {
			break
		}
	}

	block.NewRet(v)
	fmt.Println(m)
}

func (nd *node) gen(block *ir.Block, vars map[string]*ir.InstAlloca) value.Value {
	switch nd.ty {
	case ND_NUM:
		return nd.genNum()
	case ND_VAR:
		return nd.genVar(block, vars)
	case ND_RETURN:
		return nd.genReturn(block, vars)
	case int('='):
		return nd.genAssign(block, vars)
	default:
		return nd.genOp(block, vars)
	}
}

func (nd *node) genNum() *constant.Int {
	return constant.NewInt(types.I64, int64(nd.val))
}

func (nd *node) genVar(block *ir.Block, vars map[string]*ir.InstAlloca) value.Value {
	r := nd.genLval(block, vars)
	return block.NewLoad(r)
}

func (nd *node) genReturn(block *ir.Block, vars map[string]*ir.InstAlloca) value.Value {
	v := nd.lhs.gen(block, vars)
	if v.Type() != types.I64 {
		v = block.NewZExt(v, types.I64)
	}
	return v
}

func (nd *node) genAssign(block *ir.Block, vars map[string]*ir.InstAlloca) value.Value {
	r := nd.lhs.genLval(block, vars)

	v := nd.rhs.gen(block, vars)
	if v.Type() != types.I64 {
		v = block.NewZExt(v, types.I64)
	}

	block.NewStore(v, r)
	return v
}

func (nd *node) genOp(block *ir.Block, vars map[string]*ir.InstAlloca) value.Value {
	v1 := nd.lhs.gen(block, vars)
	if v1.Type() != types.I64 {
		v1 = block.NewZExt(v1, types.I64)
	}

	v2 := nd.rhs.gen(block, vars)
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

func (nd *node) genLval(block *ir.Block, vars map[string]*ir.InstAlloca) *ir.InstAlloca {
	if nd.ty != ND_VAR {
		log.Fatal("left hand value is not a variable")
	}

	_, ok := vars[nd.name]
	if !ok {
		vars[nd.name] = block.NewAlloca(types.I64)
	}

	return vars[nd.name]
}
