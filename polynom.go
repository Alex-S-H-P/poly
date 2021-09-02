package poly

import (
	"fmt"
	"math"
)

// A polynom-like object. Can handle multiplication, integration, sum, etc.
type Poly struct {
	// The coefficients of the polynom object. The first coefficient corresponds to the 0th element, and this then continues
	// Null Polynoms have a nil coefficient
	Coefficients []float64
}

// A method for getting the Degre of a polynom
// Null Polynoms will have a -1 degree while constant polynoms will have a degre of zero.
// Overall, Polynoms have a their coefficients of len = self.Degre() + 1
func (self *Poly) Degre() int {
	if self == nil {
		return -1
	}
	for {
		if self.Coefficients == nil {
			return -1 // instead of the more mathematically correct -inf
		} else if len(self.Coefficients) == 0 {
			return -1
		}
		if self.Coefficients[len(self.Coefficients)-1] != 0.0 {
			break
		}
		self.Coefficients = self.Coefficients[:len(self.Coefficients)-1]
	}
	return len(self.Coefficients) - 1
}

// allows for a composition between two Polynomials. If we were to do
//		R := P.Composite(Q), then
//		R.Call(x) == P.Call(Q.Call(x))
func (self *Poly) Composite(other *Poly) *Poly {
	var sD, oD int
	sD = self.Degre()
	oD = other.Degre()
	if sD < 0 {
		return &Poly{}
	} else if oD < 0 {
		return &Poly{Coefficients: []float64{self.Call(0.)}}
	} else if sD == 0 {
		return &Poly{Coefficients: self.Coefficients[:]}
	} else if oD == 0 {
		return &Poly{Coefficients: []float64{self.Call(other.Call(0.))}}
		// other is constant
	}
	// we now try to do it normally
	var result, P *Poly
	// handling the constant part fitst
	result = &Poly{Coefficients: []float64{self.Coefficients[0]}}
	P = other
	// now for the part actually afffected by the composition
	for i := 1; i < len(self.Coefficients); i++ {
		result = result.Add(P.Scalar_prod(self.Coefficients[i]))
		P = P.Mult(other)
	}

	return result
}

// Returns the area under the Polynom's curve in the [a, b] segment.
func (self *Poly) Integrate(a, b float64) float64 {
	if b < a {
		return -self.Integrate(b, a)
	}
	if self.Degre() < 0 {
		return 0
	}
	sum := 0.
	for i := 0; i < len(self.Coefficients); i++ {
		j := float64(i)
		sum += self.Coefficients[i] / (j + 1) * (math.Pow(b, j+1) - math.Pow(a, j+1))
	}
	return sum
}

// Creates a Primitive of the Polynom self.
// The primitive is itself a Polynom object
// The integration constant is set to 0 (the 0th Coefficient is null)
func (self *Poly) GetPrimitive() *Poly {
	if self.Degre() < 0 {
		return &Poly{Coefficients: []float64{}}
	}
	coeffs := make([]float64, self.Degre()+2)
	for i := 0; i < len(coeffs)-1; i++ {
		j := float64(i)
		coeffs[i+1] = self.Coefficients[i] / (j + 1)
	}
	return &Poly{Coefficients: coeffs}
}

// Creates a derivative polynom of self.
// Please do remember that taking the primitive of this derived polynom will not get you your original longer
func (self *Poly) GetDeriv() *Poly {
	if self.Degre() < 0 { // we have a null polynom
		return &Poly{Coefficients: []float64{}}
	}
	coeffs := make([]float64, self.Degre())
	for i := 0; i < len(coeffs); i++ {
		j := float64(i)
		coeffs[i] = self.Coefficients[i+1] * (j + 1)
	}
	return &Poly{Coefficients: coeffs}
}

// Derives the polynom and gives the result at x
func (self *Poly) Derive(x float64) float64 {
	var sum = 0.
	for i := 1; i < self.Degre()+1; i++ {
		j := float64(i)
		sum += self.Coefficients[i] * j * math.Pow(x, j-1)
	}
	return sum
}

func (self *Poly) String() string {
	var t string
	t = "P[X] = "
	if self.Degre() < 0 {
		t += "0"
		return t
	}
	var n = len(self.Coefficients)
	for i := 0; i < n; i++ {
		if self.Coefficients[i] != 0. {
			if i == 0 {
				t += fmt.Sprintf("%f", self.Coefficients[i])
			} else if i == 1 {
				t += fmt.Sprintf("%f . X", self.Coefficients[i])
			} else {
				t += fmt.Sprintf("%f . X^%d", self.Coefficients[i], i)
			}
			if i < n-1 {
				t += " + "
			}
		}
	}
	return t
}

func (self *Poly) Disp() {
	fmt.Println(self.String())
}
