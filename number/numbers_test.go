package number

import "testing"

func TestPrimeFactors(t *testing.T) {
	// cases is a list of struct
	var testNumber int = 333
	want := []int{3, 37}
	got := GetPrimeFactor(testNumber)
	if got != want {
		t.Errorf("Wrong list generated, got: %v, want %v", got, want)
	}
}
