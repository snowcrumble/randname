package randname

import (
	"fmt"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
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
	var m = make(map[string]struct{})

	maxSeq := int(g.MaxSeq())
	if maxSeq > 0xffff { //too big
		maxSeq = 0xffff
	}
	for i := 0; i <= maxSeq; i++ {
		name, _ := g.Next()
		if len(name) != length {
			t.Fatalf("invalid name length: %v, current i: %v\n", name, i)
		}
		if _, prs := m[name]; prs {
			t.Fatalf("repeated name: %v, current i: %v\n", name, i)
		}
		m[name] = struct{}{}
	}
}

func TestGeneratorConvey(t *testing.T) {
	length := 10
	var m = make(map[string]struct{})

	Convey("New a generator", t, func() {
		g := New(length)
		maxSeq := int(g.MaxSeq())
		if maxSeq > 0xffff { //too big
			maxSeq = 0xffff
		}

		Convey("Check generated IDs", func() {
			for i := 0; i <= maxSeq; i++ {
				name, _ := g.Next()
				if len(name) != length {
					t.Fatalf("invalid name length: %v, current i: %v\n", name, i)
				}
				if _, prs := m[name]; prs {
					t.Fatalf("repeated name: %v, current i: %v\n", name, i)
				}
				m[name] = struct{}{}
			}
		})
	})

}

func Test_getBits(t *testing.T) {
	if getBits(100) != 3 {
		t.Fatalf("getBits wrong")
	}
}

func ExampleGenerator() {
	g := New(8)

	for i := 0; i < 3; i++ {
		name, _ := g.Next()
		fmt.Printf("next name: %s\n", name)
	}
}
