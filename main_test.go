package main

import (
	"bufio"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestDays(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name   string
		solver Solver
		input  *bufio.Scanner
		want   int
	}{
		{
			name:   "Day 01 Part 1",
			solver: Day01Part1,
			input:  ExampleScanner(1),
			want:   11,
		}, {
			name:   "Day 01 Part 2",
			solver: Day01Part2,
			input:  ExampleScanner(1),
			want:   31,
		}, {
			name:   "Day 02 Part 1",
			solver: Day02Part1,
			input:  ExampleScanner(2),
			want:   2,
		}, {
			name:   "Day 02 Part 2",
			solver: Day02Part2,
			input:  ExampleScanner(2),
			want:   4,
		}, {
			name:   "Day 03 Part 1",
			solver: Day03Part1,
			input:  ExampleScanner(3),
			want:   161,
		}, {
			name:   "Day 03 Part 2",
			solver: Day03Part2,
			input:  ExampleScanner(3),
			want:   48,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.solver(tc.input)
			if diff := cmp.Diff(got, tc.want); diff != "" {
				t.Errorf("mismatch (-got, +want):\n%v", diff)
			}
		})
	}
}

// BenchmarkDay01Part1
// BenchmarkDay01Part1-16              1692            917222 ns/op          186044 B/op       4038 allocs/op
// BenchmarkDay01Part1-16              1242            917443 ns/op          186045 B/op       4038 allocs/op
// BenchmarkDay01Part1-16              1282            934330 ns/op          186046 B/op       4038 allocs/op
// BenchmarkDay01Part1-16              1243            860598 ns/op          186045 B/op       4038 allocs/op
func BenchmarkDay01Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day01Part1(DayScanner(1))
	}
}

// BenchmarkDay01Part2
// BenchmarkDay01Part2-16              1484            901969 ns/op          204101 B/op       4069 allocs/op
// BenchmarkDay01Part2-16              1322            940059 ns/op          204087 B/op       4069 allocs/op
// BenchmarkDay01Part2-16              1252            936612 ns/op          204083 B/op       4069 allocs/op
// BenchmarkDay01Part2-16              1195            930953 ns/op          204089 B/op       4069 allocs/op
func BenchmarkDay01Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day01Part2(DayScanner(1))
	}
}

// BenchmarkDay02Part1
// BenchmarkDay02Part1-16               646           1920380 ns/op          444965 B/op       9726 allocs/op
// BenchmarkDay02Part1-16               672           1970817 ns/op          444970 B/op       9726 allocs/op
// BenchmarkDay02Part1-16               600           1891085 ns/op          444965 B/op       9726 allocs/op
// BenchmarkDay02Part1-16               618           1874649 ns/op          444967 B/op       9726 allocs/op
func BenchmarkDay02Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day02Part1(DayScanner(2))
	}
}

// BenchmarkDay02Part2
// BenchmarkDay02Part2-16               277           4576999 ns/op          946365 B/op      22515 allocs/op
// BenchmarkDay02Part2-16               254           4559365 ns/op          946360 B/op      22515 allocs/op
// BenchmarkDay02Part2-16               262           4276266 ns/op          946359 B/op      22515 allocs/op
// BenchmarkDay02Part2-16               264           4531602 ns/op          946373 B/op      22516 allocs/op
func BenchmarkDay02Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day02Part2(DayScanner(2))
	}
}

// BenchmarkDay03Part1
// BenchmarkDay03Part1-16              1458            893275 ns/op          165204 B/op       2137 allocs/op
// BenchmarkDay03Part1-16              1190            930700 ns/op          165172 B/op       2137 allocs/op
// BenchmarkDay03Part1-16              1204            867937 ns/op          165378 B/op       2137 allocs/op
// BenchmarkDay03Part1-16              1490            899983 ns/op          165133 B/op       2137 allocs/op
func BenchmarkDay03Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day03Part1(DayScanner(3))
	}
}

// BenchmarkDay03Part2
// BenchmarkDay03Part2-16               591           1714315 ns/op          182814 B/op       1876 allocs/op
// BenchmarkDay03Part2-16               610           2047430 ns/op          182917 B/op       1876 allocs/op
// BenchmarkDay03Part2-16               589           1783785 ns/op          183069 B/op       1876 allocs/op
// BenchmarkDay03Part2-16               574           1785652 ns/op          183218 B/op       1876 allocs/op
func BenchmarkDay03Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Day03Part2(DayScanner(3))
	}
}
