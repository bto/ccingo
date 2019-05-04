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

func (tks *tokens) append(tk token) *tokens {
	tks.tks = append(tks.tks, tk)
	return tks
}

func Tokenize(rd *bufio.Reader) (tks *tokens) {
	var c byte
	var err error
	var tk token
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
			tk := token{
				ty:    TK_IDENT,
				input: []byte{c},
			}
			tks.append(tk)

			c, err = rd.ReadByte()
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
