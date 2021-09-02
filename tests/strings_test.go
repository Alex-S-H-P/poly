package poly

import (
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	var P, P0 *Poly
	P = &Poly{Coefficients: []float64{1., 2., 3., 4., 5., 6.}}
	errors := 0
	if len(strings.Split(P.String(), "+")) != len(P.Coefficients) {
		P.Disp()
		errors++
		t.Error("P does not have a valid amount of tokens [",
			len(strings.Split(P.String(), "+")), ", ",
			len(P.Coefficients), "]")
	}
	if len(strings.Split(P0.String(), "+")) != 1 {
		P.Disp()
		errors++
		t.Error("null polynomial has an invalid amount of tokens")
	}
	PrintGreen("Successfull conversion to Strings")
}
