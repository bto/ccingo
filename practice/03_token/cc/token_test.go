package cc

import (
	"bufio"
	"bytes"
	"io"
	"testing"
)

func TestTokenize(t *testing.T) {
	rd := createReader([]byte(" 1+  23\n -456 \n"))
	tks := Tokenize(rd)
	if len(tks) != 6 {
		t.Fatal("invalid number of tokens:", len(tks))
	}
	if tk := tks[0]; tk.ty != TK_NUM || tk.val != 1 || string(tk.input) != "1" {
		t.Fatal("invalid tokens[0]:", tk)
	}
	if tk := tks[1]; tk.ty != '+' || string(tk.input) != "+" {
		t.Fatal("invalid tokens[1]:", tk)
	}
	if tk := tks[2]; tk.ty != TK_NUM || tk.val != 23 || string(tk.input) != "23" {
		t.Fatal("invalid tokens[2]:", tk)
	}
	if tk := tks[3]; tk.ty != '-' || string(tk.input) != "-" {
		t.Fatal("invalid tokens[3]:", tk)
	}
	if tk := tks[4]; tk.ty != TK_NUM || tk.val != 456 || string(tk.input) != "456" {
		t.Fatal("invalid tokens[4]:", tk)
	}
	if tk := tks[5]; tk.ty != TK_EOF {
		t.Fatal("invalid tokens[5]:", tk)
	}
}

func TestTokenizeNum(t *testing.T) {
	rd := createReader([]byte{})
	tk, c, err := tokenizeNum(rd, '1')
	if tk.ty != TK_NUM || tk.val != 1 || string(tk.input) != "1" {
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
	if tk.ty != TK_NUM || tk.val != 12 || string(tk.input) != "12" {
		t.Fatal("invalid token:", tk)
	}
	if c != byte('a') {
		t.Fatal("invalid next charactor:", c)
	}
	if err != nil {
		t.Fatal("invalid error:", err)
	}
}

func createReader(v []byte) *bufio.Reader {
	return bufio.NewReader(bytes.NewReader(v))
}
