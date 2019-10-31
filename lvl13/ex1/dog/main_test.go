package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	testCases := []struct {
		desc   string
		input  int
		answer int
	}{
		{
			desc:   "Smoke",
			input:  1,
			answer: 7,
		},
		{
			desc:   "Positive",
			input:  3,
			answer: 21,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := Years(tC.input)
			if result != tC.answer {
				t.Error("Expected", tC.answer, "Got", result)
			}
		})
	}
}

func ExampleYears() {
	fmt.Println(Years(1))
	// Output:
	// 7
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(1)
	}
}

func TestYearsTwo(t *testing.T) {
	testCases := []struct {
		desc   string
		input  int
		answer int
	}{
		{
			desc:   "Smoke",
			input:  1,
			answer: 7,
		},
		{
			desc:   "Positive",
			input:  3,
			answer: 21,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			result := YearsTwo(tC.input)
			if result != tC.answer {
				t.Error("Expected", tC.answer, "Got", result)
			}
		})
	}
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(1))
	// Output:
	// 7
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(1)
	}
}
