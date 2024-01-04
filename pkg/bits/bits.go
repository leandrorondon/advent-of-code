package bits

import (
	"fmt"
	mathbits "math/bits"
)

type Bits struct {
	bits []uint64
	size int
}

func New(size int) *Bits {
	if size <= 0 {
		panic("invalid size")
	}

	var bits Bits
	bits.size = size

	if size%64 > 0 {
		bits.bits = make([]uint64, 1+size/64)
		return &bits
	}

	bits.bits = make([]uint64, size/64)

	return &bits
}

func NewFromString(s string) *Bits {
	bits := New(len(s))

	for i := range s {
		if s[i] == '1' {
			bits.Set(len(s) - i - 1)
		} else if s[i] != '0' {
			panic("unexpected binary string")
		}
	}

	return bits
}

func (b *Bits) Size() int {
	return b.size
}

func (b *Bits) Set(pos int) {
	idx, bit := b.bitpos(pos)
	b.bits[idx] = b.bits[idx] | (1 << bit)
}

func (b *Bits) Clear(pos int) {
	idx, bit := b.bitpos(pos)
	b.bits[idx] = b.bits[idx] & ^(1 << bit)
}

func (b *Bits) Reset() {
	for i := range b.bits {
		b.bits[i] = 0
	}
}

func (b *Bits) Get(pos int) int {
	idx, bit := b.bitpos(pos)
	return int((b.bits[idx] & (1 << bit)) >> bit)
}

func (b *Bits) Toggle(pos int) int {
	idx, bit := b.bitpos(pos)
	fmt.Println(idx, bit, b.bits[idx], b.bits[idx]^(1<<bit))
	b.bits[idx] = b.bits[idx] ^ (1 << bit)

	return int((b.bits[idx] & (1 << bit)) >> bit)
}

func (b *Bits) OnesCount() int {
	sum := 0
	for i := range b.bits {
		sum += mathbits.OnesCount64(b.bits[i])
	}
	return sum
}

func (b *Bits) String() string {
	s := ""
	for i := 0; i < b.size; i++ {
		s = string(rune(b.Get(i)+'0')) + s
	}
	return s
}

func (b *Bits) bitpos(pos int) (int, int) {
	if pos < 0 || pos >= b.size {
		panic("out of range")
	}

	return pos / 64, pos % 64
}
