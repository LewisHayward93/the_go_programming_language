// run some benchmark tests to compare concatenation vs string.Join
package exercise103

import (
	"strings"
	"testing"
)

var args = []string{"hi", "there", "buddy", "boy", "5", "6", "7", "8", "9"}

func concat(args []string) {
	r, sep := "", ""
	for _, a := range args {
		r += sep + a
		sep = " "
	}
}

func BenchmarkConcat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concat(args)
	}
}

func BenchmarkJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Join(args, " ")
	}
}

/*
RESULTS: ran on M1 Macbook Air 8GB

goos: darwin
goarch: arm64
pkg: gopl/chapter_01/exercise_1_03
BenchmarkConcat-8        4607358               227.7 ns/op
BenchmarkJoin-8         17120569                69.76 ns/op
PASS
ok      gopl/chapter_01/exercise_1_03   3.033s
*/
