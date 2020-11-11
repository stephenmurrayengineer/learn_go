package iteration

import "testing"
import "fmt"

func ExampleRepeat() {
	repeated := Repeat("s", 6)
	fmt.Println(repeated)
	// Output: ssssss
}

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 5)
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 1)
	}
}

func TestMakeTitle(t *testing.T) {
	titled :=  MakeTitle("  the,great,gatsby  ")
	expected := "The Great Gatsby"

	if titled != expected {
		t.Errorf("expected %q, got %q", expected,titled)
	}
}

