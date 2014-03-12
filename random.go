package random

import (
	crand "crypto/rand"
	"io"
	"math/big"
	"math/rand"
	"sort"
	"sync"
)

type asciiflag uint8

const (
	UPPER  asciiflag = 1 << iota
	LOWER  asciiflag = 1 << iota
	DIGITS asciiflag = 1 << iota
)

func Ascii(flag asciiflag) (r io.RuneReader, err error) {
	chain, err := newChain()
	if err != nil {
		return
	}

	if flag&DIGITS > 0 {
		r, err := newReader(10, 48)
		if err != nil {
			return nil, err
		}

		chain.add(r)
	}

	if flag&UPPER > 0 {
		r, err := newReader(26, 65)
		if err != nil {
			return nil, err
		}

		chain.add(r)
	}

	if flag&LOWER > 0 {
		r, err := newReader(26, 97)
		if err != nil {
			return nil, err
		}

		chain.add(r)
	}

	return chain, nil
}

type chain struct {
	readers []*reader
	max     sort.IntSlice
	maxsum  int32
	rand    *rand.Rand
	sync.Mutex
}

func newChain() (c *chain, err error) {
	i, err := crand.Int(crand.Reader, big.NewInt(1<<63-1))
	if err != nil {
		return
	}

	c = &chain{
		readers: make([]*reader, 0, 3),
		max:     make(sort.IntSlice, 0, 3),
		maxsum:  0,
		rand:    rand.New(rand.NewSource(i.Int64())),
	}
	return
}

func (c *chain) add(r *reader) {
	c.Lock()
	defer c.Unlock()

	c.readers = append(c.readers, r)
	c.maxsum += r.max
	c.max = append(c.max, int(c.maxsum))
}

func (c *chain) ReadRune() (rune, int, error) {
	i := c.rand.Int31n(c.maxsum)
	return c.readers[c.max.Search(int(i))].ReadRune()
}

type reader struct {
	rand   *rand.Rand
	max    int32
	prefix int32
}

func newReader(max, prefix int32) (r *reader, err error) {
	i, err := crand.Int(crand.Reader, big.NewInt(1<<63-1))
	if err != nil {
		return
	}

	r = &reader{
		rand:   rand.New(rand.NewSource(i.Int64())),
		max:    max,
		prefix: prefix,
	}
	return
}

func (r *reader) randint32() int32 {
	return r.rand.Int31n(r.max)
}

func (r *reader) ReadRune() (rune, int, error) {
	return rune(r.randint32() + r.prefix), 1, nil
}
