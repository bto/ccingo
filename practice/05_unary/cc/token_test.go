package cc

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"testing"
)

func TestTokenize(t *testing.T) {
	rd := newReader(" 12-  3/(3\n +4)*5 \n")
	tks := Tokenize(rd)
	if len(tks.tks) != 12 {
		t.Fatal("invalid number of tokens:", len(tks.tks))
	}
	if tk := tks.current(); !tk.checkNum(12) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('-') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkNum(3) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('/') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('(') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkNum(3) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('+') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkNum(4) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar(')') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkChar('*') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); !tk.checkNum(5) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks.next(); tk.ty != TK_EOF {
		t.Fatal("invalid token:", tk)
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

func newReader(str string) *bufio.Reader {
	return bufio.NewReader(bytes.NewReader([]byte(str)))
}
