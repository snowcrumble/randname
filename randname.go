package randname

import (
	"math"
	"fmt"
)

var defaultSequence int32 = math.MinInt32

func pseudoEncrypt(seed int32) int32 {
	var (
		l1 = (seed >> 16) & 0xffff
		r1 = seed & 0xffff
		l2, r2 int32
	)

	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ int32(math.Round((float64((1366*r1 + 150889) % 714025) / 714025.0) * 32767.0))
		l1 = l2
		r1 = r2
	}
	return (r1 << 16) + l1
}

type Generator struct{
	seq int32
}

func (x *Generator) SetSeq(seq int32) {
	x.seq = seq
}

func (x *Generator) Next() string {
	n := pseudoEncrypt(x.seq)
	x.seq++
	return fmt.Sprintf("%d", uint32(n))
}

func New() *Generator{
	return &Generator{seq: defaultSequence}
}
