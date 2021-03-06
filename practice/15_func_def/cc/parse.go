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
	ND_BLOCK
	ND_FUNC_CALL
	ND_FUNC_DEF
)

type node struct {
	ty, val  int
	name     string
	lhs, rhs *node
	nds      nodes
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
		nds = append(nds, *tks.funcDef())
	}
	return
}

func (tks *tokens) funcDef() *node {
	tk := tks.current()
	if tk.ty != TK_IDENT {
		log.Fatal("関数定義ではありません: ", string(tk.input))
	}
	name := string(tk.input)
	tks.next()

	if !tks.consume('(') {
		log.Fatal("関数定義の開きカッコではありません: ", string(tks.current().input))
	}

	nds := tks.funcDefArgs()

	if !tks.consume(')') {
		log.Fatal("関数定義の閉じカッコではありません: ", string(tks.current().input))
	}

	nd := tks.block()

	return &node{
		ty:   ND_FUNC_DEF,
		name: name,
		nds:  nds,
		lhs:  nd,
	}
}

func (tks *tokens) funcDefArgs() (nds nodes) {
	switch tk := tks.current(); tk.ty {
	case ')':
		return
	case TK_IDENT:
		nd := node{
			ty:   ND_VAR,
			name: string(tk.input),
		}
		nds = append(nds, nd)
		tks.next()
	default:
		log.Fatal("関数定義の変数ではありません: ", string(tk.input))
	}

	for tks.current().ty != ')' {
		if !tks.consume(',') {
			break
		}

		tk := tks.current()
		if tk.ty != TK_IDENT {
			log.Fatal("関数定義の変数ではありません: ", string(tk.input))
		}
		nd := node{
			ty:   ND_VAR,
			name: string(tk.input),
		}
		nds = append(nds, nd)
		tks.next()
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
	case TK_IF, TK_WHILE, TK_FOR:
		return tks.control()
	case '{':
		return tks.block()
	default:
		nd := tks.assign()
		if !tks.consume(';') {
			log.Fatal("';'ではないトークンです: ", string(tks.current().input))
		}
		return nd
	}
}

func (tks *tokens) block() *node {
	if !tks.consume('{') {
		log.Fatal("ブロックの開きカッコがありません: ", string(tks.current().input))
	}
	nds := tks.blockItems()
	if !tks.consume('}') {
		log.Fatal("ブロックの閉じカッコがありません: ", string(tks.current().input))
	}
	return &node{
		ty:  ND_BLOCK,
		nds: nds,
	}
}

func (tks *tokens) blockItems() (nds nodes) {
	for tks.current().ty != '}' {
		nds = append(nds, *tks.stmt())
	}
	return
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
	case tks.consume(TK_FOR):
		if !tks.consume('(') {
			log.Fatal("forの開きカッコがありません: ", string(tks.current().input))
		}
		ndAssign1 := tks.assign()
		if !tks.consume(';') {
			log.Fatal("';'ではないトークンです:", string(tks.current().input))
		}
		ndAssign2 := tks.assign()
		if !tks.consume(';') {
			log.Fatal("';'ではないトークンです:", string(tks.current().input))
		}
		ndAssign3 := tks.assign()
		if !tks.consume(')') {
			log.Fatal("forの閉じカッコがありません: ", string(tks.current().input))
		}
		ndStmt := tks.stmt()
		return &node{
			lhs: ndAssign1,
			rhs: &node{
				ty:  ND_WHILE,
				lhs: ndAssign2,
				rhs: &node{
					lhs: ndStmt,
					rhs: ndAssign3,
				},
			},
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
		name := string(tk.input)
		tks.next()
		if tks.consume('(') {
			return tks.funcCall(name)
		} else {
			return &node{
				ty:   ND_VAR,
				name: name,
			}
		}
	}

	log.Fatal("不正なトークンです: ", string(tk.input))
	return
}

func (tks *tokens) funcCall(name string) *node {
	nds := tks.funcCallArgs()
	if !tks.consume(')') {
		log.Fatal("関数の閉じカッコがありません: ", string(tks.current().input))
	}
	return &node{
		ty:   ND_FUNC_CALL,
		name: name,
		nds:  nds,
	}
}

func (tks *tokens) funcCallArgs() (nds nodes) {
	if tks.current().ty == ')' {
		return
	}

	nds = append(nds, *tks.assign())

	for tks.current().ty != ')' {
		if !tks.consume(',') {
			break
		}
		nds = append(nds, *tks.assign())
	}
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
