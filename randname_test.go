package randname

import (
	"testing"
	"fmt"
)

func Test_pseudoEncrypt(t *testing.T) {
	var seed int32 = -0xffff
	var m = make(map[int32]struct{})
	for ; seed <= 0xffff; seed++ {
		r := pseudoEncrypt(seed)
		if _, prs := m[r]; prs {
			t.Fatalf("repeated number: input = %v, output = %v\n", seed, r)
		}
		m[r] = struct{}{}
	}
}

func TestGenerator(t *testing.T) {
	g := New()
	g.SetSeq(-0xffff)
	
	var m = make(map[string]struct{})
	for i:=0; i <= 0xffff; i++ {
		name := g.Next()
		if _, prs := m[name]; prs {
			t.Fatalf("repeated name: %v\n", name)
		}
		m[name] = struct{}{}
	}
}

func ExampleGenerator() {
	g := New()
	g.SetSeq(-0xffff)

	for i := 0; i < 3; i++ {
		name := g.Next()
		fmt.Printf("next name: %s\n", name)
	}
}