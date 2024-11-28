package benchmark

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/PlayerR9/go-verify/OLD/common"
)

type Benchmark struct {
	test_count uint
	base_value uint
	fn_names   []string
	benchmarks [][]Instance
}

func NewBenchmark(fn_names ...string) *Benchmark {
	if len(fn_names) == 0 {
		return nil
	}

	return &Benchmark{
		test_count: 4,
		base_value: 10,
		fn_names:   fn_names,
	}
}

func (bm *Benchmark) ChangeTestCount(count uint) error {
	if bm == nil {
		return common.ErrNilReceiver
	}

	bm.test_count = count

	return nil
}

func (bm *Benchmark) ChangeBaseValue(base uint) error {
	if base == 0 {
		return common.ErrInvalidBase
	} else if bm == nil {
		return common.ErrNilReceiver
	}

	bm.base_value = base

	return nil
}

type BenchmarkFn func(exp uint)

type Instance struct {
	arg      string
	function BenchmarkFn
}

func (bm *Benchmark) AddBenchmark(arg string, fns ...BenchmarkFn) error {
	if bm == nil {
		return common.ErrNilReceiver
	}

	if len(fns) != len(bm.fn_names) {
		return fmt.Errorf("expected %d functions, got %d", len(bm.fn_names), len(fns))
	}

	for i, fn := range fns {
		if fn == nil {
			return fmt.Errorf("function at index %d is nil", i)
		}
	}

	instances := make([]Instance, 0, len(fns))

	if arg == "" {
		for i, fn := range fns {
			instances = append(instances, Instance{
				arg:      "/" + bm.fn_names[i],
				function: fn,
			})
		}
	} else {
		for i, fn := range fns {
			instances = append(instances, Instance{
				arg:      "/" + arg + "/" + bm.fn_names[i],
				function: fn,
			})
		}
	}

	bm.benchmarks = append(bm.benchmarks, instances)

	return nil
}

func (bm Benchmark) Run(b *testing.B) int {
	exponents, _ := common.UintPowSlice(bm.base_value, bm.test_count)

	var count int

	for _, instances := range bm.benchmarks {
		for _, n := range exponents {
			for _, instance := range instances {
				ok := b.Run(fmt.Sprintf("n=%d%s", n, instance.arg), func(b *testing.B) {
					for j := 0; j < b.N; j++ {
						instance.function(n)
					}
				})
				if ok {
					count++
				}
			}
		}
	}

	return count
}

func BenchmarkStringConcat(b *testing.B) {
	bm := NewBenchmark("+", "string_builder")

	sampleStrings := []string{
		"hello",
		"hellohello",
		"hellohellohello",
	}

	for _, str := range sampleStrings {
		_ = bm.AddBenchmark(
			"str="+strconv.Quote(str),
			func(exp uint) {
				var result string

				for j := uint(0); j < exp; j++ {
					result += str
				}
			},
			func(exp uint) {
				var builder strings.Builder

				for j := uint(0); j < exp; j++ {
					builder.WriteString(str)
				}
			},
		)
	}

	_ = bm.Run(b)
}
