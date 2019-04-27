package cc

import (
	"bufio"
	"io"
	"log"
	"strconv"
)

const (
	TK_NUM = iota + 256
	TK_EOF
)

type token struct {
	ty, val int
	input   []byte
}

func Tokenize(rd *bufio.Reader) (tks []token) {
	var c byte
	var err error
	var tk token

	for c, err = rd.ReadByte(); err == nil; {
		switch c {
		case 0, byte(' '), byte('\n'):
			c, err = rd.ReadByte()
			continue
		case byte('+'), byte('-'):
			tk := token{
				ty:    int(c),
				input: []byte{c},
			}
			tks = append(tks, tk)

			c, err = rd.ReadByte()
			continue
		}

		if byte('0') <= c && c <= byte('9') {
			tk, c, err = tokenizeNum(rd, c)
			tks = append(tks, tk)
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
	tks = append(tks, tk)

	return
}

func tokenizeNum(rd *bufio.Reader, v byte) (tk token, c byte, err error) {
	var num []byte
	for c = v; err == nil; c, err = rd.ReadByte() {
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
