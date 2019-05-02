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
		gen(&nd, vars, lb)
		fmt.Println("  pop rax")
	}

	fmt.Println("  mov rsp, rbp")
	fmt.Println("  pop rbp")
	fmt.Println("  ret")
}

func gen(nd *node, vars *variables, lb *label) {
	switch nd.ty {
	case ND_NUM:
		fmt.Println("  push", nd.val)
		return
	case ND_VAR:
		genLval(nd, vars)
		fmt.Println("  pop rax")
		fmt.Println("  mov rax, [rax]")
		fmt.Println("  push rax")
		return
	case ND_RETURN:
		gen(nd.lhs, vars, lb)
		fmt.Println("  pop rax")
		fmt.Println("  mov rsp, rbp")
		fmt.Println("  pop rbp")
		fmt.Println("  ret")
		return
	case ND_IF:
		lbIf := lb.get("if")
		gen(nd.lhs, vars, lb)
		fmt.Println("  pop rax")
		fmt.Println("  cmp rax, 0")
		fmt.Println("  je", lbIf)
		gen(nd.rhs, vars, lb)
		fmt.Println(lbIf + ":")
		return
	case ND_WHILE:
		lbBegin := lb.get("begin")
		lbEnd := lb.get("end")
		fmt.Println(lbBegin + ":")
		gen(nd.lhs, vars, lb)
		fmt.Println("  pop rax")
		fmt.Println("  cmp rax, 0")
		fmt.Println("  je", lbEnd)
		gen(nd.rhs, vars, lb)
		fmt.Println("  jmp", lbBegin)
		fmt.Println(lbEnd + ":")
		return
	case int('='):
		genLval(nd.lhs, vars)
		gen(nd.rhs, vars, lb)
		fmt.Println("  pop rdi")
		fmt.Println("  pop rax")
		fmt.Println("  mov [rax], rdi")
		fmt.Println("  push rdi")
		return
	}

	gen(nd.lhs, vars, lb)
	gen(nd.rhs, vars, lb)

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

func genLval(nd *node, vars *variables) {
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
