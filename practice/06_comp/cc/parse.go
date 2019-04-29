package cc

import (
	"log"
)

const (
	ND_NUM = iota + 256
	ND_EQ
	ND_NE
	ND_LE
)

type node struct {
	ty, val  int
	lhs, rhs *node
}

func Parse(tks *tokens) *node {
	return equality(tks)
}

func equality(tks *tokens) *node {
	nd := relational(tks)
	return equalityx(tks, nd)
}

func equalityx(tks *tokens, nd *node) *node {
	switch {
	case tks.consume(TK_EQ):
		ndRel := relational(tks)
		nd = &node{
			ty:  ND_EQ,
			lhs: nd,
			rhs: ndRel,
		}
		return equalityx(tks, nd)
	case tks.consume(TK_NE):
		ndRel := relational(tks)
		nd = &node{
			ty:  ND_NE,
			lhs: nd,
			rhs: ndRel,
		}
		return equalityx(tks, nd)
	default:
		return nd
	}
}

func relational(tks *tokens) *node {
	nd := add(tks)
	return relationalx(tks, nd)
}

func relationalx(tks *tokens, nd *node) *node {
	switch {
	case tks.consume('<'):
		ndAdd := add(tks)
		nd = &node{
			ty:  '<',
			lhs: nd,
			rhs: ndAdd,
		}
		return relationalx(tks, nd)
	case tks.consume(TK_LE):
		ndAdd := add(tks)
		nd = &node{
			ty:  ND_LE,
			lhs: nd,
			rhs: ndAdd,
		}
		return relationalx(tks, nd)
	case tks.consume('>'):
		ndAdd := add(tks)
		nd = &node{
			ty:  '<',
			lhs: ndAdd,
			rhs: nd,
		}
		return relationalx(tks, nd)
	case tks.consume(TK_GE):
		ndAdd := add(tks)
		nd = &node{
			ty:  ND_LE,
			lhs: ndAdd,
			rhs: nd,
		}
		return relationalx(tks, nd)
	default:
		return nd
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
	nd := unary(tks)
	return mulx(tks, nd)
}

func mulx(tks *tokens, nd *node) *node {
	switch {
	case tks.consume('*'):
		ndUnary := unary(tks)
		nd = &node{
			ty:  '*',
			lhs: nd,
			rhs: ndUnary,
		}
		return mulx(tks, nd)
	case tks.consume('/'):
		ndUnary := unary(tks)
		nd = &node{
			ty:  '/',
			lhs: nd,
			rhs: ndUnary,
		}
		return mulx(tks, nd)
	default:
		return nd
	}
}

func unary(tks *tokens) (nd *node) {
	switch {
	case tks.consume('+'):
		return term(tks)
	case tks.consume('-'):
		ndZero := &node{
			ty:  ND_NUM,
			val: 0,
		}
		ndTerm := term(tks)
		return &node{
			ty:  '-',
			lhs: ndZero,
			rhs: ndTerm,
		}
	default:
		return term(tks)
	}
}

func term(tks *tokens) *node {
	switch {
	case tks.consume('('):
		nd := equality(tks)
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
