package main

import (
	"flag"
	"log"
	"strings"
	"time"

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
	log.Printf("Starting")
	start := time.Now()
	scene.Lowpoly.ReadOBJ(*lowName, false)
	log.Printf("Reading lowpoly mesh.. %s", time.Since(start))
	scene.Highpoly.ReadOBJ(*highName, true)
	log.Printf("Reading highpoly mesh.. %s", time.Since(start))
	scene.Bake()
	log.Printf("Baking...")
	log.Printf("Finished baking took %s", time.Since(start))
	scene.BakedDiffuse.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_diff.png")
	scene.BakedNormal.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_norm.png")
	log.Printf("Finished saving images took %s", time.Since(start))
}
