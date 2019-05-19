package cc

import (
	"testing"
)

func TestParseNum(t *testing.T) {
	tks := &tokens{}
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: TK_EOF})
	nd := tks.Parse()
	if !nd.checkNum(1) {
		t.Fatal("invalid node:", nd)
	}
}

func TestParseAddSub(t *testing.T) {
	tks := &tokens{}
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: '+'})
	tks.append(token{ty: TK_NUM, val: 2})
	tks.append(token{ty: '-'})
	tks.append(token{ty: TK_NUM, val: 3})
	tks.append(token{ty: TK_EOF})
	nd := tks.Parse()
	if !nd.checkOp('-') {
		t.Fatal("invalid node:", nd)
	}
	ndr := nd.rhs
	if !ndr.checkNum(3) {
		t.Fatal("invalid node:", ndr)
	}
	ndl := nd.lhs
	if !ndl.checkOp('+') {
		t.Fatal("invalid node:", ndl)
	}
	ndlr := ndl.rhs
	if !ndlr.checkNum(2) {
		t.Fatal("invalid node:", ndr)
	}
	ndll := ndl.lhs
	if !ndll.checkNum(1) {
		t.Fatal("invalid node:", ndr)
	}
}

func TestParseMulDiv(t *testing.T) {
	tks := &tokens{}
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: '+'})
	tks.append(token{ty: TK_NUM, val: 2})
	tks.append(token{ty: '*'})
	tks.append(token{ty: TK_NUM, val: 3})
	tks.append(token{ty: '-'})
	tks.append(token{ty: TK_NUM, val: 4})
	tks.append(token{ty: '/'})
	tks.append(token{ty: TK_NUM, val: 5})
	tks.append(token{ty: TK_EOF})
	nd := tks.Parse()
	if !nd.checkOp('-') {
		t.Fatal("invalid node:", nd)
	}
	ndr := nd.rhs
	if !ndr.checkOp('/') {
		t.Fatal("invalid node:", ndr)
	}
	ndrr := ndr.rhs
	if !ndrr.checkNum(5) {
		t.Fatal("invalid node:", ndrr)
	}
	ndrl := ndr.lhs
	if !ndrl.checkNum(4) {
		t.Fatal("invalid node:", ndrl)
	}
	ndl := nd.lhs
	if !ndl.checkOp('+') {
		t.Fatal("invalid node:", ndl)
	}
	ndlr := ndl.rhs
	if !ndlr.checkOp('*') {
		t.Fatal("invalid node:", ndlr)
	}
	ndlrr := ndlr.rhs
	if !ndlrr.checkNum(3) {
		t.Fatal("invalid node:", ndlrr)
	}
	ndlrl := ndlr.lhs
	if !ndlrl.checkNum(2) {
		t.Fatal("invalid node:", ndlrl)
	}
	ndll := ndl.lhs
	if !ndll.checkNum(1) {
		t.Fatal("invalid node:", ndll)
	}
}

func TestParseTerm(t *testing.T) {
	tks := &tokens{}
	tks.append(token{ty: '('})
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: '+'})
	tks.append(token{ty: TK_NUM, val: 2})
	tks.append(token{ty: ')'})
	tks.append(token{ty: '*'})
	tks.append(token{ty: TK_NUM, val: 3})
	tks.append(token{ty: TK_EOF})
	nd := tks.Parse()
	if !nd.checkOp('*') {
		t.Fatal("invalid node:", nd)
	}
	ndr := nd.rhs
	if !ndr.checkNum(3) {
		t.Fatal("invalid node:", ndr)
	}
	ndl := nd.lhs
	if !ndl.checkOp('+') {
		t.Fatal("invalid node:", ndl)
	}
	ndlr := ndl.rhs
	if !ndlr.checkNum(2) {
		t.Fatal("invalid node:", ndr)
	}
	ndll := ndl.lhs
	if !ndll.checkNum(1) {
		t.Fatal("invalid node:", ndr)
	}
}

func TestParseUnary(t *testing.T) {
	tks := &tokens{}
	tks.append(token{ty: '+'})
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: '+'})
	tks.append(token{ty: '-'})
	tks.append(token{ty: TK_NUM, val: 2})
	tks.append(token{ty: TK_EOF})
	nd := tks.Parse()
	if !nd.checkOp('+') {
		t.Fatal("invalid node:", nd)
	}
	ndr := nd.rhs
	if !ndr.checkOp('-') {
		t.Fatal("invalid node:", ndr)
	}
	ndrr := ndr.rhs
	if !ndrr.checkNum(2) {
		t.Fatal("invalid node:", ndrr)
	}
	ndrl := ndr.lhs
	if !ndrl.checkNum(0) {
		t.Fatal("invalid node:", ndrl)
	}
	ndl := nd.lhs
	if !ndl.checkNum(1) {
		t.Fatal("invalid node:", ndl)
	}
}

func (nd *node) checkNum(val int) bool {
	return nd.ty == ND_NUM && nd.val == val && nd.lhs == nil && nd.rhs == nil
}

func (nd *node) checkOp(ty int) bool {
	return nd.ty == ty && nd.lhs != nil && nd.rhs != nil
}
