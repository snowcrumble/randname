//Package randname 一个基于顺序整数的随机唯一ID生成器
package randname

import (
	"fmt"
	"math"
)

var (
	defaultSequence uint32
	defaultLength   = 10
)

func pseudoEncrypt(seed uint32, bits uint) uint32 {
	var (
		mask   uint32 = (1 << bits) - 1
		l1            = (seed >> bits) & mask
		r1            = seed & mask
		l2, r2 uint32
	)

	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ (uint32(math.Round((float64((1366*r1+150889)%714025)/714025.0)*32767.0)) & mask)
		l1 = l2
		r1 = r2
	}
	return (r1 << bits) + l1
}

//Acc Accumulator
type Acc interface {
	//Incr 返回+1后的结果
	Incr() (uint32, error)
}

//SimpleAcc 简单计数器
type SimpleAcc struct {
	Seq uint32
}

//Incr 返回+1后的结果
func (x *SimpleAcc) Incr() (uint32, error) {
	x.Seq++
	return x.Seq, nil
}

//Generator 唯一ID生成器
type Generator struct {
	acc    Acc
	length int
	bits   uint
}

//SetAcc 设置序列累加器
func (x *Generator) SetAcc(acc Acc) {
	x.acc = acc
}

//Next 获取一个新的ID
func (x *Generator) Next() (string, error) {
	seq, err := x.acc.Incr()
	if err != nil {
		return "", err
	}
	n := pseudoEncrypt(seq, x.bits)

	format := fmt.Sprintf("%%.%dd", x.length) //字符串前补0格式
	return fmt.Sprintf(format, uint32(n)), nil
}

//New 创建一个唯一ID生成器，可指定数字最大长度，但需要注意，支持的sequence数值范围为 [0, MaxSeq()]
func New(length int) *Generator {
	if length > 10 || length < 1 { //限制最大最小长度
		length = defaultLength
	}

	bits := getBits(int(math.Pow10(length) - 1))

	return &Generator{
		acc:    &SimpleAcc{},
		length: length,
		bits:   bits,
	}
}

//MaxSeq 返回可用的最大sequence值
func (x *Generator) MaxSeq() uint32 {
	return (1 << (x.bits * 2)) - 1
}

//getBits 返回 (小于n的最大的二进制位全为1的整数所占用的位) / 2
// n 为偶数时损失的ID数量会少一些
func getBits(n int) (bits uint) {
	var x int
	var i uint
	for ; i < 64; i++ {
		y := x | (1 << i) //将第i位设为1
		if y > n {
			break
		}
		x = y
	}
	bits = i / 2
	return
}
