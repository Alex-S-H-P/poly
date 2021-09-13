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

### Polynoms

`Poly` are a go struct that are the center of this library. They are basically a list of coefficients. Polynomials will be considered null if they have no coefficients, if they are `nil`, or if all of their coefficients are null.

### Degree

You can access this value by simply using the Degre method. Despite degree being mathematically `-inf` for null Polys, in this library we have chosen to give it the value of **-1**. This value has been chosen so that the number of coefficients can be established as `p.Degre() + 1`

It is recommended to test first for negative degree to ensure that the Polynomials are not edge-cases (no coefficients, nil pointer etc.) As such, most functions in this package will test first for nil Polys and, should that case be met, give back a simpler answer that meets this case.

The Degree method will also remove leading zeros. Please note that due to floating point inaccuracies, some coefficients that should *mathematically* be null won't be and vice-versa. It is a rare occurence, however, and we have chosen to ignore it at least for the time being.

### Call and Composite

The `Call` method takes in a floating point input (float64) and returns the value of the polynomial function at this point.

The `Composite` function takes in a Polynomial (or rather it's pointer typed `*Poly`), and returns a pointer of a new Polynomial. This new `*Poly` will be made such that `P.Call(Q.Call(x)) == P.Composite(Q).Call(x)` whatever the value of P, Q both `*Poly` and x, a `float64`.

### Integration and derivation.

We have recognized the need for Integrating and Derivating polynomial as polynomial objects and as functions.

As such, despite there being two functionnalities, we 

----
WIP
----
