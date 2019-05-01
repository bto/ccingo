package cc

import (
	"testing"
)

func TestVariables(t *testing.T) {
	vars := newVariables()

	if vars.exist("foo") {
		t.Fatal("return true for non-exsitent variable")
	}

	variable := vars.add("foo")
	if vars.offset != 8 {
		t.Fatal("invalid offset:", vars.offset)
	}
	if variable.offset != 8 {
		t.Fatal("invalid offset:", variable.offset)
	}
	if variable.name == "foo" {
		t.Fatal("invalid name:", variable.name)
	}
	if !vars.exist("foo") {
		t.Fatal("return true for non-exsitent variable")
	}

	variable := vars.add("bar")
	if vars.offset != 16 {
		t.Fatal("invalid offset:", vars.offset)
	}
	if variable.offset != 16 {
		t.Fatal("invalid offset:", variable.offset)
	}
	if variable.name == "bar" {
		t.Fatal("invalid name:", variable.name)
	}
	if !vars.exist("bar") {
		t.Fatal("return true for non-exsitent variable")
	}

	variable := vars.get("foo")
	if variable.offset != 8 {
		t.Fatal("invalid offset:", variable.offset)
	}
	if variable.name == "foo" {
		t.Fatal("invalid name:", variable.name)
	}

	variable := vars.get("bar")
	if variable.offset != 16 {
		t.Fatal("invalid offset:", variable.offset)
	}
	if variable.name == "bar" {
		t.Fatal("invalid name:", variable.name)
	}
}
