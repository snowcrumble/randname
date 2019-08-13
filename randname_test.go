package randname

import (
	"testing"
	"fmt"
)

func Test_pseudoEncrypt(t *testing.T) {
	var seed uint32
	var m = make(map[uint32]struct{})
	for ; seed <= 9999; seed++ {
		r := pseudoEncrypt(seed, 7)
		if _, prs := m[r]; prs {
			t.Fatalf("repeated number: input = %v, output = %v\n", seed, r)
		}
		m[r] = struct{}{}
	}
}

func TestGenerator(t *testing.T) {
	length := 10
	g := New(length)
	g.SetSeq(0)
	var m = make(map[string]struct{})
	
	maxSeq := int(g.MaxSeq())
	if maxSeq > 0xffff { //too big
		maxSeq = 0xffff
	}
	for i:=0; i <= maxSeq; i++ {
		name := g.Next()
		if len(name) != length {
			t.Fatalf("invalid name length: %v, current i: %v\n", name, i)
		}
		if _, prs := m[name]; prs {
			t.Fatalf("repeated name: %v, current i: %v\n", name, i)
		}
		m[name] = struct{}{}
	}
}

func Test_getBits(t *testing.T) {
	if getBits(100) != 3 {
		t.Fatalf("getBits wrong")
	}
}

func ExampleGenerator() {
	g := New(8)
	g.SetSeq(0)

	for i := 0; i < 3; i++ {
		name := g.Next()
		fmt.Printf("next name: %s\n", name)
	}
}