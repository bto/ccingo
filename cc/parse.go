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

func (tks *tokens) Parse() *node {
	nd := tks.add()
	if !tks.consume(TK_EOF) {
		log.Fatal("不正なトークンです: ", string(tks.current().input))
	}
	return nd
}

func (tks *tokens) add() *node {
	nd := tks.mul()
	return tks.addx(nd)
}

func (tks *tokens) addx(nd *node) *node {
	switch {
	case tks.consume('+'):
		ndMul := tks.mul()
		nd = &node{
			ty:  '+',
			lhs: nd,
			rhs: ndMul,
		}
		return tks.addx(nd)
	case tks.consume('-'):
		ndMul := tks.mul()
		nd = &node{
			ty:  '-',
			lhs: nd,
			rhs: ndMul,
		}
		return tks.addx(nd)
	default:
		return nd
	}
}

func (tks *tokens) mul() *node {
	nd := tks.unary()
	return tks.mulx(nd)
}

func (tks *tokens) mulx(nd *node) *node {
	switch {
	case tks.consume('*'):
		ndUnary := tks.unary()
		nd = &node{
			ty:  '*',
			lhs: nd,
			rhs: ndUnary,
		}
		return tks.mulx(nd)
	case tks.consume('/'):
		ndUnary := tks.unary()
		nd = &node{
			ty:  '/',
			lhs: nd,
			rhs: ndUnary,
		}
		return tks.mulx(nd)
	default:
		return nd
	}
}

func (tks *tokens) unary() (nd *node) {
	switch {
	case tks.consume('+'):
		return tks.term()
	case tks.consume('-'):
		ndZero := &node{
			ty:  ND_NUM,
			val: 0,
		}
		ndTerm := tks.term()
		return &node{
			ty:  '-',
			lhs: ndZero,
			rhs: ndTerm,
		}
	default:
		return tks.term()
	}
}

func (tks *tokens) term() *node {
	switch {
	case tks.consume('('):
		nd := tks.add()
		if !tks.consume(')') {
			log.Fatal("閉じカッコがありません: ", string(tks.current().input))
		}
		return nd
	default:
		return tks.num()
	}
}

func (tks *tokens) num() *node {
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
