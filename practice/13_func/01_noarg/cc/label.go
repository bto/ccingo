package cc

import (
	"strconv"
)

type label struct {
	n int
}

func (lb *label) get(key string) string {
	lb.n++
	return "." + key + strconv.Itoa(lb.n)
}

func newLabel() *label {
	return &label{}
}
