package halma

type Set struct {
	data [16]byte
}

func (s *Set) clear() {
	for n:=0; n<16; n++ {
		s.data[n] = 0
	}
}

func (s *Set) include(n int) {
	s.data[(n)>>3] |= 1 << ((n-1) & 7)
}

func (s *Set) in(n int) bool {
	return (s.data[(n)>>3] & ( 1 << ((n-1) & 7))) != 0
}