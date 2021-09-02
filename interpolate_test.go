package polynom

import (
	"fmt"
	"testing"
)

func TestInterpolate(t *testing.T) {
	Xs := []float64{1., 2., 3.}
	Ys := []float64{5., 6., 7.}
	P, err := Interpolate(Xs, Ys)
	errors := 0
	if err != nil {
		t.Error(err)
		errors++
		P.Disp()
	}
	if P.Degre() <= 0 {
		fmt.Printf("%d instead of %d\n", P.Degre(), len(Xs)-1)
		t.Error("Failed to make a polynomial.")
		errors++
	}
	for i := 0; i < len(Xs); i++ {
		if P.Call(Xs[i]) != Ys[i] {
			t.Error("Failed to find correct y for defined x :", P.Call(Xs[i]), "instead of", Ys[i])
			errors++
		}
	}
	Ys2 := []float64{1., 2.}
	P, err = Interpolate(Xs, Ys2)
	if P == nil || err == nil {
		t.Error("Failed to detect non-fatal_error [", P, ", ", err, "]")
		errors++
	}
	Ys3 := []float64{}
	P, err = Interpolate(Xs, Ys3)
	if P != nil || err == nil {
		t.Error("Failed to detect fatal error [", P, ", ", err, "]")
	}
	if errors == 0 {
		PrintGreen("Interpolation is valid")
	}
}
