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

	for _, nd := range nds {
		nd.gen()
		fmt.Println("  pop rax")
	}

	fmt.Println("  mov rsp, rbp")
	fmt.Println("  pop rbp")
	fmt.Println("  ret")
}

func (nd *node) gen() {
	switch nd.ty {
	case ND_NUM:
		fmt.Println("  push", nd.val)
		return
	case ND_VAR:
		nd.genLval()
		fmt.Println("  pop rax")
		fmt.Println("  mov rax, [rax]")
		fmt.Println("  push rax")
		return
	case int('='):
		nd.lhs.genLval()
		nd.rhs.gen()
		fmt.Println("  pop rdi")
		fmt.Println("  pop rax")
		fmt.Println("  mov [rax], rdi")
		fmt.Println("  push rdi")
		return
	}

	nd.lhs.gen()
	nd.rhs.gen()

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

func (nd *node) genLval() {
	if nd.ty != ND_VAR {
		log.Fatal("代入の左辺値が変数ではありません")
	}

	offset := (byte('z') - nd.name[0] + 1) * 8
	fmt.Println("  mov rax, rbp")
	fmt.Println("  sub rax,", offset)
	fmt.Println("  push rax")
}
