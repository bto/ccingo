package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	TK_NUM = iota + 256
	TK_IDENT
	TK_RETURN
	TK_EOF
)

type token struct {
	ty, val int
	input   []byte
}

type tokens struct {
	tks []token
	i   int
}

func (tks *tokens) append(tk token) {
	tks.tks = append(tks.tks, tk)
}

func (tks *tokens) consume(ty int) bool {
	if tks.tks[tks.i].ty == ty {
		tks.i++
		return true
	} else {
		return false
	}
}

func (tks *tokens) current() token {
	return tks.tks[tks.i]
}

func (tks *tokens) next() token {
	tks.i++
	return tks.tks[tks.i]
}

func tokenize(rd *bufio.Reader) (tks *tokens) {
	var c byte
	var err error
	var tk token
	var name []byte
	tks = &tokens{}

	for c, err = rd.ReadByte(); err == nil; {
		switch c {
		case 0, byte(' '), byte('\n'):
			c, err = rd.ReadByte()
			continue
		case byte('+'), byte('-'), byte('*'), byte('/'), byte('('), byte(')'), byte('='), byte(';'):
			tk := token{
				ty:    int(c),
				input: []byte{c},
			}
			tks.append(tk)

			c, err = rd.ReadByte()
			continue
		}

		if byte('0') <= c && c <= byte('9') {
			tk, c, err = tokenizeNum(rd, c)
			tks.append(tk)
			continue
		}

		if byte('a') <= c && c <= byte('z') {
			name, c, err = tokenizeAlphabet(rd, c)
			tk := token{
				input: name,
			}

			if len(name) == 1 {
				tk.ty = TK_IDENT
				tks.append(tk)
				continue
			}

			switch string(name) {
			case "return":
				tk.ty = TK_RETURN
				tks.append(tk)
				continue
			}

			log.Fatal("トークナイズできません: ", name)
		}

		log.Fatal("トークナイズできません: ", string([]byte{c}))
	}
	if err != io.EOF {
		log.Fatal(err)
	}

	tk = token{
		ty: TK_EOF,
	}
	tks.append(tk)

	return
}

func tokenizeAlphabet(rd *bufio.Reader, v byte) (name []byte, c byte, err error) {
	for c = v; err == nil; c, err = rd.ReadByte() {
		if c < byte('a') || byte('z') < c {
			break
		}

		name = append(name, c)
	}

	return
}

func tokenizeNum(rd *bufio.Reader, v byte) (tk token, c byte, err error) {
	var num []byte
	for c = v; err == nil; c, err = rd.ReadByte() {
		if c < byte('0') || byte('9') < c {
			break
		}

		num = append(num, c)
	}

	val, err := strconv.Atoi(string(num))
	if err != nil {
		log.Fatal(err)
	}

	tk = token{
		ty:    TK_NUM,
		val:   val,
		input: num,
	}

	return
}

const (
	ND_NUM = iota + 256
	ND_IDENT
	ND_RETURN
)

type node struct {
	ty, val  int
	name     string
	lhs, rhs *node
}

func program(tks *tokens) (nds []node) {
	for tks.current().ty != TK_EOF {
		nds = append(nds, *stmt(tks))
	}
	return
}

func stmt(tks *tokens) (nd *node) {
	if tks.consume(TK_RETURN) {
		ndAssign := assign(tks)
		nd = &node{
			ty:  ND_RETURN,
			lhs: ndAssign,
		}
	} else {
		nd = assign(tks)
	}

	if !tks.consume(';') {
		log.Fatal("';'ではないトークンです:", string(tks.current().input))
	}

	return
}

func assign(tks *tokens) *node {
	nd := add(tks)

	if !tks.consume('=') {
		return nd
	}

	ndAssign := assign(tks)
	return &node{
		ty:  '=',
		lhs: nd,
		rhs: ndAssign,
	}
}

func add(tks *tokens) *node {
	nd := mul(tks)
	return addx(tks, nd)
}

func addx(tks *tokens, nd *node) *node {
	switch {
	case tks.consume('+'):
		ndMul := mul(tks)
		nd = &node{
			ty:  '+',
			lhs: nd,
			rhs: ndMul,
		}
		return addx(tks, nd)
	case tks.consume('-'):
		ndMul := mul(tks)
		nd = &node{
			ty:  '-',
			lhs: nd,
			rhs: ndMul,
		}
		return addx(tks, nd)
	default:
		return nd
	}
}

func mul(tks *tokens) *node {
	nd := term(tks)
	return mulx(tks, nd)
}

func mulx(tks *tokens, nd *node) *node {
	switch {
	case tks.consume('*'):
		ndTerm := term(tks)
		nd = &node{
			ty:  '*',
			lhs: nd,
			rhs: ndTerm,
		}
		return mulx(tks, nd)
	case tks.consume('/'):
		ndTerm := term(tks)
		nd = &node{
			ty:  '/',
			lhs: nd,
			rhs: ndTerm,
		}
		return mulx(tks, nd)
	default:
		return nd
	}
}

func term(tks *tokens) (nd *node) {
	tk := tks.current()

	switch {
	case tks.consume('('):
		nd = assign(tks)
		if !tks.consume(')') {
			log.Fatal("閉じカッコがありません: ", string(tks.current().input))
		}
		return
	case tk.ty == TK_NUM:
		return num(tks)
	case tk.ty == TK_IDENT:
		return ident(tks)
	}

	log.Fatal("不正なトークンです: ", string(tk.input))
	return
}

func num(tks *tokens) *node {
	tk := tks.current()
	if tk.ty != TK_NUM {
		log.Fatal("数値ではないトークンです: ", string(tk.input))
	}

	tks.next()
	return &node{
		ty:  ND_NUM,
		val: tk.val,
	}
}

func ident(tks *tokens) *node {
	tk := tks.current()
	if tk.ty != TK_IDENT {
		log.Fatal("変数ではないトークンです: ", string(tk.input))
	}

	tks.next()
	return &node{
		ty:   ND_IDENT,
		name: string(tk.input),
	}
}

func gen(nd *node) {
	switch nd.ty {
	case ND_NUM:
		fmt.Println("  push", nd.val)
		return
	case ND_IDENT:
		genLval(nd)
		fmt.Println("  pop rax")
		fmt.Println("  mov rax, [rax]")
		fmt.Println("  push rax")
		return
	case ND_RETURN:
		gen(nd.lhs)
		fmt.Println("  pop rax")
		fmt.Println("  mov rsp, rbp")
		fmt.Println("  pop rbp")
		fmt.Println("  ret")
		return
	case int('='):
		genLval(nd.lhs)
		gen(nd.rhs)
		fmt.Println("  pop rdi")
		fmt.Println("  pop rax")
		fmt.Println("  mov [rax], rdi")
		fmt.Println("  push rdi")
		return
	}

	gen(nd.lhs)
	gen(nd.rhs)

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

func genLval(nd *node) {
	if nd.ty != ND_IDENT {
		log.Fatal("代入の左辺値が変数ではありません")
	}

	offset := (byte('z') - nd.name[0] + 1) * 8
	fmt.Println("  mov rax, rbp")
	fmt.Println("  sub rax,", offset)
	fmt.Println("  push rax")
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := tokenize(rd)

	nds := program(tks)

	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	fmt.Println("  push rbp")
	fmt.Println("  mov rbp, rsp")
	fmt.Println("  sub rsp, 208")

	for _, nd := range nds {
		gen(&nd)
		fmt.Println("  pop rax")
	}

	fmt.Println("  mov rsp, rbp")
	fmt.Println("  pop rbp")
	fmt.Println("  ret")
}
