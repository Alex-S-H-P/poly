package polynom

import (
	"fmt"
	"math"
	"testing"
)

func PrintGreen(format string, a ...interface{}) (int, error) {
	// fmt.Printn wrapper that also handles coloring text in green
	fmt.Printf("\u001B[32m")
	res1, res2 := fmt.Printf(format, a...)
	fmt.Println("\u001B[0m")
	return res1, res2
}

func TestDegre(t *testing.T) {
	errors := 0
	for i := 1; i < 10; i++ {
		X := Poly{Coefficients: make([]float64, i)}
		if X.Degre() != -1 {
			t.Fail()
			errors++
		}
		X = Poly{Coefficients: make([]float64, i)}
		X.Coefficients[i-1] = 1.
		if X.Degre() != i-1 {
			fmt.Println("Got invalid value instead of ", i, ":", X.Degre())
			t.Fail()
			errors++
		}
	}
	PrintGreen("Degre handles standard cases")
	Y := Poly{}
	if Y.Degre() != -1 {
		t.Fail()
		errors++
	}
	if errors != 0 {
		PrintGreen("Degre seems to work...")
	}
}

func TestCall(t *testing.T) {
	var EmptyNull, HiddenNull, NonEmptyNull, Standard Poly
	EmptyNull = Poly{}
	HiddenNull = Poly{Coefficients: []float64{}}
	NonEmptyNull = Poly{Coefficients: []float64{0., 0., 0.}}
	Standard = Poly{Coefficients: []float64{1., 2., 3.}}
	errors := 0
	if EmptyNull.Call(0.) != 0. || EmptyNull.Call(1.) != 0 {
		t.Fail()
		errors++
	}
	if NonEmptyNull.Call(0.) != 0. || NonEmptyNull.Call(1.) != 0 {
		t.Fail()
		errors++
	}
	if HiddenNull.Call(0.) != 0. || EmptyNull.Call(1.) != 0 {
		t.Fail()
		errors++
	}
	if Standard.Call(0.) != 1. || Standard.Call(1.) != 6 {
		t.Fail()
		errors++
	}
	if errors == 0 {
		PrintGreen("Polynomial works as intended")
	}
}

func TestMult(t *testing.T) {
	var EmptyNull, HiddenNull, NonEmptyNull, Standard Poly
	EmptyNull = Poly{}
	HiddenNull = Poly{Coefficients: []float64{}}
	NonEmptyNull = Poly{Coefficients: []float64{0., 0., 0.}}
	Standard = Poly{Coefficients: []float64{1., 2., 3.}}
	errors := 0
	if (EmptyNull.Mult(&Standard)).Degre() != -1 {
		t.Fail()
		errors++
	}
	if (HiddenNull.Mult(&Standard)).Degre() != -1 {
		t.Fail()
		errors++
	}
	if (NonEmptyNull.Mult(&Standard)).Degre() != -1 {
		t.Fail()
		errors++
	}
	if Standard.Mult(&Standard).Degre() != 2*Standard.Degre() {
		t.Fail()
		Standard.Mult(&Standard).Disp()
		Standard.Disp()
		errors++
	}
	if Standard.Mult(&EmptyNull).Degre() != -1 {
		t.Fail()
		errors++
	}
	P1 := Poly{Coefficients: []float64{1., 3.}}
	P2 := Poly{Coefficients: []float64{-2., 3.}}
	if P1.Mult(&P2).Coefficients[0] != -2. {
		t.Fail()
		errors++
	}
	if errors == 0 {
		PrintGreen("Inter-Polynomial multiplication works fully")
	}
}

func TestAdd(t *testing.T) {
	var P1, P2, P3, P0 *Poly
	var Pnot Poly
	P1 = &Poly{Coefficients: []float64{1., 2., 3.}}
	P2 = &Poly{Coefficients: []float64{1., 2., 3., 4.}}
	P3 = &Poly{Coefficients: []float64{1., 2., 3., 4., 5.}}
	P0 = &Poly{}
	errors := 0
	if P0.Add(P1).Degre() != P1.Degre() {
		t.Fail()
		errors++
	}
	if P3.Add(P2).Degre() != P3.Degre() {
		t.Fail()
		errors++
	}
	if P2.Add(P2).Degre() != P2.Degre() {
		t.Fail()
		errors++
	}
	if Pnot.Add(P3).Degre() != P3.Degre() {
		t.Fail()
		errors++
	}
	if errors == 0 {
		PrintGreen("Addition of polynoms seems to work")
	}
}

func TestMinus(t *testing.T) {
	var P1, P2, P3, P0 *Poly
	P1 = &Poly{Coefficients: []float64{1., 2., 3.}}
	P2 = &Poly{Coefficients: []float64{1., 2., 3., 4.}}
	P3 = &Poly{Coefficients: []float64{1., 2., 3., 4., 5.}}
	P0 = &Poly{}
	errors := 0
	if P0.Minus(P1).Degre() != P1.Degre() {
		fmt.Printf("P0: ")
		P0.Disp()
		fmt.Printf("P1: ")
		P1.Disp()
		fmt.Printf("P0 - P1: ")
		P0.Minus(P1).Disp()
		t.Error("Failed to subtract from null")
		errors++
	}
	if P3.Minus(P2).Degre() != P3.Degre() {
		fmt.Printf("P3: ")
		P3.Disp()
		fmt.Printf("P2: ")
		P2.Disp()
		fmt.Printf("P3 - P2: ")
		P3.Minus(P2).Disp()
		t.Error("Failed to subtract from larger polynom")
		errors++
	}
	if P3.Minus(P3).Degre() != P0.Degre() {
		t.Error("Failed to subtract from self")
		errors++
	}
	if errors == 0 {
		PrintGreen("Minus operator operationnal")
	}
}

func TestScalar_prod(t *testing.T) {
	var P0 *Poly
	P0 = &Poly{}
	P0.Scalar_prod_in_place(-1)
	var P1 *Poly
	P1 = &Poly{Coefficients: []float64{1.}}
	P1.Scalar_prod_in_place(3.)
	if P1.Coefficients[0] != 3. {
		t.Fail()
	}
	P1 = &Poly{Coefficients: []float64{1., 2., 3., 2., 1.}}
	P1.Scalar_prod_in_place(3.)
	if P1.Coefficients[4] != 3. {
		t.Fail()
	}
	if P1.Coefficients[2] != 9. {
		t.Fail()
	}
}

func TestDerive(t *testing.T) {
	var P0, P1 *Poly
	P1 = &Poly{Coefficients: []float64{0., 1.}}
	var Xs = []float64{1., 2., 25., -0.13}
	errors := 0
	if P0.Derive(0) != 0 {
		errors++
		t.Error("Failed to Derive")
	}
	for i, x := range Xs {
		if P0.Derive(x) != P0.GetDeriv().Call(x) || P0.Derive(x) != 0 {
			t.Error("Failed to get correct value for sample",
				i, ": x =", x, "[Null polynom]",
				"\n Got (", P0.Derive(x), ", ", P0.GetDeriv().Call(x), ") instead of 1")
		}
		if P1.Derive(x) != 1 || P1.GetDeriv().Call(x) != 1 {
			t.Error("Failed to get correct value for sample",
				i, ": x =", x, "[normal polynom]",
				"\n Got (", P1.Derive(x), ", ", P1.GetDeriv().Call(x), ") instead of 1")
		}
	}
	if errors == 0 {
		PrintGreen("Derivation doesn't veer")
	}
}

func TestIntegration(t *testing.T) {
	var P0, P1, P2 Poly
	errors := 0
	P1 = Poly{Coefficients: []float64{1.}}
	P2 = Poly{Coefficients: []float64{1., 2.}}
	var a, b float64
	a = 0.
	b = 15.
	fmt.Println("Testing between", a, "and", b, "...")
	if P0.GetPrimitive().Degre() > 0 {
		errors++
		t.Error("Primitive of null polynomial isn't made correctly (", P0.GetPrimitive().Degre(), ")")
	}
	if P0.Integrate(a, b) != 0. {
		errors++
		t.Error("Integate of null polynomial isn't a null function [", P0.Integrate(a, b), "]")
	}
	// Testing PP1
	PP1 := P1.GetPrimitive()
	if PP1.Call(b)-PP1.Call(a) != P1.Integrate(a, b) {
		errors++
		t.Error("Primitives doesn't equalize to correct Integration :",
			PP1.Call(b)-PP1.Call(a), "!=", P1.Integrate(a, b))
	}
	if PP1.GetDeriv().Coefficients[0] != P1.Coefficients[0] {
		errors++
		t.Error("Derivation of a primitive doesn't give back the original Polynomial")
	}
	// Testing PP2
	PP2 := P2.GetPrimitive()
	if PP2.Call(b)-PP2.Call(a) != P2.Integrate(a, b) {
		errors++
		t.Error("Primitive of null polynomial isn't made correctly")
	}
	if PP2.GetDeriv().Coefficients[1] != P2.Coefficients[1] {
		errors++
		t.Error("Derivation of a primitive doesn't give back the original Polynomial")
	}
	if errors == 0 {
		PrintGreen("Integration does work !")
	}
}

// tests polynomial composition
func TestComposite(t *testing.T) {
	var P0, P, Q, R *Poly
	P = &Poly{Coefficients: []float64{0., 0., 1.}} // P[X] = X²
	Q = &Poly{Coefficients: []float64{1., 1.}}     // Q[X] = X + 1
	R = &Poly{Coefficients: []float64{1.}}
	errors := 0
	if P0.Composite(Q).Degre() >= 0 {
		errors++
		t.Error("Expected null polynomial when composing null polynomial.")
	}
	if P.Composite(P0).Degre() >= 0 {
		errors++
		t.Error("Expected null polynomial when composing null polynomial")
	}
	if P.Composite(R).Degre() != 0 {
		errors++
		t.Error("Expected constant polynomial when composing constant polynomial")
	}
	if R.Composite(P).Degre() != 0 {
		errors++
		t.Error("Expected constant polynomial when composing from constant polynomial")
	}
	if P.Composite(Q).Degre() != 2 { // (X+1)² = X²+2X+1
		errors++
		t.Error("Composition does not give back the correct Polynoms :",
			P.Composite(Q).Degre(), "instead of", 2)
	} else if Q.Composite(P).Degre() != 2 { // X²+1
		errors++
		t.Error("Inverting the order of the polynomials",
			"does not give back the right value when compositing")
	}
	Xs := [...]float64{1., 2., 3., 4., 5.}
	for _, x := range Xs {
		if math.Abs(P.Composite(Q).Call(x)-P.Call(Q.Call(x))) > .00000000000005 {
			errors++
			t.Error("Composition does not return the right values")
		}
		if math.Abs(Q.Composite(P).Call(x)-Q.Call(P.Call(x))) > .00000000000005 {
			errors++
			t.Error("Composition with inverted order does",
				"not return the right values.",
				"\nExpected :", Q.Composite(P).Call(x), "Got :", Q.Call(P.Call(x)))
		}
	}
	if errors == 0 {
		PrintGreen("Composition stays on track")
	}
}

func TestPow(t *testing.T) {
	fmt.Println("Testing Pow")
	errors := 0
	var P = &Poly{Coefficients: []float64{0., 4.}}
	for i := 0; i < 10; i++ {
		PowP, err := P.Pow(i)
		if i == 0 {
			if PowP.Degre() != 0 || err != nil {
				errors++
				t.Error("Failed to handle null exponent")
			}
		} else {
			if PowP.Degre() != i {
				errors++
				t.Error("Failed to use valid exponent. Got", PowP.Degre(), "instead of", i)
			} else if err != nil {
				errors++
				t.Error(err.Error())
			}
			for k := 0; k < i; k++ {
				if PowP.Coefficients[k] != 0 {
					errors++
					t.Error("Invalid (non-zero) internal value in the Polynom. Does not match the example")
					break
				}
			}
			if math.Abs(PowP.Coefficients[i]-math.Pow(4, float64(i))) > .00000000000005 {
				errors++
				t.Error("Invalid value as head. Got",
					PowP.Coefficients[i], "instead of", math.Pow(4, float64(i)),
					"(Delta = ",
					math.Abs(PowP.Coefficients[i]-math.Pow(4, float64(i))), ")")
			}
		}
	}
	if errors == 0 {
		PrintGreen("Power is well within expectations")
	}
}

func TestEquality(t *testing.T) {
	A := &Poly{Coefficients: []float64{1., 2., 3.}}
	B := &Poly{Coefficients: []float64{1., 2., 3.}}
	C := &Poly{Coefficients: []float64{1., 2., 4.}}
	D := &Poly{}
	E := &Poly{Coefficients: []float64{1., 2.}}
	F := &Poly{Coefficients: []float64{0., 0., 0., 0.}}
	var errors int
	if !A.Eq(B) {
		errors++
		t.Error("A :", A, "and B :", B, "should be equals")
	}
	if A.Eq(C) {
		errors++
		t.Error("Cannot Distinguish two polynoms that aren't the same")
	}
	if E.Eq(A) {
		errors++
		t.Error("Cannot operate fully when comparing to longer polynomials")
	}
	if !D.Eq(F) {
		errors++
		t.Error("Two null polynomials in different forms are not detected as equals")
	}
	if errors == 0 {
		PrintGreen("Equality functions")
	}
}
