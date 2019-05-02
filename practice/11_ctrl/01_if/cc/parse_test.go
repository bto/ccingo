package cc

import (
	"testing"
)

func TestParseNum(t *testing.T) {
	tks := &tokens{}
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: ';'}).append(token{ty: TK_EOF})
	nd := tks.Parse()[0]
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
	tks.append(token{ty: ';'}).append(token{ty: TK_EOF})
	nd := tks.Parse()[0]
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
	tks.append(token{ty: ';'}).append(token{ty: TK_EOF})
	nd := tks.Parse()[0]
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
	tks.append(token{ty: ';'}).append(token{ty: TK_EOF})
	nd := tks.Parse()[0]
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
	tks.append(token{ty: ';'}).append(token{ty: TK_EOF})
	nd := tks.Parse()[0]
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

func TestParseComp(t *testing.T) {
	// 1 == 2 < 3 != 4 <= 5 != 6 > 7 == 8 >= 9
	tks := &tokens{}
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: TK_EQ})
	tks.append(token{ty: TK_NUM, val: 2})
	tks.append(token{ty: '<'})
	tks.append(token{ty: TK_NUM, val: 3})
	tks.append(token{ty: TK_NE})
	tks.append(token{ty: TK_NUM, val: 4})
	tks.append(token{ty: TK_LE})
	tks.append(token{ty: TK_NUM, val: 5})
	tks.append(token{ty: TK_NE})
	tks.append(token{ty: TK_NUM, val: 6})
	tks.append(token{ty: '>'})
	tks.append(token{ty: TK_NUM, val: 7})
	tks.append(token{ty: TK_EQ})
	tks.append(token{ty: TK_NUM, val: 8})
	tks.append(token{ty: TK_GE})
	tks.append(token{ty: TK_NUM, val: 9})
	tks.append(token{ty: ';'}).append(token{ty: TK_EOF})
	nd := tks.Parse()[0]
	if !nd.checkOp(ND_EQ) {
		t.Fatal("invalid node:", nd)
	}
	ndr := nd.rhs
	if !ndr.checkOp(ND_LE) {
		t.Fatal("invalid node:", ndr)
	}
	ndrr := ndr.rhs
	if !ndrr.checkNum(8) {
		t.Fatal("invalid node:", ndrr)
	}
	ndrl := ndr.lhs
	if !ndrl.checkNum(9) {
		t.Fatal("invalid node:", ndrl)
	}
	ndl := nd.lhs
	if !ndl.checkOp(ND_NE) {
		t.Fatal("invalid node:", ndl)
	}
	ndlr := ndl.rhs
	if !ndlr.checkOp('<') {
		t.Fatal("invalid node:", ndlr)
	}
	ndlrr := ndlr.rhs
	if !ndlrr.checkNum(6) {
		t.Fatal("invalid node:", ndlrr)
	}
	ndlrl := ndlr.lhs
	if !ndlrl.checkNum(7) {
		t.Fatal("invalid node:", ndlrl)
	}
	ndll := ndl.lhs
	if !ndll.checkOp(ND_NE) {
		t.Fatal("invalid node:", ndll)
	}
	ndllr := ndll.rhs
	if !ndllr.checkOp(ND_LE) {
		t.Fatal("invalid node:", ndllr)
	}
	ndllrr := ndllr.rhs
	if !ndllrr.checkNum(5) {
		t.Fatal("invalid node:", ndllrr)
	}
	ndllrl := ndllr.lhs
	if !ndllrl.checkNum(4) {
		t.Fatal("invalid node:", ndllrl)
	}
	ndlll := ndll.lhs
	if !ndlll.checkOp(ND_EQ) {
		t.Fatal("invalid node:", ndlll)
	}
	ndlllr := ndlll.rhs
	if !ndlllr.checkOp('<') {
		t.Fatal("invalid node:", ndlllr)
	}
	ndlllrr := ndlllr.rhs
	if !ndlllrr.checkNum(3) {
		t.Fatal("invalid node:", ndlllrr)
	}
	ndlllrl := ndlllr.lhs
	if !ndlllrl.checkNum(2) {
		t.Fatal("invalid node:", ndlllrl)
	}
	ndllll := ndlll.lhs
	if !ndllll.checkNum(1) {
		t.Fatal("invalid node:", ndllll)
	}
}

func TestParseVar(t *testing.T) {
	// a=b=1==2;a;
	tks := &tokens{}
	tks.append(token{ty: TK_IDENT, input: []byte("a")})
	tks.append(token{ty: '='})
	tks.append(token{ty: TK_IDENT, input: []byte("b")})
	tks.append(token{ty: '='})
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: TK_EQ})
	tks.append(token{ty: TK_NUM, val: 2})
	tks.append(token{ty: ';'})
	tks.append(token{ty: TK_IDENT, input: []byte("a")})
	tks.append(token{ty: ';'}).append(token{ty: TK_EOF})
	nds := tks.Parse()

	nd := nds[0]
	if !nd.checkOp('=') {
		t.Fatal("invalid node:", nd)
	}
	ndr := nd.rhs
	if !ndr.checkOp('=') {
		t.Fatal("invalid node:", ndr)
	}
	ndrr := ndr.rhs
	if !ndrr.checkOp(ND_EQ) {
		t.Fatal("invalid node:", ndrr)
	}
	ndrrr := ndrr.rhs
	if !ndrrr.checkNum(2) {
		t.Fatal("invalid node:", ndrrr)
	}
	ndrrl := ndrr.lhs
	if !ndrrl.checkNum(1) {
		t.Fatal("invalid node:", ndrrl)
	}
	ndrl := ndr.lhs
	if !ndrl.checkVar("b") {
		t.Fatal("invalid node:", ndrl)
	}
	ndl := nd.lhs
	if !ndl.checkVar("a") {
		t.Fatal("invalid node:", ndl)
	}

	nd = nds[1]
	if !nd.checkVar("a") {
		t.Fatal("invalid node:", nd)
	}
}

func TestParseReturn(t *testing.T) {
	// return a=1;
	tks := &tokens{}
	tks.append(token{ty: TK_RETURN})
	tks.append(token{ty: TK_IDENT, input: []byte("a")})
	tks.append(token{ty: '='})
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: ';'}).append(token{ty: TK_EOF})
	nd := tks.Parse()[0]
	if !nd.checkUnary(ND_RETURN) {
		t.Fatal("invalid node:", nd)
	}
	ndl := nd.lhs
	if !ndl.checkOp('=') {
		t.Fatal("invalid node:", ndl)
	}
	ndlr := ndl.rhs
	if !ndlr.checkNum(1) {
		t.Fatal("invalid node:", ndlr)
	}
	ndll := ndl.lhs
	if !ndll.checkVar("a") {
		t.Fatal("invalid node:", ndll)
	}
}

func TestParseIf(t *testing.T) {
	// if(a=1)return 2;
	tks := &tokens{}
	tks.append(token{ty: TK_IF})
	tks.append(token{ty: '('})
	tks.append(token{ty: TK_IDENT, input: []byte("a")})
	tks.append(token{ty: '='})
	tks.append(token{ty: TK_NUM, val: 1})
	tks.append(token{ty: ')'})
	tks.append(token{ty: TK_RETURN})
	tks.append(token{ty: TK_NUM, val: 2})
	tks.append(token{ty: ';'}).append(token{ty: TK_EOF})
	nd := tks.Parse()[0]
	if !nd.checkOp(ND_IF) {
		t.Fatal("invalid node:", nd)
	}
	ndr := nd.rhs
	if !ndr.checkUnary(ND_RETURN) {
		t.Fatal("invalid node:", ndr)
	}
	ndrl := ndr.lhs
	if !ndrl.checkNum(2) {
		t.Fatal("invalid node:", ndrl)
	}
	ndl := nd.lhs
	if !ndl.checkOp('=') {
		t.Fatal("invalid node:", ndl)
	}
	ndlr := ndl.rhs
	if !ndlr.checkNum(1) {
		t.Fatal("invalid node:", ndlr)
	}
	ndll := ndl.lhs
	if !ndll.checkVar("a") {
		t.Fatal("invalid node:", ndll)
	}
}

func (nd *node) checkNum(val int) bool {
	return nd.ty == ND_NUM && nd.val == val && nd.lhs == nil && nd.rhs == nil
}

func (nd *node) checkOp(ty int) bool {
	return nd.ty == ty && nd.lhs != nil && nd.rhs != nil
}

func (nd *node) checkUnary(ty int) bool {
	return nd.ty == ty && nd.lhs != nil && nd.rhs == nil
}

func (nd *node) checkVar(name string) bool {
	return nd.ty == ND_VAR && nd.name == name
}
