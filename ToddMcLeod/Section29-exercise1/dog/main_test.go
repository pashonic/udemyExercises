package dog

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	got := Years(7)
	if got != 49 {
		t.Errorf("Got %v, expected 70", got)
	}
}

func TestYearsTwo(t *testing.T) {
	got := YearsTwo(7)
	if got != 49 {
		t.Errorf("Got %v, expected 70", got)
	}
}

func ExampleYears() {
	fmt.Println(Years(10))
	// Output:
	// 70
}

func ExampleYearsTwo() {
	fmt.Println(YearsTwo(10))
	// Output:
	// 70
}

func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(10)
	}
}

func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(10)
	}
}
