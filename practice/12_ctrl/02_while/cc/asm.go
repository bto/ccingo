package cc

import (
	"fmt"
	"log"
)

func (nds nodes) PrintAsm() {
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	fmt.Println("  push rbp")
	fmt.Println("  mov rbp, rsp")
	fmt.Println("  sub rsp, 208")

	vars := newVariables()
	lb := newLabel()
	for _, nd := range nds {
		nd.gen(vars, lb)
		fmt.Println("  pop rax")
	}

	fmt.Println("  mov rsp, rbp")
	fmt.Println("  pop rbp")
	fmt.Println("  ret")
}

func (nd *node) gen(vars *variables, lb *label) {
	switch nd.ty {
	case ND_NUM:
		nd.genNum()
	case ND_VAR:
		nd.genVar(vars)
	case ND_RETURN:
		nd.genReturn(vars, lb)
	case ND_IF:
		nd.genIf(vars, lb)
	case ND_WHILE:
		nd.genWhile(vars, lb)
	case int('='):
		nd.genAssign(vars, lb)
	default:
		nd.genOp(vars, lb)
	}
}

func (nd *node) genNum() {
	fmt.Println("  push", nd.val)
}

func (nd *node) genVar(vars *variables) {
	nd.genLval(vars)
	fmt.Println("  pop rax")
	fmt.Println("  mov rax, [rax]")
	fmt.Println("  push rax")
}

func (nd *node) genReturn(vars *variables, lb *label) {
	nd.lhs.gen(vars, lb)
	fmt.Println("  pop rax")
	fmt.Println("  mov rsp, rbp")
	fmt.Println("  pop rbp")
	fmt.Println("  ret")
}

func (nd *node) genIf(vars *variables, lb *label) {
	lbIf := lb.get("if")
	nd.lhs.gen(vars, lb)
	fmt.Println("  pop rax")
	fmt.Println("  cmp rax, 0")
	fmt.Println("  je", lbIf)
	nd.rhs.gen(vars, lb)
	fmt.Println(lbIf + ":")
}

func (nd *node) genWhile(vars *variables, lb *label) {
	lbBegin := lb.get("begin")
	lbEnd := lb.get("end")
	fmt.Println(lbBegin + ":")
	nd.lhs.gen(vars, lb)
	fmt.Println("  pop rax")
	fmt.Println("  cmp rax, 0")
	fmt.Println("  je", lbEnd)
	nd.rhs.gen(vars, lb)
	fmt.Println("  jmp", lbBegin)
	fmt.Println(lbEnd + ":")
}

func (nd *node) genAssign(vars *variables, lb *label) {
	nd.lhs.genLval(vars)
	nd.rhs.gen(vars, lb)
	fmt.Println("  pop rdi")
	fmt.Println("  pop rax")
	fmt.Println("  mov [rax], rdi")
	fmt.Println("  push rdi")
}

func (nd *node) genOp(vars *variables, lb *label) {
	nd.lhs.gen(vars, lb)
	nd.rhs.gen(vars, lb)

	fmt.Println("  pop rdi")
	fmt.Println("  pop rax")

	switch nd.ty {
	case '+':
		fmt.Println("  add rax, rdi")
	case '-':
		fmt.Println("  sub rax, rdi")
	case '*':
		fmt.Println("  mul rdi")
	case '/':
		fmt.Println("  mov rdx, 0")
		fmt.Println("  div rdi")
	case ND_EQ:
		fmt.Println("  cmp rax, rdi")
		fmt.Println("  sete al")
		fmt.Println("  movzb rax, al")
	case ND_NE:
		fmt.Println("  cmp rax, rdi")
		fmt.Println("  setne al")
		fmt.Println("  movzb rax, al")
	case '<':
		fmt.Println("  cmp rax, rdi")
		fmt.Println("  setl al")
		fmt.Println("  movzb rax, al")
	case ND_LE:
		fmt.Println("  cmp rax, rdi")
		fmt.Println("  setle al")
		fmt.Println("  movzb rax, al")
	}

	fmt.Println("  push rax")
}

func (nd *node) genLval(vars *variables) {
	if nd.ty != ND_VAR {
		log.Fatal("代入の左辺値が変数ではありません")
	}

	if !vars.exist(nd.name) {
		vars.add(nd.name)
	}
	v := vars.get(nd.name)
	offset := v.offset

	fmt.Println("  mov rax, rbp")
	fmt.Println("  sub rax,", offset)
	fmt.Println("  push rax")
}
