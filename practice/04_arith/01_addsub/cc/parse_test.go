package cc

import (
	"testing"
)

func TestParse(t *testing.T) {
	tks := &tokens{}
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: TK_EOF})
	nd := Parse(tks)
	if !checkNum(nd, 1) {
		t.Fatal("invalid node:", nd)
	}

	tks = &tokens{}
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: '+'})
	tks.append(token{ty: TK_NUM, val: 2})
	tks.append(token{ty: TK_EOF})
	nd = Parse(tks)
	if !checkOp(nd, '+') {
		t.Fatal("invalid node:", nd)
	}
	ndr := nd.rhs
	if !checkNum(ndr, 2) {
		t.Fatal("invalid node:", ndr)
	}
	ndl := nd.lhs
	if !checkNum(ndl, 1) {
		t.Fatal("invalid node:", ndl)
	}

	tks = &tokens{}
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: '+'})
	tks.append(token{ty: TK_NUM, val: 2})
	tks.append(token{ty: '-'})
	tks.append(token{ty: TK_NUM, val: 3})
	tks.append(token{ty: TK_EOF})
	nd = Parse(tks)
	if !checkOp(nd, '-') {
		t.Fatal("invalid node:", nd)
	}
	ndr = nd.rhs
	if !checkNum(ndr, 3) {
		t.Fatal("invalid node:", ndr)
	}
	ndl = nd.lhs
	if !checkOp(ndl, '+') {
		t.Fatal("invalid node:", ndl)
	}
	ndlr := ndl.rhs
	if !checkNum(ndlr, 2) {
		t.Fatal("invalid node:", ndr)
	}
	ndll := ndl.lhs
	if !checkNum(ndll, 1) {
		t.Fatal("invalid node:", ndr)
	}
}

func checkNum(nd *node, val int) bool {
	return nd.ty == ND_NUM && nd.val == val && nd.lhs == nil && nd.rhs == nil
}

func checkOp(nd *node, ty int) bool {
	return nd.ty == ty && nd.lhs != nil && nd.rhs != nil
}
