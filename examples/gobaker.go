package main

import (
	"flag"

	"github.com/pkg/profile"
	"github.com/rtropisz/gobaker"
)

var bakedDiffuse *gobaker.Texture
var bakedNormal *gobaker.Texture
var offset float64
var low gobaker.Mesh
var high gobaker.Mesh
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
	scene := gobaker.NewScene(*size)
	scene.Lowpoly.ReadOBJ(*lowName, false)
	scene.Highpoly.ReadOBJ(*highName, true)
	scene.Bake()
	scene.BakedDiffuse.SaveImage(*lowName + "_diff.png")
	scene.BakedNormal.SaveImage(*lowName + "_norm.png")
}
