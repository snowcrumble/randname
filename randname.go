package randname

import (
	"math"
	"fmt"
)

var (
	defaultSequence uint32 = 0
	defaultLength int = 10
)

func pseudoEncrypt(seed uint32, bits uint) uint32 {
	var (
		mask uint32 = (1<<bits) -1
		l1 = (seed >> bits) & mask
		r1 = seed & mask
		l2, r2 uint32
	)

	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ (uint32(math.Round((float64((1366*r1 + 150889) % 714025) / 714025.0) * 32767.0)) & mask)
		l1 = l2
		r1 = r2
	}
	return (r1 << bits) + l1
}

type Generator struct{
	seq uint32
	length int
	bits uint
}

func (x *Generator) SetSeq(seq uint32) {
	x.seq = seq
}

func (x *Generator) Next() string {
	n := pseudoEncrypt(x.seq, x.bits)
	x.seq++

	format := fmt.Sprintf("%%.%dd", x.length) //字符串前补0格式
	return fmt.Sprintf(format, uint32(n))
}

//New 创建一个唯一ID生成器，可指定数字最大长度，但需要注意，支持的sequence数值范围为 [0, MaxSeq()]
func New(length int) *Generator{
	if length > 10 || length < 1{ //限制最大最小长度
		length = defaultLength
	}
	
	bits := getBits(int(math.Pow10(length)-1))
	// fmt.Println("math.Pow10(length)-1: ", math.Pow10(length)-1)
	// fmt.Println("new bits: ", bits)
	
	return &Generator{
		seq: defaultSequence,
		length: length,
		bits: bits,
	}
}

//MaxSeq 返回可用的最大sequence值
func (x *Generator) MaxSeq() uint32 {
	return (1 << (x.bits*2)) -1
}

//getBits 返回 (一个整数占用多少二进制位) / 2
func getBits(n int) (bits uint) {
	var x int
	var i uint
	for ; i < 64; i++{
		y := x | (1 << i)
		if y > n {
			break
		}
		x = y
	}
	bits = i/2
	return
}
