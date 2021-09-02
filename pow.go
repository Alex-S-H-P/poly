package polynom

import (
	"fmt"
)

// checks whether a bit in an int number is set
func hasBit(number int, pos uint) bool {
	val := number & (1 << pos)
	return (val > 0)
}

// sets the bit
func setBit(number uint64, pos uint) uint64 {
	number |= (1 << pos)
	return number
}

func pow(P *Poly, n uint64) *Poly {
	switch {
	case n == 0:
		return &Poly{Coefficients: []float64{1.}}
	case n == 1:
		return P
	case n%2 == 0: //n even
		return pow(P.Mult(P), n/2)
	default:
		return P.Mult(pow(P.Mult(P), (n-1)/2))
	}
}

// Allows for a polynomial to be squared, cubed and so on...
func (self *Poly) Pow(exponent int) (*Poly, error) {
	if exponent < 0 {
		return nil, fmt.Errorf("Cannot use negative exponent")
	}
	if exponent == 0 {
		return &Poly{Coefficients: []float64{1.}}, nil
	}
	n := uint64(exponent)
	return pow(self, n), nil

}
