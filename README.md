![Logo unavailable](logo.jpg)

# Poly

Poly is a go library that allows for manipulation of polynomial object.

## Table of contents

- [OverView](#overview)
- [Detailed documentation](#detailed-documentation)

## Overview

### Generalities

This library is centered around the `Poly` struct. This struct has a `Coefficients` field, which gives the coefficients (the 0th coefficient is the constant coefficient, the 1st is the linear coefficient and so on ...). 
These coefficients are floating points values (`float64` in go).

Here's an example :
```go
P := Poly{Coefficients:[]float64{1., 2., 3.}} 
// P[X] = 1 + 2X + 3X²
// You can check this with
P.Disp()
```

You can then use the polynomial object as a normal function just by calling its `Call` method:
```go
fmt.Println(P.Call(0.)) // prints 1.000 because of the constant coefficient being 1
```

If you think you have found a mistake,
or have thought of an improvement, please either tell me or 
(even better) make a pull request.

### Importing the package

You can import this package by simply adding `import "github.com/Alex-S-H-P/poly"` at the top of your code file. 
You can look through the test files that are soon to come for help and inspiration on this matter. 

## Detailed documentation

WIP