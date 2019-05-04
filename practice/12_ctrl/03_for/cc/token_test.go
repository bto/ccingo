package cc

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"testing"
)

func TestTokenizeAddSub(t *testing.T) {
	rd := newReader(" 1+  23\n -456 \n")
	tks := Tokenize(rd)
	if len(tks.tks) != 6 {
		t.Fatal("invalid number of tokens:", len(tks.tks))
	}
	if tk := tks.current(); !tk.checkNum(1) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('+') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkNum(23) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('-') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkNum(456) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); tk.ty != TK_EOF {
		t.Fatal("invalid token:", tk)
	}
}

func TestTokenizeMulDiv(t *testing.T) {
	rd := newReader("1*2/3")
	tks := Tokenize(rd)
	if len(tks.tks) != 6 {
		t.Fatal("invalid number of tokens:", len(tks.tks))
	}
	if tk := tks.current(); !tk.checkNum(1) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('*') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkNum(2) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('/') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkNum(3) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); tk.ty != TK_EOF {
		t.Fatal("invalid token:", tk)
	}
}

func TestTokenizeTerm(t *testing.T) {
	rd := newReader("(1+2)")
	tks := Tokenize(rd)
	if len(tks.tks) != 6 {
		t.Fatal("invalid number of tokens:", len(tks.tks))
	}
	if tk := tks.current(); !tk.checkChar('(') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkNum(1) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('+') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkNum(2) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar(')') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); tk.ty != TK_EOF {
		t.Fatal("invalid token:", tk)
	}
}

func TestTokenizeComp(t *testing.T) {
	rd := newReader("==!=<<=>>=")
	tks := Tokenize(rd)
	if len(tks.tks) != 7 {
		t.Fatal("invalid number of tokens:", len(tks.tks))
	}
	if tk := tks.current(); !tk.checkWord(TK_EQ, "==") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkWord(TK_NE, "!=") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('<') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkWord(TK_LE, "<=") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('>') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkWord(TK_GE, ">=") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); tk.ty != TK_EOF {
		t.Fatal("invalid token:", tk)
	}
}

func TestTokenizeVar(t *testing.T) {
	rd := newReader("a===b;")
	tks := Tokenize(rd)
	if len(tks.tks) != 6 {
		t.Fatal("invalid number of tokens:", len(tks.tks))
	}
	if tk := tks.current(); !tk.checkWord(TK_IDENT, "a") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkWord(TK_EQ, "==") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('=') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkWord(TK_IDENT, "b") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar(';') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); tk.ty != TK_EOF {
		t.Fatal("invalid token:", tk)
	}
}

func TestTokenizeWord(t *testing.T) {
	rd := newReader("return if while for")
	tks := Tokenize(rd)
	if len(tks.tks) != 5 {
		t.Fatal("invalid number of tokens:", len(tks.tks))
	}
	if tk := tks.current(); !tk.checkWord(TK_RETURN, "return") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkWord(TK_IF, "if") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkWord(TK_WHILE, "while") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkWord(TK_FOR, "for") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); tk.ty != TK_EOF {
		t.Fatal("invalid token:", tk)
	}
}

func TestTokenizeIdent(t *testing.T) {
	rd := newReader("foo;bar")
	tks := Tokenize(rd)
	if len(tks.tks) != 4 {
		t.Fatal("invalid number of tokens:", len(tks.tks))
	}
	if tk := tks.current(); !tk.checkWord(TK_IDENT, "foo") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar(';') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkWord(TK_IDENT, "bar") {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); tk.ty != TK_EOF {
		t.Fatal("invalid token:", tk)
	}
}

func TestTokenizeAlNum(t *testing.T) {
	rd := newReader("")
	name, c, err := tokenizeAlNum(rd, 'a')
	if string(name) != "a" {
		t.Fatal("invalid name:", string(name))
	}
	if c != 0 {
		t.Fatal("invalid next charactor:", c)
	}
	if err != io.EOF {
		t.Fatal("not EOF:", err)
	}

	rd = newReader("1a9=")
	name, c, err = tokenizeAlNum(rd, 'z')
	if string(name) != "z1a9" {
		t.Fatal("invalid name:", string(name))
	}
	if c != byte('=') {
		t.Fatal("invalid next charactor:", c)
	}
	if err != nil {
		t.Fatal("invalid error:", err)
	}
}

func TestTokenizeNum(t *testing.T) {
	rd := newReader("")
	tk, c, err := tokenizeNum(rd, '1')
	if !tk.checkNum(1) {
		t.Fatal("invalid token:", tk)
	}
	if c != 0 {
		t.Fatal("invalid next charactor:", c)
	}
	if err != io.EOF {
		t.Fatal("not EOF:", err)
	}

	rd = newReader("2a")
	tk, c, err = tokenizeNum(rd, '1')
	if !tk.checkNum(12) {
		t.Fatal("invalid token:", tk)
	}
	if c != byte('a') {
		t.Fatal("invalid next charactor:", c)
	}
	if err != nil {
		t.Fatal("invalid error:", err)
	}
}

func (tk *token) checkChar(op int) bool {
	return tk.ty == op && string(tk.input) == string(op)
}

func (tk *token) checkNum(val int) bool {
	return tk.ty == TK_NUM && tk.val == val && string(tk.input) == strconv.Itoa(val)
}

func (tk *token) checkWord(ty int, word string) bool {
	return tk.ty == ty && string(tk.input) == word
}

func newReader(str string) *bufio.Reader {
	return bufio.NewReader(bytes.NewReader([]byte(str)))
}
