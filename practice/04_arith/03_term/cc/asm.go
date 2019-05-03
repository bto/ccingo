package cc

import (
	"fmt"
)

func (nd *node) PrintAsm() {
	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	nd.gen()

	fmt.Println("  pop rax")
	fmt.Println("  ret")
}

func (nd *node) gen() {
	if nd.ty == ND_NUM {
		nd.genNum()
	} else {
		nd.genOp()
	}
}

func (nd *node) genNum() {
	fmt.Println("  push", nd.val)
}

func (nd *node) genOp() {
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
	}

	fmt.Println("  push rax")
}
