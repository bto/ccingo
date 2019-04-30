package cc

import (
	"bufio"
	"bytes"
	"io"
	"strconv"
	"testing"
)

func TestTokenize(t *testing.T) {
	rd := createReader([]byte(" 1+  23\n -456 \n"))
	tks := Tokenize(rd)
	if len(tks) != 6 {
		t.Fatal("invalid number of tokens:", len(tks))
	}
	if tk := tks[0]; !tk.checkNum(1) {
		t.Fatal("invalid tokens[0]:", tk)
	}
	if tk := tks[1]; !tk.checkOp('+') {
		t.Fatal("invalid tokens[1]:", tk)
	}
	if tk := tks[2]; !tk.checkNum(23) {
		t.Fatal("invalid tokens[2]:", tk)
	}
	if tk := tks[3]; !tk.checkOp('-') {
		t.Fatal("invalid tokens[3]:", tk)
	}
	if tk := tks[4]; !tk.checkNum(456) {
		t.Fatal("invalid tokens[4]:", tk)
	}
	if tk := tks[5]; tk.ty != TK_EOF {
		t.Fatal("invalid tokens[5]:", tk)
	}
}

func TestTokenizeNum(t *testing.T) {
	rd := createReader([]byte{})
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

	rd = createReader([]byte("2a"))
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

func (tk *token)checkNum(val int) bool {
    return tk.ty == TK_NUM && tk.val == val && string(tk.input) == strconv.Itoa(val)
}

func (tk *token)checkOp(op int) bool {
    return tk.ty == op && string(tk.input) == string(op)
}

func createReader(v []byte) *bufio.Reader {
	return bufio.NewReader(bytes.NewReader(v))
}
