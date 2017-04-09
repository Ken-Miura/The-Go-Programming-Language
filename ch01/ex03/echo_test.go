// Copyright 2017 Ken Mirua
package ex03

import "testing"

func BenchmarkEchoInefficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoInefficient()
	}
}

func BenchmarkEchoEfficient(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoEfficient()
	}
}
