package main

import (
	"flag"
	"strings"

	"github.com/pkg/profile"
	"github.com/rtropisz/gobaker"
)

func main() {
	var (
		size      = flag.Int("s", 1024, "size of the output in pixels")
		lowName   = flag.String("l", "", "pathToFile of lowpoly mesh")
		highName  = flag.String("h", "", "pathToFile of lowpoly mesh")
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
	scene.BakedDiffuse.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_diff.png")
	scene.BakedNormal.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_norm.png")
}
