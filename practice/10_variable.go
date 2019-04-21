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
	tks = &tokens{}

	for c, err = rd.ReadByte(); err == nil; {
		switch c {
		case 0, byte(' '), byte('\n'):
			c, err = rd.ReadByte()
			continue
		case byte('+'), byte('-'), byte('*'), byte('/'), byte('('), byte(')'), byte(';'):
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
			tk := token{
				ty:    TK_IDENT,
				input: []byte{c},
			}
			tks.tks = append(tks.tks, tk)

			c, err = rd.ReadByte()
			continue
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
)

type node struct {
	ty, val  int
	name     []byte
	lhs, rhs *node
}

func program(tks *tokens) (nds []node) {
	for tks.current().ty != TK_EOF {
		nds = append(nds, *stmt(tks))
	}
	return
}

func stmt(tks *tokens) *node {
	nd := assign(tks)
	if !tks.consume(';') {
		log.Fatal("';'ではないトークンです:", string(tks.current().input))
	}
	return nd
}

func assign(tks *tokens) *node {
	nd := add(tks)
	if !tks.consume('=') {
		return nd
	}

	nd = &node{
		ty:  '=',
		lhs: nd,
		rhs: assign(tks),
	}
	return nd
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

func term(tks *tokens) *node {
	var nd *node
	tk := tks.current()

	switch {
	case tks.consume('('):
		nd = add(tks)
		if !tks.consume(')') {
			log.Fatal("閉じカッコがありません: ", string(tks.current().input))
		}
	case tk.ty == TK_NUM:
		nd = num(tks)
	case tk.ty == TK_IDENT:
		nd = ident(tks)
	default:
		log.Fatal("不正なトークンです: ", string(tk.input))
	}

	return nd
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
		ty:  ND_IDENT,
		name: tk.input,
	}
}

func gen(nd *node) {
	if nd.ty == ND_NUM {
		fmt.Println("  push", nd.val)
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

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := tokenize(rd)

	nds := program(tks)

	fmt.Println(".intel_syntax noprefix")
	fmt.Println(".global main")
	fmt.Println("main:")

	for _, nd := range nds {
		gen(&nd)
	}

	fmt.Println("  pop rax")
	fmt.Println("  ret")
}
