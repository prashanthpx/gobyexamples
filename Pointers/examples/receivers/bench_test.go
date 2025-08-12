package receivers

import "testing"

type Medium struct {
	A, B, C, D, E, F, G, H, I, J int64
}

func (m Medium) SumV() int64 { return m.A + m.B + m.C + m.D + m.E + m.F + m.G + m.H + m.I + m.J }
func (m *Medium) SumP() int64 { return m.A + m.B + m.C + m.D + m.E + m.F + m.G + m.H + m.I + m.J }

func BenchmarkValueRecv(b *testing.B) {
	m := Medium{1,2,3,4,5,6,7,8,9,10}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = m.SumV()
	}
}

func BenchmarkPtrRecv(b *testing.B) {
	m := Medium{1,2,3,4,5,6,7,8,9,10}
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = (&m).SumP()
	}
}

