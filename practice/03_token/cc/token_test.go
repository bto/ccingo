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
	if len(tks) != 6 {
		t.Fatal("invalid number of tokens:", len(tks))
	}
	if tk := tks[0]; !tk.checkNum(1) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks[1]; !tk.checkChar('+') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks[2]; !tk.checkNum(23) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks[3]; !tk.checkChar('-') {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks[4]; !tk.checkNum(456) {
		t.Fatal("invalid token:", tk)
	}
	if tk := tks[5]; tk.ty != TK_EOF {
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
