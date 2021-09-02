package poly

import (
	"math"
)

// Gets the value of the polynomial function for input x
func (self *Poly) Call(x float64) float64 {
	if self == nil {
		return 0.
	}
	n := float64(len(self.Coefficients))
	sum := 0.
	var i float64
	for i = 0.; i < n; i += 1 {
		sum += self.Coefficients[int(i)] * math.Pow(x, i)
	}
	return sum
}

// tests Whether two polynoms are the same.
// P.Eq(Q) <=> P.Call(x) == Q.Call(x) for all x (real)
func (self *Poly) Eq(other *Poly) bool {
	if self.Degre() != other.Degre() {
		return false
	} else if self.Degre() < 0 {
		return true
	}
	for i := 0; i < len(self.Coefficients); i++ {
		if self.Coefficients[i] != other.Coefficients[i] {
			return false
		}
	}
	return true
}

// Multiplies two polynoms together with standard polynomial product
func (self *Poly) Mult(other *Poly) *Poly {
	coeffs := make([]float64, len(self.Coefficients)*len(other.Coefficients))
	if self.Coefficients == nil || other.Coefficients == nil {
		// a null Polynom will always return a null Polynom, no matter what it's being multiplied with
		return &Poly{}
	}
	for i := 0; i < len(self.Coefficients); i++ {
		for j := 0; j < len(other.Coefficients); j++ {
			coeffs[i+j] += self.Coefficients[i] * other.Coefficients[j]
		}
	}
	return &Poly{Coefficients: coeffs}
}

// Multiplies a polynomial by a scalar. Returns the result
func (self *Poly) Scalar_prod(x float64) *Poly {
	if self.Degre() < 0 {
		return nil
	}
	coeffs := make([]float64, len(self.Coefficients))
	for i := 0; i < len(self.Coefficients); i++ {
		coeffs[i] = self.Coefficients[i] * x
	}
	return &Poly{Coefficients: coeffs}
}

// Multiplies all of the coefficients of a Polynom. Is done in place.
func (self *Poly) Scalar_prod_in_place(x float64) {
	if self.Degre() < 0 {
		return
	}
	for i := 0; i < len(self.Coefficients); i++ {
		self.Coefficients[i] *= x
	}
}

// Adds two polynoms together
func (self *Poly) Add(other *Poly) *Poly {

	// we handle cases of polynoms having different degrees
	var n_max, n_min int
	var longer, shorter *Poly
	if self.Degre() > other.Degre() {
		longer = self
		shorter = other
		n_max = self.Degre() + 1 // degre = len - 1
		n_min = other.Degre() + 1
	} else {
		longer = other
		shorter = self
		n_max = other.Degre() + 1 // degre = len - 1
		n_min = self.Degre() + 1
	}
	coeffs := make([]float64, n_max)
	for i := 0; i < n_max; i++ {
		if i < n_min {
			coeffs[i] = longer.Coefficients[i] + shorter.Coefficients[i]
		} else {
			coeffs[i] = longer.Coefficients[i]
		}
	}
	return &Poly{Coefficients: coeffs}
}

// subracts the other polynom from self
func (self *Poly) Minus(other *Poly) *Poly {
	// we handle cases of polynoms having different degrees
	if self.Degre() < other.Degre() {
		P := other.Minus(self)
		P.Scalar_prod_in_place(-1.)
		return P
	}
	coeffs := make([]float64, len(self.Coefficients))
	for i := 0; i < len(self.Coefficients); i++ {
		if i < len(other.Coefficients) {
			coeffs[i] = self.Coefficients[i] - other.Coefficients[i]
		} else {
			coeffs[i] = self.Coefficients[i]
		}
	}
	return &Poly{Coefficients: coeffs}
}
