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
		fmt.Println("  push", nd.val)
		return
	case ND_VAR:
		nd.genLval(vars)
		fmt.Println("  pop rax")
		fmt.Println("  mov rax, [rax]")
		fmt.Println("  push rax")
		return
	case ND_RETURN:
		nd.lhs.gen(vars, lb)
		fmt.Println("  pop rax")
		fmt.Println("  mov rsp, rbp")
		fmt.Println("  pop rbp")
		fmt.Println("  ret")
		return
	case ND_IF:
		lbIf := lb.get("if")
		nd.lhs.gen(vars, lb)
		fmt.Println("  pop rax")
		fmt.Println("  cmp rax, 0")
		fmt.Println("  je", lbIf)
		nd.rhs.gen(vars, lb)
		fmt.Println(lbIf + ":")
		return
	case ND_WHILE:
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
		return
	case ND_BLOCK:
		for _, nd1 := range nd.nds {
			nd1.gen(vars, lb)
			fmt.Println("  pop rax")
		}
		return
	case int('='):
		nd.lhs.genLval(vars)
		nd.rhs.gen(vars, lb)
		fmt.Println("  pop rdi")
		fmt.Println("  pop rax")
		fmt.Println("  mov [rax], rdi")
		fmt.Println("  push rdi")
		return
	}

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
