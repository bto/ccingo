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

	m := ir.NewModule()
	block := m.NewFunc("main", types.I64).NewBlock("")

	vars := make(map[string]*ir.InstAlloca)
	vars["a"] = block.NewAlloca(types.I64)
	vars["b"] = block.NewAlloca(types.I64)
	vars["c"] = block.NewAlloca(types.I64)
	vars["d"] = block.NewAlloca(types.I64)
	vars["e"] = block.NewAlloca(types.I64)
	vars["f"] = block.NewAlloca(types.I64)
	vars["g"] = block.NewAlloca(types.I64)
	vars["h"] = block.NewAlloca(types.I64)
	vars["i"] = block.NewAlloca(types.I64)
	vars["j"] = block.NewAlloca(types.I64)
	vars["k"] = block.NewAlloca(types.I64)
	vars["l"] = block.NewAlloca(types.I64)
	vars["m"] = block.NewAlloca(types.I64)
	vars["n"] = block.NewAlloca(types.I64)
	vars["o"] = block.NewAlloca(types.I64)
	vars["p"] = block.NewAlloca(types.I64)
	vars["q"] = block.NewAlloca(types.I64)
	vars["r"] = block.NewAlloca(types.I64)
	vars["s"] = block.NewAlloca(types.I64)
	vars["t"] = block.NewAlloca(types.I64)
	vars["u"] = block.NewAlloca(types.I64)
	vars["v"] = block.NewAlloca(types.I64)
	vars["w"] = block.NewAlloca(types.I64)
	vars["x"] = block.NewAlloca(types.I64)
	vars["y"] = block.NewAlloca(types.I64)
	vars["z"] = block.NewAlloca(types.I64)

	for _, nd := range nds {
		v = nd.gen(block, vars)
		if v.Type() != types.I64 {
			v = block.NewZExt(v, types.I64)
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
	r := nd.genLval(vars)
	return block.NewLoad(r)
}

func (nd *node) genReturn(block *ir.Block, vars map[string]*ir.InstAlloca) value.Value {
	return nd.lhs.gen(block, vars)
}

func (nd *node) genAssign(block *ir.Block, vars map[string]*ir.InstAlloca) value.Value {
	r := nd.lhs.genLval(vars)

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

func (nd *node) genLval(vars map[string]*ir.InstAlloca) *ir.InstAlloca {
	if nd.ty != ND_VAR {
		log.Fatal("left hand value is not a variable")
	}

	return vars[nd.name]
}
