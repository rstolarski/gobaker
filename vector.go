package gobaker

import (
	"fmt"
	"image/color"
	"math"
	"strconv"
	"strings"
)

// Vector defines 3 point or axis with X,Y,Z values
type Vector struct {
	X, Y, Z float64
}

// Unit Vectors
var (
	Zero    = Vector{0, 0, 0}
	One     = Vector{1, 1, 1}
	Forward = Vector{0, 1, 0}
	Back    = Vector{0, -1, 0}
	Left    = Vector{-1, 0, 0}
	Right   = Vector{1, 0, 0}
	Up      = Vector{0, 0, 1}
	Down    = Vector{0, 0, -1}
)

// Add adds two Vectors together
func (a Vector) Add(b Vector) Vector {
	return Vector{a.X + b.X, a.Y + b.Y, a.Z + b.Z}
}

// Sub subtracts set Vector from current one
func (a Vector) Sub(b Vector) Vector {
	return Vector{a.X - b.X, a.Y - b.Y, a.Z - b.Z}
}

// Mul multiplies Vector by a scalar
func (a Vector) Mul(n float64) Vector {
	return Vector{a.X * n, a.Y * n, a.Z * n}
}

// MulVec multiplies Vector by a Vector
func (a Vector) MulVec(b Vector) Vector {
	return Vector{a.X * b.X, a.Y * b.Y, a.Z * b.Z}
}

// Div divides Vector by scalar
func (a Vector) Div(b float64) Vector {
	return Vector{a.X / b, a.Y / b, a.Z / b}
}

// DivVec divides each Vector value by corresponding value from second Vector
func (a Vector) DivVec(b Vector) Vector {
	return Vector{a.X / b.X, a.Y / b.X, a.Z / b.X}
}

// Dot returns the dot product of two Vector's.
func (a Vector) Dot(b Vector) float64 {
	return a.X*b.X + a.Y*b.Y + a.Z*b.Z
}

// Cross returns the cross product of two Vectors.
func (a Vector) Cross(b Vector) Vector {
	return Vector{a.Y*b.Z - a.Z*b.Y, a.Z*b.X - a.X*b.Z, a.X*b.Y - a.Y*b.X}
}

// Negate returns negated Vector.
func (a Vector) Negate() Vector {
	return Vector{-a.X, -a.Y, -a.Z}
}

// Len finds the length of the Vector.
func (a Vector) Len() float64 {
	return math.Sqrt(a.X*a.X + a.Y*a.Y + a.Z*a.Z)
}

// Equals compares two Vectors.
func (a Vector) Equals(b Vector) bool {
	return a.X == b.X && a.Y == b.Y && a.Z == b.Z
}

// Normalize return Normalize current Vector.
func (a Vector) Normalize() Vector {
	l := a.Len()
	return Vector{a.X / l, a.Y / l, a.Z / l}
}

// Square returns new Vector with each values squared.
func (a Vector) Square() Vector {
	return Vector{math.Sqrt(a.X), math.Sqrt(a.Y), math.Sqrt(a.Z)}
}

// Lerp interpolates each Vector's value
// with another Vector, scaling it by n.
func (a Vector) Lerp(b Vector, n float64) Vector {
	m := 1 - n
	return Vector{a.X*m + b.X*n, a.Y*m + b.Y*n, a.Z*m + b.Z*n}
}

// ColorToFloat converts uint8 values into 0.0-1.0 number in Vector format.
func ColorToFloat(r, g, b uint8) Vector {
	return Vector{float64(r) / 255, float64(g) / 255, float64(b) / 255}
}

// FloatToColor converts 0.0-1.0 number into color.NRGBA format.
func (a Vector) FloatToColor() color.NRGBA {
	a = a.Mul(255)

	return color.NRGBA{uint8(a.X), uint8(a.Y), uint8(a.Z), 255}
}

// Abs negates each Vectors component if it's lower then zero.
func (a Vector) Abs() Vector {
	return Vector{math.Abs(a.X), math.Abs(a.Y), math.Abs(a.Z)}
}

// Barycentric performs a barycentric interpolation between
// 3 Vectors 'a', 'b' and 'c' based on values of a Vector 'l'.
func Barycentric(a, b, c, l Vector) Vector {
	return a.Mul(l.X).Add(b.Mul(l.Y)).Add(c.Mul(l.Z))
}

// ParseVector parses string values seperated with comma into a Vector.
func ParseVector(s string) (v Vector, err error) {
	xyz := strings.Split(s, ",")
	if len(xyz) != 3 {
		return v, fmt.Errorf(
			"3 values are required for Vector, received %v values",
			len(xyz),
		)
	}
	v.X, err = strconv.ParseFloat(xyz[0], 64)
	if err != nil {
		return v, err
	}
	v.Y, err = strconv.ParseFloat(xyz[1], 64)
	if err != nil {
		return v, err
	}
	v.Z, err = strconv.ParseFloat(xyz[2], 64)
	return v, err
}

// String implements Stringer interface.
// It displays each Vector components value seperated with a comma.
func (a Vector) String() string {
	return fmt.Sprintf("%.5f, %.5f, %.5f", a.X, a.Y, a.Z)
}
