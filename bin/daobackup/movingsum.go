package main

type MovingSum struct {
	buf [8194]byte
	i   uint32
	sum uint32
}

func (s *MovingSum) Add(b byte) {
	s.i = (s.i + 1) % 8194
	s.sum -= uint32(s.buf[s.i])
	s.sum += uint32(b)
	s.buf[s.i] = b
}

func (s *MovingSum) OnSplit(n uint) bool {
	return s.sum == s.sum & ^((uint32(1)<<n)-1)
}
