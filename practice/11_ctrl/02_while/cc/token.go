package cc

import (
	"bufio"
	"io"
	"log"
	"strconv"
)

const (
	TK_NUM = iota + 256
	TK_EQ
	TK_NE
	TK_LE
	TK_GE
	TK_IDENT
	TK_RETURN
	TK_IF
	TK_WHILE
	TK_EOF
)

type token struct {
	ty, val int
	input   []byte
}

type tokens struct {
	tks []token
	i   int
}

func (tks *tokens) append(tk token) {
	tks.tks = append(tks.tks, tk)
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

func Tokenize(rd *bufio.Reader) (tks *tokens) {
	var c byte
	var err error
	var tk token
	var name []byte
	tks = &tokens{}

	for c, err = rd.ReadByte(); err == nil; {
		switch c {
		case byte(' '), byte('\n'):
			c, err = rd.ReadByte()
			continue
		case byte('='):
			c, err = rd.ReadByte()
			if c == byte('=') {
				tk := token{
					ty:    TK_EQ,
					input: []byte{'=', c},
				}
				tks.append(tk)

				c, err = rd.ReadByte()
			} else {
				tk := token{
					ty:    int('='),
					input: []byte{'='},
				}
				tks.append(tk)
			}
			continue
		case byte('!'):
			c, err = rd.ReadByte()
			if c != byte('=') {
				log.Fatal("トークナイズできません: ", string([]byte{'!', c}))
			}

			tk := token{
				ty:    TK_NE,
				input: []byte{'!', c},
			}
			tks.append(tk)

			c, err = rd.ReadByte()
			continue
		case byte('<'):
			c, err = rd.ReadByte()
			if c == byte('=') {
				tk := token{
					ty:    TK_LE,
					input: []byte{'<', c},
				}
				tks.append(tk)

				c, err = rd.ReadByte()
			} else {
				tk := token{
					ty:    int('<'),
					input: []byte{'<'},
				}
				tks.append(tk)
			}
			continue
		case byte('>'):
			c, err = rd.ReadByte()
			if c == byte('=') {
				tk := token{
					ty:    TK_GE,
					input: []byte{'>', c},
				}
				tks.append(tk)

				c, err = rd.ReadByte()
			} else {
				tk := token{
					ty:    int('>'),
					input: []byte{'>'},
				}
				tks.append(tk)
			}
			continue
		case byte('+'), byte('-'), byte('*'), byte('/'), byte('('), byte(')'), byte(';'):
			tk := token{
				ty:    int(c),
				input: []byte{c},
			}
			tks.append(tk)

			c, err = rd.ReadByte()
			continue
		}

		if byte('0') <= c && c <= byte('9') {
			tk, c, err = tokenizeNum(rd, c)
			tks.append(tk)
			continue
		}

		if byte('a') <= c && c <= byte('z') {
			name, c, err = tokenizeAlNum(rd, c)
			tk := token{
				input: name,
			}

			switch string(name) {
			case "return":
				tk.ty = TK_RETURN
				tks.append(tk)
				continue
			case "if":
				tk.ty = TK_IF
				tks.append(tk)
				continue
			case "while":
				tk.ty = TK_WHILE
				tks.append(tk)
				continue
			}

			tk.ty = TK_IDENT
			tks.append(tk)
			continue
		}

		log.Fatal("トークナイズできません: ", string([]byte{c}))
	}
	if err != io.EOF {
		log.Fatal(err)
	}

	tk = token{
		ty: TK_EOF,
	}
	tks.append(tk)

	return
}

func tokenizeAlNum(rd *bufio.Reader, v byte) (name []byte, c byte, err error) {
	for c = v; err == nil; c, err = rd.ReadByte() {
		if (c < byte('a') || byte('z') < c) && (c < byte('0') || byte('9') < c) {
			break
		}

		name = append(name, c)
	}

	return
}

func tokenizeNum(rd *bufio.Reader, v byte) (tk token, c byte, rdErr error) {
	var num []byte
	for c = v; rdErr == nil; c, rdErr = rd.ReadByte() {
		if c < byte('0') || byte('9') < c {
			break
		}

		num = append(num, c)
	}

	val, err := strconv.Atoi(string(num))
	if err != nil {
		log.Fatal(err)
	}

	tk = token{
		ty:    TK_NUM,
		val:   val,
		input: num,
	}

	return
}
