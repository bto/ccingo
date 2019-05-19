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

func (nd *node) checkNum(val int) bool {
	return nd.ty == ND_NUM && nd.val == val && nd.lhs == nil && nd.rhs == nil
}

func (nd *node) checkOp(ty int) bool {
	return nd.ty == ty && nd.lhs != nil && nd.rhs != nil
}
