package polynom

import "fmt"

// A 2D point. Meant mostly for internal handling of the Interpolation
type Point struct {
	x float64
	y float64
}

func Build(x, y interface{}) *Point {
	var X, Y float64
	switch t1 := x.(type) {
	case int:
		X = float64(t1)
	case float64:
		X = t1
	case float32:
		X = float64(t1)
	}
	switch t2 := y.(type) {
	case int:
		Y = float64(t2)
	case float64:
		Y = t2
	case float32:
		Y = float64(t2)
	}
	return &Point{x: X, y: Y}
}

// Transforms 2 arrays into an array of points
func ToPoints(Xs, Ys []float64) ([]Point, error) {
	var n int
	var err error
	if len(Xs)*len(Ys) == 0 {
		return nil, fmt.Errorf("Contained an empty array")
	}
	if len(Xs) != len(Ys) {
		n = len(Xs)
		if len(Ys) < n {
			n = len(Ys)
		}
		err = fmt.Errorf("Arrays of different size. Decided to take the first %d elements", n)
	} else {
		n = len(Xs)
	}
	arr := make([]Point, 0, n)
	for i := 0; i < n; i++ {
		arr = append(arr, Point{x: Xs[i],
			y: Ys[i]})
	}
	return arr, err
}
