package poly

import "fmt"

type PointLike interface {
	ToPoint() Point
}

// Point implements PointLike
func (self Point) ToPoint() Point {
	return self
}

// Realises a lagrangian interpolation of the data_series.
func Interpolate(Xs, Ys []float64) (*Poly, error) {
	points, err := ToPoints(Xs, Ys)
	if err != nil {
		if points == nil {
			return nil, err
		}
	}
	res, err2 := TrueInterpolate(points)
	if err2 == nil && err != nil {
		return res, err
	} else if err2 != nil && err != nil {
		return res, fmt.Errorf(err.Error() + "\nOn top of this\n" + err2.Error())
	} else if err2 == nil {
		return res, nil
	} else {
		return res, err2
	}
}

// Actually makes the interpolation. To simply give two arrays to Interpolate,
// just call the Interpolate function
func TrueInterpolate(points []Point) (*Poly, error) {
	var subPoly, final *Poly
	n := len(points)
	final = &Poly{}
	for i := 0; i < n; i++ {
		subPoly = nil
		for j := 0; j < n; j++ {
			if j != i {
				tmp_poly := &Poly{Coefficients: []float64{-points[j].x, 1.}}
				tmp_poly.Scalar_prod_in_place(1 / (points[i].x - points[j].x))
				if subPoly == nil {
					subPoly = tmp_poly
				} else {
					subPoly = subPoly.Mult(tmp_poly)
				}
			}
		}
		subPoly.Scalar_prod_in_place(points[i].y)
		final = final.Add(subPoly)
	}
	return final, nil
}
