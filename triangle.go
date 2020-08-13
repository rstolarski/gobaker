package gobaker

import (
	"fmt"
	"math"
)

// Vertex desribes a 3D mesh vertex
type Vertex struct {
	v  Vector
	vt Vector // Texture coordinate
	vn Vector // Normal Vector
}

// Triangle describes single triangle of a 3D mesh
type Triangle struct {
	V0, V1, V2 Vertex // Each triangle vertex
	Tangent    Vector
	Bitangent  Vector
	Normal     Vector
	Bar        Vector // Baricentric coordinates from intersection with a ray
	Color      Vector
	Material   *Material
	distance   float64
	hitFront   bool
}

// String implements Stringer interface.
// It displays each triangle Vertex, Texture and Normals Vectors
func (t Triangle) String() string {
	var s string
	s += fmt.Sprintf("Verticies: %.5f, %.5f, %.5f\n", t.V0.v, t.V1.v, t.V2.v)
	s += fmt.Sprintf("Texture: %.5f, %.5f, %.5f\n", t.V0.vt, t.V1.vt, t.V2.vt)
	s += fmt.Sprintf("Normals: %.5f, %.5f, %.5f\n", t.V0.vn, t.V1.vn, t.V2.vn)
	s += fmt.Sprintf("Colors: %.5f, %.5f, %.5f\n", t.Color.X, t.Color.Y, t.Color.Z)
	return s
}

// Intersect checks if ray intersects with a triangle
func (t *Triangle) Intersect(r *Ray) bool {
	a := t.V0.v
	b := t.V1.v
	c := t.V2.v

	t.Normal = (b.Sub(a)).Cross(c.Sub(a)).Normalize()

	p := plane{Position: a, Normal: t.Normal}
	if !p.Intersect(r) {
		return false
	}

	// v0 := b.Sub(a)
	// v1 := c.Sub(a)
	// v2 := r.HitPosition().Sub(a)

	// d00 := v0.Dot(v0)
	// d01 := v0.Dot(v1)
	// d11 := v1.Dot(v1)
	// d20 := v2.Dot(v0)
	// d21 := v2.Dot(v1)

	// denom := d00*d11 - d01*d01
	// beta := (d11*d20 - d01*d21) / denom
	// gamma := (d00*d21 - d01*d20) / denom
	// a := v0
	// b := v1
	// c := v2
	o := r.Origin
	d := r.Direction

	M := [3][3](float64){
		{a.X - b.X, a.X - c.X, d.X},
		{a.Y - b.Y, a.Y - c.Y, d.Y},
		{a.Z - b.Z, a.Z - c.Z, d.Z}}

	Mbeta := [3][3](float64){
		{a.X - o.X, a.X - c.X, d.X},
		{a.Y - o.Y, a.Y - c.Y, d.Y},
		{a.Z - o.Z, a.Z - c.Z, d.Z}}

	Mgamma := [3][3](float64){
		{a.X - b.X, a.X - o.X, d.X},
		{a.Y - b.Y, a.Y - o.Y, d.Y},
		{a.Z - b.Z, a.Z - o.Z, d.Z}}

	detBeta := det(Mbeta)
	detGamma := det(Mgamma)
	detM := det(M)
	beta := detBeta / detM
	gamma := detGamma / detM

	alpha := 1.0 - beta - gamma
	t.Bar = Vector{alpha, beta, gamma}

	if beta < 0.0 || gamma < 0.0 || alpha < 0.0 {
		return false
	}

	p.Intersect(r)
	t.distance = r.Distance

	uv0 := t.V0.vt
	uv1 := t.V1.vt
	uv2 := t.V2.vt

	deltaPos1 := b.Sub(a)
	deltaPos2 := c.Sub(a)

	deltaUV1 := uv1.Sub(uv0)
	deltaUV2 := uv2.Sub(uv0)

	f := 1.0 / (deltaUV1.X*deltaUV2.Y - deltaUV1.Y*deltaUV2.X)
	t.Tangent = (deltaPos1.Mul(deltaUV2.Y).Sub(deltaPos2.Mul(deltaUV1.Y))).Mul(f)
	t.Bitangent = (deltaPos2.Mul(deltaUV1.X).Sub(deltaPos1.Mul(deltaUV2.X))).Mul(f)

	t.Tangent = t.Tangent.Normalize()
	t.Bitangent = t.Bitangent.Normalize()
	return true
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

func det(a [3][3]float64) float64 {
	r1 := (a[0][0] * a[1][1] * a[2][2]) + (a[0][1] * a[1][2] * a[2][0]) + (a[0][2] * a[1][0] * a[2][1])
	r2 := (a[0][2] * a[1][1] * a[2][0]) + (a[0][0] * a[1][2] * a[2][1]) + (a[0][1] * a[1][0] * a[2][2])

	return (r1 - r2)
}

type plane struct {
	Position Vector
	Normal   Vector
}

func (p *plane) Intersect(r *Ray) bool {
	denom := p.Normal.Dot(r.Direction)
	nom := p.Position.Sub(r.Origin)
	d := nom.Dot(p.Normal) / denom

	if math.Abs(denom) < math.SmallestNonzeroFloat64 {
		return false
	}

	r.Distance = d
	if d >= math.SmallestNonzeroFloat64 {
		return true
	}
	return false
}
