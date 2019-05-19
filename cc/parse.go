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
	nd := tks.num()
	return tks.addx(nd)
}

func (tks *tokens) addx(nd *node) *node {
	switch {
	case tks.consume('+'):
		ndNum := tks.num()
		nd = &node{
			ty:  '+',
			lhs: nd,
			rhs: ndNum,
		}
		return tks.addx(nd)
	case tks.consume('-'):
		ndNum := tks.num()
		nd = &node{
			ty:  '-',
			lhs: nd,
			rhs: ndNum,
		}
		return tks.addx(nd)
	default:
		return nd
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
