package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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

type tokens struct {
	tks []token
	i   int
}

func tokenize(rd *bufio.Reader) (tks *tokens) {
	var c byte
	var err error
	var tk token
	tks = &tokens{}

	space := byte(' ')
	lf := byte('\n')
	plus := byte('+')
	minus := byte('-')
	zero := byte('0')
	nine := byte('9')

	for c, err = rd.ReadByte(); err == nil; {
		if c == 0 || c == space || c == lf {
			c, err = rd.ReadByte()
			continue
		}

		if c == plus || c == minus {
			tk := token{
				ty:    int(c),
				input: []byte{c},
			}
			tks.tks = append(tks.tks, tk)

			c, err = rd.ReadByte()
			continue
		}

		if zero <= c && c <= nine {
			tk, c, err = tokenizeNum(rd, c)
			tks.tks = append(tks.tks, tk)
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
	tks.tks = append(tks.tks, tk)

	return
}

func tokenizeNum(rd *bufio.Reader, v byte) (tk token, c byte, err error) {
	zero := byte('0')
	nine := byte('9')

	var num []byte
	for c = v; err == nil; c, err = rd.ReadByte() {
		if c < zero || nine < c {
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

type node struct {
	ty, val  int
	lhs, rhs *node
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

func add(tks *tokens) (nd *node) {
	var tk token

	if tk = tks.current(); tk.ty != TK_NUM {
		log.Fatal("数値ではないトークンです: ", string(tk.input))
	}
	nd = &node{
		ty:  tk.ty,
		val: tk.val,
	}

	for tk = tks.next(); ; {
		switch {
		case tks.consume('+'):
			nd = &node{
				ty:  '+',
				lhs: nd,
				rhs: add(tks),
			}
		case tks.consume('-'):
			nd = &node{
				ty:  '-',
				lhs: nd,
				rhs: add(tks),
			}
		default:
			return
		}
	}
}

func main() {
	rd := bufio.NewReader(os.Stdin)
	tks := tokenize(rd)
	fmt.Println(tks)

	nd := add(tks)
	fmt.Println(nd)
}
