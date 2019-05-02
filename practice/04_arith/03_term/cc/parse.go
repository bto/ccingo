package cc

import (
	"log"
)

const (
	ND_NUM = iota + 256
)

type node struct {
	ty, val  int
	lhs, rhs *node
}

func Parse(tks *tokens) *node {
	nd := add(tks)
	if !tks.consume(TK_EOF) {
		log.Fatal("不正なトークンです: ", string(tks.current().input))
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
	switch {
	case tks.consume('('):
		nd := add(tks)
		if !tks.consume(')') {
			log.Fatal("閉じカッコがありません: ", string(tks.current().input))
		}
		return nd
	default:
		return num(tks)
	}
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
