package cc

import (
	"log"
)

const (
	ND_NUM = iota + 256
	ND_EQ
	ND_NE
	ND_LE
	ND_VAR
	ND_RETURN
	ND_IF
	ND_WHILE
)

type node struct {
	ty, val  int
	name     string
	lhs, rhs *node
}

type nodes []node

func (tks *tokens) Parse() nodes {
	nds := tks.program()
	if !tks.consume(TK_EOF) {
		log.Fatal("不正なトークンです: ", string(tks.current().input))
	}
	return nds
}

func (tks *tokens) program() (nds nodes) {
	for tks.current().ty != TK_EOF {
		nds = append(nds, *tks.stmt())
	}
	return
}

func (tks *tokens) stmt() *node {
	tk := tks.current()
	switch tk.ty {
	case TK_RETURN:
		tks.next()
		ndAssign := tks.assign()
		if !tks.consume(';') {
			log.Fatal("';'ではないトークンです: ", string(tks.current().input))
		}
		return &node{
			ty:  ND_RETURN,
			lhs: ndAssign,
		}
	case TK_IF, TK_WHILE:
		return tks.control()
	default:
		nd := tks.assign()
		if !tks.consume(';') {
			log.Fatal("';'ではないトークンです: ", string(tks.current().input))
		}
		return nd
	}
}

func (tks *tokens) control() *node {
	switch {
	case tks.consume(TK_IF):
		if !tks.consume('(') {
			log.Fatal("ifの開きカッコがありません: ", string(tks.current().input))
		}
		ndAssign := tks.assign()
		if !tks.consume(')') {
			log.Fatal("ifの閉じカッコがありません: ", string(tks.current().input))
		}
		ndStmt := tks.stmt()
		return &node{
			ty:  ND_IF,
			lhs: ndAssign,
			rhs: ndStmt,
		}
	case tks.consume(TK_WHILE):
		if !tks.consume('(') {
			log.Fatal("whileの開きカッコがありません: ", string(tks.current().input))
		}
		ndAssign := tks.assign()
		if !tks.consume(')') {
			log.Fatal("whileの閉じカッコがありません: ", string(tks.current().input))
		}
		ndStmt := tks.stmt()
		return &node{
			ty:  ND_WHILE,
			lhs: ndAssign,
			rhs: ndStmt,
		}
	}

	log.Fatal("不正なトークンです: ", string(tks.current().input))
	return &node{}
}

func (tks *tokens) assign() *node {
	nd := tks.equality()

	if !tks.consume('=') {
		return nd
	}

	ndEq := tks.assign()
	return &node{
		ty:  '=',
		lhs: nd,
		rhs: ndEq,
	}
}

func (tks *tokens) equality() *node {
	nd := tks.relational()
	return tks.equalityx(nd)
}

func (tks *tokens) equalityx(nd *node) *node {
	switch {
	case tks.consume(TK_EQ):
		ndRel := tks.relational()
		nd = &node{
			ty:  ND_EQ,
			lhs: nd,
			rhs: ndRel,
		}
		return tks.equalityx(nd)
	case tks.consume(TK_NE):
		ndRel := tks.relational()
		nd = &node{
			ty:  ND_NE,
			lhs: nd,
			rhs: ndRel,
		}
		return tks.equalityx(nd)
	default:
		return nd
	}
}

func (tks *tokens) relational() *node {
	nd := tks.add()
	return tks.relationalx(nd)
}

func (tks *tokens) relationalx(nd *node) *node {
	switch {
	case tks.consume('<'):
		ndAdd := tks.add()
		nd = &node{
			ty:  '<',
			lhs: nd,
			rhs: ndAdd,
		}
		return tks.relationalx(nd)
	case tks.consume(TK_LE):
		ndAdd := tks.add()
		nd = &node{
			ty:  ND_LE,
			lhs: nd,
			rhs: ndAdd,
		}
		return tks.relationalx(nd)
	case tks.consume('>'):
		ndAdd := tks.add()
		nd = &node{
			ty:  '<',
			lhs: ndAdd,
			rhs: nd,
		}
		return tks.relationalx(nd)
	case tks.consume(TK_GE):
		ndAdd := tks.add()
		nd = &node{
			ty:  ND_LE,
			lhs: ndAdd,
			rhs: nd,
		}
		return tks.relationalx(nd)
	default:
		return nd
	}
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

func (tks *tokens) term() (nd *node) {
	tk := tks.current()

	switch {
	case tks.consume('('):
		nd = tks.assign()
		if !tks.consume(')') {
			log.Fatal("閉じカッコがありません: ", string(tks.current().input))
		}
		return
	case tk.ty == TK_NUM:
		return tks.num()
	case tk.ty == TK_IDENT:
		return tks.ident()
	}

	log.Fatal("不正なトークンです: ", string(tk.input))
	return
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

func (tks *tokens) ident() *node {
	tk := tks.current()
	if tk.ty != TK_IDENT {
		log.Fatal("変数ではないトークンです: ", string(tk.input))
	}

	tks.next()
	return &node{
		ty:   ND_VAR,
		name: string(tk.input),
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
