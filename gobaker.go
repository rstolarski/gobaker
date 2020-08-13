package gobaker

import (
	"flag"

	"github.com/pkg/profile"
)

var bakedDiffuse *Texture
var bakedNormal *Texture
var offset float64
var low Mesh
var high Mesh
var s int

func main() {
	var (
		size      = flag.Int("s", 1024, "size of the output in pixels")
		lowName   = flag.String("l", "", "filename of lowpoly mesh")
		highName  = flag.String("h", "", "filename of lowpoly mesh")
		profiling = flag.Bool("p", false, "turn on trace profiling")
	)
	flag.Parse()

	//Profiling
	if *profiling {
		defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	}
	scene := NewScene(*size)
	scene.Lowpoly.ReadOBJ(*lowName, false)
	scene.Highpoly.ReadOBJ(*highName, true)
	scene.Bake()
	scene.BakedDiffuse.SaveImage(*lowName + "_diff.png")
	scene.BakedNormal.SaveImage(*lowName + "_norm.png")
}
