// run some benchmark tests to compare popcount using loop, single expression and rightmost shift
package pocount

import (
	"testing"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCountTable(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountTableLoop(x uint64) int {
	sum := 0
	for i := 0; i < 8; i++ {
		sum += int(pc[byte(x>>uint(i))])
	}
	return sum
}

func PopCountShift(x uint64) int {
  count := 0
  mask := uint64(1)
  for i:=0; i<64; i++ {
    if x&mask > 0 {
      count++
    }
    x >>= 1
  }
  return count
}

func bench(b *testing.B, f func(uint64) int) {
	for i := 0; i < b.N; i++ {
		f(uint64(i))
	}
}

func BenchmarkTable(b *testing.B) {
	bench(b, PopCountTable)
}

func BenchmarkTableLoop(b *testing.B) {
	bench(b, PopCountTableLoop)
}

func BenchmarkShift(b *testing.B) {
	bench(b, PopCountShift)
}

/*
RESULTS:
goos: darwin
goarch: arm64
pkg: gopl/chapter_02/exercise_2_04
BenchmarkTable-8        1000000000               0.3185 ns/op
BenchmarkTableLoop-8    318265418                3.763 ns/op
BenchmarkShift-8        55860936                21.31 ns/op
PASS
ok      gopl/chapter_02/exercise_2_04   3.405s
*/
