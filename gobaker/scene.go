package gobaker

import (
	"log"
	"sort"
	"sync"
	"time"

	bep "github.com/gen2brain/beeep"
)

// Scene represents an object in which baking process is being handled
// It cointais a lowpoly and highpoly meshes and output textures
// with their final size
type Scene struct {
	Lowpoly      Mesh
	Highpoly     Mesh
	BakedDiffuse *Texture
	//BakedNormal       *Texture
	//BakedObjectNormal *Texture
	BakedID            *Texture
	OutputSize         int
	ReadID             bool
	MaxFrontalDistance float64
	MaxRearDistance    float64
}

// NewScene return a new Scene with output textures of a given size 's'
func NewScene(s int, readID bool, MaxFrontalDistance, MaxRearDistance float64) Scene {
	return Scene{
		BakedDiffuse: NewTexture(s),
		//BakedNormal:       NewTexture(s),
		//BakedObjectNormal: NewTexture(s),
		BakedID:            NewTexture(s),
		OutputSize:         s,
		ReadID:             readID,
		MaxFrontalDistance: MaxFrontalDistance,
		MaxRearDistance:    MaxRearDistance,
	}
}

// ReadTexturesForHighpoly read textures into highpoly model
func (s *Scene) ReadTexturesForHighpoly() (err error) {
	err = s.Highpoly.ReadTexturesToMaterials(s.ReadID)
	if err != nil {
		return err
	}
	return nil
}

// Bake processes each pixel of an output texture
// it computes each texture in one process\
func (s *Scene) Bake(workers int) {
	defer duration(track("Baking took"))
	// Get current number of CPU threads

	// Offset used in UV coordinate calculations
	offset := 1.0 / (2.0 * float64(s.OutputSize))

	depth := make([]float64, s.OutputSize*s.OutputSize)

	c := make(chan int, s.OutputSize)
	var wg sync.WaitGroup
	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func() {
			for row := range c {
				for col := 0; col < s.OutputSize; col++ {
					depth[col+s.OutputSize*row] = s.processPixel(col, row, offset)
				}
			}
			wg.Done()
		}()
	}

	for row := 0; row < s.OutputSize; row++ {
		c <- row
	}
	close(c)
	wg.Wait()

	if s.ReadID {
		// added alpha channel from depth array
		maxDistance := -1.0

		for row := 0; row < s.OutputSize; row++ {
			for col := 0; col < s.OutputSize; col++ {
				if depth[col+s.OutputSize*row] <= 0 {
					depth[col+s.OutputSize*row] = 0
				} else {
					if depth[col+s.OutputSize*row] > maxDistance {
						maxDistance = depth[col+s.OutputSize*row]
					}
				}
			}
		}

		for y := s.BakedID.Image.Bounds().Min.Y; y < s.BakedID.Image.Bounds().Max.Y; y++ {
			for x := s.BakedID.Image.Bounds().Min.X; x < s.BakedID.Image.Bounds().Max.X; x++ {
				sample := s.BakedID.Image.NRGBAAt(x, y)
				sample.A = uint8(depth[x+s.OutputSize*y] / maxDistance * 255.0)
				s.BakedID.Image.Set(x, y, sample)
			}
		}
	}

	bep.Alert("BAKER", "Baking finished!", "cooper.png")
}

func (s *Scene) processPixel(x, y int, offset float64) float64 {
	// Get uv coordinates for a given pixel
	uv := Vector{
		(float64(x) / float64(s.BakedDiffuse.h)) + offset,
		(float64(y) / float64(s.BakedDiffuse.h)) + offset,
		0.0,
	}

	// Iterate through all low poly triangles and check if current
	// uv coordinates are inside a given triangle
	var lowpolyTriangle *Triangle

	for i := 0; i < len(s.Lowpoly.Triangles); i++ {
		if checkIfInside(
			s.Lowpoly.Triangles[i].V0.vt.X,
			s.Lowpoly.Triangles[i].V0.vt.Y,
			s.Lowpoly.Triangles[i].V1.vt.X,
			s.Lowpoly.Triangles[i].V1.vt.Y,
			s.Lowpoly.Triangles[i].V2.vt.X,
			s.Lowpoly.Triangles[i].V2.vt.Y,
			uv.X,
			uv.Y,
		) {
			lowpolyTriangle = &s.Lowpoly.Triangles[i] // If it intersects with a triangle, stop the loop
			break
		}
	}
	if lowpolyTriangle == nil {
		return -1.0
	}

	v0 := lowpolyTriangle.V0.v
	v1 := lowpolyTriangle.V1.v
	v2 := lowpolyTriangle.V2.v

	n0 := lowpolyTriangle.V0.vn
	n1 := lowpolyTriangle.V1.vn
	n2 := lowpolyTriangle.V2.vn

	b := lowpolyTriangle.Barycentric(uv.X, uv.Y)

	// Origin is an origin point of a ray, that is from lowpoly mesh
	origin := Barycentric(v0, v1, v2, b)

	// Ray shoot the same direction as lowpoly triangle normal direction
	// direction := (v1.Sub(v0)).Cross(v2.Sub(v0))
	direction := Barycentric(n0, n1, n2, b)

	// Create new Ray for back and front shooting
	rayFront := Ray{origin, direction.Normalize().Negate(), 0.0, s.MaxFrontalDistance, s.MaxRearDistance}
	rayBack := Ray{origin, direction.Normalize(), 0.0, s.MaxFrontalDistance, s.MaxRearDistance}

	// List of highpoly triangles that are going to be hit by Rays
	highpolyHit := make([]Triangle, 0)

	// Check interstions with each highpoly triangles
	for i := 0; i < len(s.Highpoly.Triangles); i++ {
		if s.Highpoly.Triangles[i].Intersect(&rayFront) {
			if rayFront.Distance <= rayFront.MaxFrontalDistance {
				highpolyHit = append(highpolyHit, s.Highpoly.Triangles[i])
			}
		}
		if s.Highpoly.Triangles[i].Intersect(&rayBack) {
			if rayBack.Distance <= rayBack.MaxRearDistance {
				s.Highpoly.Triangles[i].hitFront = false
				s.Highpoly.Triangles[i].Distance = -s.Highpoly.Triangles[i].Distance
				highpolyHit = append(highpolyHit, s.Highpoly.Triangles[i])
			}

		}
	}

	// Return early
	if len(highpolyHit) == 0 {
		return -1.0
	}

	// Sort each hit triangle by a hit distance from longest to shortest
	// This is to ensure that we get farthest triangles first and then the closest ones
	sort.SliceStable(highpolyHit, func(i, j int) bool {
		return highpolyHit[i].Distance > highpolyHit[j].Distance
	})

	for i := 0; i < len(highpolyHit); i++ {
		// Get barycentric coordinates on a given highpoly triangle intesection
		uvhighpolyHit := Barycentric(highpolyHit[i].V0.vt, highpolyHit[i].V1.vt, highpolyHit[i].V2.vt, highpolyHit[i].Bar)

		// Sample color from highpoly texture based on triangle's material
		highpolyHitDiffuseColor := highpolyHit[i].Material.Diffuse.SamplePixel(uvhighpolyHit.X, uvhighpolyHit.Y)

		//	Alpha checking
		if highpolyHitDiffuseColor.A <= uint8(20) {
			continue
		}

		// Setting output diffuse texture color
		s.BakedDiffuse.Image.SetNRGBA(x, y, highpolyHitDiffuseColor)

		// ID map baking
		if s.ReadID {
			highpolyHitIDColor := highpolyHit[i].Material.ID.SamplePixel(uvhighpolyHit.X, uvhighpolyHit.Y)

			// blue color multiply by va
			blueColor := float64(highpolyHitIDColor.B) / 255.0
			blueColor *= (highpolyHit[i].V0.va*highpolyHit[i].Bar.X + highpolyHit[i].V1.va*highpolyHit[i].Bar.Y + highpolyHit[i].V2.va*highpolyHit[i].Bar.Z)
			highpolyHitIDColor.B = uint8(255.0 * blueColor)

			// Setting output ID texture color
			s.BakedID.Image.SetNRGBA(x, y, highpolyHitIDColor)
		}

		//an attempt to rendering normals :P
		//normalAthighpolyHit := Barycentric(t.V0.vn, t.V1.vn, t.V2.vn, t.Bar).Normalize()

		// // Calculating TBN matrix
		// v0 := t.V0
		// v1 := t.V1
		// v2 := t.V2

		// Edge1 := v1.v.Sub(v0.v)
		// Edge2 := v2.v.Sub(v0.v)

		// DeltaU1 := v1.vt.X - v0.vt.X
		// DeltaV1 := v1.vt.Y - v0.vt.Y
		// DeltaU2 := v2.vt.X - v0.vt.X
		// DeltaV2 := v2.vt.Y - v0.vt.Y

		// f := 1.0 / (DeltaU1*DeltaV2 - DeltaU2*DeltaV1)

		// T, B, N := Vector{}, Vector{}, Vector{}

		// T.X = f * (DeltaV2*Edge1.X - DeltaV1*Edge2.X)
		// T.Y = f * (DeltaV2*Edge1.Y - DeltaV1*Edge2.Y)
		// T.Z = f * (DeltaV2*Edge1.Z - DeltaV1*Edge2.Z)

		// B.X = f * (-DeltaU2*Edge1.X - DeltaU1*Edge2.X)
		// B.Y = f * (-DeltaU2*Edge1.Y - DeltaU1*Edge2.Y)
		// B.Z = f * (-DeltaU2*Edge1.Z - DeltaU1*Edge2.Z)

		// N = Edge1.Cross(Edge2).Normalize()
		// T = T.Normalize()
		// B = B.Normalize()

		// T = T.Sub(N.Mul(T.Dot(N)))
		// B = B.Sub(N.Mul(B.Dot(N))).Sub(T.Mul(B.Dot(T)))

		// TBN := NewMatrix(T, B, N)

		// // v1 - Just one vector
		// normalAthighpolyHit = TBN.MulDirection(normalAthighpolyHit).Normalize()

		// // v2 - Matrix by triangle normals
		// n0 := TBN.MulDirection(v0.vn).Normalize()
		// n1 := TBN.MulDirection(v1.vn).Normalize()
		// n2 := TBN.MulDirection(v2.vn).Normalize()

		// // Barycentric for no reason
		// normalAthighpolyHit = Barycentric(n0, n1, n2, t.Bar).Normalize()

		//if t.distance < 0 {
		//normalAthighpolyHit = normalAthighpolyHit.Mul(-1.0)
		//}
		//normalAthighpolyHit = normalAthighpolyHit.Add(One).Div(2.0)

		// Saving pixels to images
		//s.BakedObjectNormal.Image.SetNRGBA(x, y, normalAthighpolyHit.FloatToColor())
		//s.BakedNormal.Image.SetNRGBA(x, y, t.Material.Normal.SamplePixel(uvhighpolyHit.X, uvhighpolyHit.Y))
		return highpolyHit[i].Distance
	}
	return -1.0
}

// Checking if point with coordinates 'xp' and 'yp' is inside triangle
// with coordinates x1-x3 and y1-y3
func checkIfInside(x1, y1, x2, y2, x3, y3, xp, yp float64) bool {
	x2 -= x1
	y2 -= y1
	x3 -= x1
	y3 -= y1
	xp -= x1
	yp -= y1
	d := x2*y3 - x3*y2
	w1 := xp*(y2-y3) + yp*(x3-x2) + x2*y3 - x3*y2
	w2 := xp*y3 - yp*x3
	w3 := yp*x2 - xp*y2
	if w1 >= 0.0 && w1 <= d && w2 >= 0.0 && w2 <= d && w3 >= 0.0 && w3 <= d {
		return true
	}
	return false
}

func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %s\n", msg, time.Since(start))
}
