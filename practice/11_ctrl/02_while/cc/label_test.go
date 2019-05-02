package cc

import (
	"testing"
)

func TestGet(t *testing.T) {
	lb := newLabel()
	if s := lb.get("if"); s != ".if1" {
		t.Fatal("invalid label:", s)
	}
	if s := lb.get("begin"); s != ".begin2" {
		t.Fatal("invalid label:", s)
	}
	if s := lb.get("end"); s != ".end3" {
		t.Fatal("invalid label:", s)
	}
}
