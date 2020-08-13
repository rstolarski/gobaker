package gobaker

// Matrix represents a 3x3 matrix
type Matrix struct {
	x00, x01, x02 float64
	x10, x11, x12 float64
	x20, x21, x22 float64
}

// NewMatrix return new Matrix based
func NewMatrix(v0, v1, v2 Vector) Matrix {
	return Matrix{
		v0.X, v0.Y, v0.Z,
		v1.X, v1.Y, v1.Z,
		v2.X, v2.Y, v2.Z,
	}
}

// MulDirection multiplies Matrix 'a' by a Vector 'b'
// It returns a normalized Vector
func (a Matrix) MulDirection(b Vector) Vector {
	x := a.x00*b.X + a.x01*b.Y + a.x02*b.Z
	y := a.x10*b.X + a.x11*b.Y + a.x12*b.Z
	z := a.x20*b.X + a.x21*b.Y + a.x22*b.Z
	return Vector{x, y, z}.Normalize()
}

// Transpose returns a transposed 3x3 Matrix
func (a Matrix) Transpose() Matrix {
	return Matrix{
		a.x00, a.x10, a.x20,
		a.x01, a.x11, a.x21,
		a.x02, a.x12, a.x22}
}
