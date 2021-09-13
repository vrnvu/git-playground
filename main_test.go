package main

import "testing"

func TestFoo(t *testing.T) {
	if Foo() != "foo" {
		t.Error("foo")
	}
}

func TestBar(t *testing.T) {
	if Bar() != "bar" {
		t.Error("bar")
	}
}

func TestAyy(t *testing.T) {
	if Ayy() != "lmao" {
		t.Error("lmao")
	}
}

func TestKek(t *testing.T) {
	if Kek() != "kek" {
		t.Error("kek")
	}
}
