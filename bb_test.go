package main

import (
	"testing"

	"github.com/miltfra/tools/ds/graph"
)

func BenchmarkTSPBB2887(b *testing.B) {
	b.StopTimer()
	g := graph.FromFile("/home/miltfra/projects/example_files/graphs/wild20.txt", 1)
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		TSPBB(g.Edges, 8, 1<<28, 7)
		b.StopTimer()
	}
}
func BenchmarkTSPBB2888(b *testing.B) {
	b.StopTimer()
	g := graph.FromFile("/home/miltfra/projects/example_files/graphs/wild20.txt", 1)
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		TSPBB(g.Edges, 8, 1<<28, 8)
		b.StopTimer()
	}
}

func BenchmarkTSPBB28810(b *testing.B) {
	b.StopTimer()
	g := graph.FromFile("/home/miltfra/projects/example_files/graphs/wild20.txt", 1)
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		TSPBB(g.Edges, 8, 1<<28, 10)
		b.StopTimer()
	}
}

func BenchmarkTSPBB28815(b *testing.B) {
	b.StopTimer()
	g := graph.FromFile("/home/miltfra/projects/example_files/graphs/wild20.txt", 1)
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		TSPBB(g.Edges, 8, 1<<28, 15)
		b.StopTimer()
	}
}
func BenchmarkTSPBB2787(b *testing.B) {
	b.StopTimer()
	g := graph.FromFile("/home/miltfra/projects/example_files/graphs/wild20.txt", 1)
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		TSPBB(g.Edges, 8, 1<<27, 7)
		b.StopTimer()
	}
}

func BenchmarkTSPBB2687(b *testing.B) {
	b.StopTimer()
	g := graph.FromFile("/home/miltfra/projects/example_files/graphs/wild20.txt", 1)
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		TSPBB(g.Edges, 8, 1<<26, 7)
		b.StopTimer()
	}
}
func BenchmarkTSPBB2987(b *testing.B) {
	b.StopTimer()
	g := graph.FromFile("/home/miltfra/projects/example_files/graphs/wild20.txt", 1)
	for n := 0; n < b.N; n++ {
		b.StartTimer()
		TSPBB(g.Edges, 8, 1<<29, 7)
		b.StopTimer()
	}
}
