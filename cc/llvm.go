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

type context struct {
	mod  *ir.Module
	fn   *ir.Func
	bl   *ir.Block
	fns  map[string]*ir.Func
	vars map[string]*ir.InstAlloca
}

func (nds nodes) PrintLlvm() {
	var v value.Value
	cn := &context{
		mod: ir.NewModule(),
		fns: make(map[string]*ir.Func),
	}

	for _, nd := range nds {
		v = nd.gen(cn)
		if v == nil {
			continue
		}
		if v.Type() != types.I64 {
			v = cn.bl.NewZExt(v, types.I64)
		}
		if nd.ty == ND_RETURN {
			break
		}
	}

	fmt.Println(cn.mod)
}

func (nd *node) gen(cn *context) value.Value {
	switch nd.ty {
	case ND_NUM:
		return nd.genNum()
	case ND_VAR:
		return nd.genVar(cn)
	case ND_RETURN:
		return nd.genReturn(cn)
	case ND_IF:
		return nd.genIf(cn)
	case ND_WHILE:
		return nd.genWhile(cn)
	case ND_BLOCK:
		return nd.genBlock(cn)
	case ND_FUNC_CALL:
		return nd.genFuncCall(cn)
	case ND_FUNC_DEF:
		return nd.genFuncDef(cn)
	case int('='):
		return nd.genAssign(cn)
	case 0:
		return nd.genNoop(cn)
	default:
		return nd.genOp(cn)
	}
}

func (nd *node) genNum() *constant.Int {
	return constant.NewInt(types.I64, int64(nd.val))
}

func (nd *node) genVar(cn *context) value.Value {
	r := nd.genLval(cn)
	return cn.bl.NewLoad(r)
}

func (nd *node) genReturn(cn *context) value.Value {
	v := nd.lhs.gen(cn)
	if v.Type() != types.I64 {
		v = cn.bl.NewZExt(v, types.I64)
	}
	cn.bl.NewRet(v)
	return v
}

func (nd *node) genIf(cn *context) value.Value {
	blThen := cn.fn.NewBlock("then")
	blElse := cn.fn.NewBlock("else")
	blEnd := cn.fn.NewBlock("endif")

	cond := nd.lhs.gen(cn)
	if cond.Type() != types.I1 {
		cond = cn.bl.NewICmp(enum.IPredNE, cond, constant.NewInt(types.I64, 0))
	}
	cn.bl.NewCondBr(cond, blThen, blElse)

	cnThen := &context{
		fn:   cn.fn,
		bl:   blThen,
		vars: cn.vars,
	}
	nd.rhs.gen(cnThen)
	if blThen.Term == nil {
		blThen.NewBr(blEnd)
	}

	blElse.NewBr(blEnd)

	cn.bl = blEnd
	return nil
}

func (nd *node) genWhile(cn *context) value.Value {
	blCond := cn.fn.NewBlock("cond")
	blBegin := cn.fn.NewBlock("begin")
	blEnd := cn.fn.NewBlock("end")

	cn.bl.NewBr(blCond)

	cnCond := &context{
		fn:   cn.fn,
		bl:   blCond,
		vars: cn.vars,
	}
	cond := nd.lhs.gen(cnCond)
	if cond.Type() != types.I1 {
		cond = blCond.NewICmp(enum.IPredNE, cond, constant.NewInt(types.I64, 0))
	}
	blCond.NewCondBr(cond, blBegin, blEnd)

	cnBegin := &context{
		fn:   cn.fn,
		bl:   blBegin,
		vars: cn.vars,
	}
	nd.rhs.gen(cnBegin)
	blBegin.NewBr(blCond)

	cn.bl = blEnd
	return nil
}

func (nd *node) genBlock(cn *context) value.Value {
	var v value.Value
	for _, nd1 := range nd.nds {
		v = nd1.gen(cn)
		if nd1.ty == ND_RETURN {
			break
		}
	}
	return v
}

func (nd *node) genFuncCall(cn *context) value.Value {
	_, ok := cn.fns[nd.name]
	if !ok {
		params := []*ir.Param{}
		for i := 0; i < len(nd.nds); i++ {
			name := fmt.Sprintf("a%d", i)
			params = append(params, ir.NewParam(name, types.I64))
		}

		cn.fns[nd.name] = cn.mod.NewFunc(nd.name, types.I64, params...)
	}

	args := []value.Value{}
	for i := 0; i < len(nd.nds); i++ {
		args = append(args, nd.nds[i].gen(cn))
	}
	return cn.bl.NewCall(cn.fns[nd.name], args...)
}

func (nd *node) genFuncDef(cn *context) value.Value {
	params := []*ir.Param{}
	for i := 0; i < len(nd.nds); i++ {
		name := nd.nds[i].name
		params = append(params, ir.NewParam(name, types.I64))
	}
	cn.fn = cn.mod.NewFunc(nd.name, types.I64, params...)
	cn.fns[nd.name] = cn.fn
	cn.bl = cn.fn.NewBlock(nd.name)

	cn.vars = make(map[string]*ir.InstAlloca)
	for i := 0; i < len(params); i++ {
		param := params[i]
		v := cn.bl.NewAlloca(param.Typ)
		cn.vars[param.LocalIdent.Name()] = v
		cn.bl.NewStore(param, v)
	}

	v := nd.lhs.gen(cn)
	if v.Type() != types.I64 {
		v = cn.bl.NewZExt(v, types.I64)
	}
	cn.bl.NewRet(v)

	return nil
}

func (nd *node) genAssign(cn *context) value.Value {
	r := nd.lhs.genLval(cn)

	v := nd.rhs.gen(cn)
	if v.Type() != types.I64 {
		v = cn.bl.NewZExt(v, types.I64)
	}

	cn.bl.NewStore(v, r)
	return v
}

func (nd *node) genNoop(cn *context) value.Value {
	nd.lhs.gen(cn)
	nd.rhs.gen(cn)
	return nil
}

func (nd *node) genOp(cn *context) value.Value {
	v1 := nd.lhs.gen(cn)
	if v1.Type() != types.I64 {
		v1 = cn.bl.NewZExt(v1, types.I64)
	}

	v2 := nd.rhs.gen(cn)
	if v2.Type() != types.I64 {
		v2 = cn.bl.NewZExt(v2, types.I64)
	}

	switch nd.ty {
	case '+':
		return cn.bl.NewAdd(v1, v2)
	case '-':
		return cn.bl.NewSub(v1, v2)
	case '*':
		return cn.bl.NewMul(v1, v2)
	case '/':
		return cn.bl.NewUDiv(v1, v2)
	case ND_EQ:
		return cn.bl.NewICmp(enum.IPredEQ, v1, v2)
	case ND_NE:
		return cn.bl.NewICmp(enum.IPredNE, v1, v2)
	case '<':
		return cn.bl.NewICmp(enum.IPredULT, v1, v2)
	case ND_LE:
		return cn.bl.NewICmp(enum.IPredULE, v1, v2)
	}

	log.Fatal("invalid node: ", nd)
	return nil
}

func (nd *node) genLval(cn *context) *ir.InstAlloca {
	if nd.ty != ND_VAR {
		log.Fatal("left hand value is not a variable")
	}

	_, ok := cn.vars[nd.name]
	if !ok {
		cn.vars[nd.name] = cn.bl.NewAlloca(types.I64)
	}

	return cn.vars[nd.name]
}
