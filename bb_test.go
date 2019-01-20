package main

import (
	"testing"

	"github.com/miltfra/tools"
)

func BenchmarkTSPBB10(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		g := tools.RndDropoutGraph(10, 0.5)
		b.StartTimer()
		TSPBB(g)
		b.StopTimer()
	}

}
func BenchmarkTSPBB20(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		g := tools.RndDropoutGraph(20, 0.5)
		b.StartTimer()
		TSPBB(g)
		b.StopTimer()
	}
}

func BenchmarkTSPBB30(b *testing.B) {
	b.StopTimer()
	for n := 0; n < b.N; n++ {
		g := tools.RndDropoutGraph(30, 0.5)
		b.StartTimer()
		TSPBB(g)
		b.StopTimer()
	}
}
func TestTSPBB20(t *testing.T) {
	n := 20
	g := tools.RndGraph(n)
	cost, path := TSPBB(g)
	actualCost := float64(0)
	last := path[len(path)-1]
	for i := 0; i < len(path); i++ {
		actualCost += g[last][path[i]]
		last = path[i]
	}
	if actualCost != cost {
		t.FailNow()
	}
}
