![Logo unavailable](logo.jpg)

# Poly

Poly is a go library that allows for manipulation of polynomial object.

## Table of contents

- [OverView](#overview)
- [Detailed documentation](#detailed-documentation)

## Overview

This library is centered around the `Poly` struct. This struct has a `Coefficients` field, which gives the coefficients (the 0th coefficient is the constant coefficient, the 1st is the linear coefficient and so on ...). 
These coefficients are floating points values (float64 in go)

Here's an example :
```go
P := Poly{Coefficients:[]float64{1., 2., 3.}}
```

## Detailed documentation