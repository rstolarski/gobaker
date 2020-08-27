package gobaker

import (
	"fmt"
	"math"
)

// Vertex desribes a 3D mesh vertex
type Vertex struct {
	v  Vector
	vt Vector  // Texture coordinate
	vn Vector  // Normal Vector
	va float64 // Vertex Color Alpha
}

//SetVertexAlpha set vertex alpha color
func (v *Vertex) SetVertexAlpha(vAlpha float64) {
	v.va = vAlpha
}

// Triangle describes single triangle of a 3D mesh
type Triangle struct {
	V0, V1, V2 Vertex // Each triangle vertex
	Tangent    Vector
	Bitangent  Vector
	Normal     Vector
	Bar        Vector // Baricentric coordinates from intersection with a ray
	Material   *Material
	Distance   float64
	hitFront   bool
}

// String implements Stringer interface.
// It displays each triangle Vertex, Texture and Normals Vectors
func (t Triangle) String() string {
	var s string
	s += fmt.Sprintf("Verticies: %.5f, %.5f, %.5f\n", t.V0.v, t.V1.v, t.V2.v)
	s += fmt.Sprintf("Texture: %.5f, %.5f, %.5f\n", t.V0.vt, t.V1.vt, t.V2.vt)
	s += fmt.Sprintf("Normals: %.5f, %.5f, %.5f\n", t.V0.vn, t.V1.vn, t.V2.vn)
	s += fmt.Sprintf("Color: %.5f, %.5f, %.5f\n", t.V0.va, t.V1.va, t.V2.va)
	return s
}

// Intersect checks if ray intersects with a triangle
func (t *Triangle) Intersect(r *Ray) bool {
	a := t.V0.v
	b := t.V1.v
	c := t.V2.v

	t.Normal = (b.Sub(a)).Cross(c.Sub(a))

	if t.Normal.X > 1.0 || t.Normal.Y > 1.0 || t.Normal.Z > 1.0 {
		t.Normal = t.Normal.Normalize()
	}

	den := t.Normal.Dot(r.Direction)
	if math.Abs(den) < math.SmallestNonzeroFloat64 {
		return false
	}

	nom := a.Sub(r.Origin)
	d := nom.Dot(t.Normal) / den

	r.Distance = d
	if d >= math.SmallestNonzeroFloat64 {
		v0 := b.Sub(a)
		v1 := c.Sub(a)
		v2 := r.HitPosition().Sub(a)

		d00 := v0.Dot(v0)
		d01 := v0.Dot(v1)
		d11 := v1.Dot(v1)
		d20 := v2.Dot(v0)
		d21 := v2.Dot(v1)

		denom := d00*d11 - d01*d01
		beta := (d11*d20 - d01*d21) / denom
		gamma := (d00*d21 - d01*d20) / denom

		alpha := 1.0 - beta - gamma

		if beta < 0.0 || gamma < 0.0 || alpha < 0.0 {
			return false
		}
		t.Bar = Vector{alpha, beta, gamma}
		t.Distance = r.Distance
		return true
	}
	return false
}

// Barycentric return barycentric coorditnates based on UV position of a triangle.
func (t Triangle) Barycentric(u, v float64) Vector {
	xa := t.V0.vt.X
	ya := t.V0.vt.Y
	xb := t.V1.vt.X
	yb := t.V1.vt.Y
	xc := t.V2.vt.X
	yc := t.V2.vt.Y
	xp := u
	yp := v

	d := (yb-yc)*(xa-xc) + (xc-xb)*(ya-yc)
	d1 := ((yb-yc)*(xp-xc) + (xc-xb)*(yp-yc)) / d
	d2 := ((yc-ya)*(xp-xc) + (xa-xc)*(yp-yc)) / d
	d3 := 1 - d1 - d2

	return Vector{d1, d2, d3}

}
