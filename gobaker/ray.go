package gobaker

// Ray describes a 3-dimensional ray with an origin and a direction Vector
// It also contains two distance values, used to checks intersection with
// multiple surfaces
type Ray struct {
	Origin, Direction  Vector
	Distance           float64
	MaxFrontalDistance float64
	MaxRearDistance    float64
}

// HitPosition return 3D coordinates of a ray intersetion
func (r Ray) HitPosition() Vector {
	return r.Origin.Add(r.Direction.Mul(r.Distance))
}
