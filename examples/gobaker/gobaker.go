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
		size            = flag.Int("s", 1024, "size of the output images in pixels")
		lowName         = flag.String("l", "", "path to lowpoly mesh")
		highName        = flag.String("h", "", "path to highpoly mesh")
		highPLYName     = flag.String("hp", "", "path to highpoly PLY mesh")
		cpuProfiling    = flag.Bool("cpuP", false, "turn on cpu profiling")
		memProfiling    = flag.Bool("memP", false, "turn on memory profiling")
		tracecProfiling = flag.Bool("traceP", false, "turn on trace profiling")
	)
	flag.Parse()

	// if *lowName == "" || *highName == "" || *highPLYName == "" {
	// 	log.Fatal("One of file needed for baking was not provided")
	// }

	//Profiling
	if *cpuProfiling {
		defer profile.Start(profile.CPUProfile, profile.ProfilePath(".")).Stop()
	}
	if *memProfiling {
		defer profile.Start(profile.MemProfile, profile.ProfilePath(".")).Stop()
	}
	if *tracecProfiling {
		defer profile.Start(profile.TraceProfile, profile.ProfilePath(".")).Stop()
	}
	scene := gobaker.NewScene(*size)
	log.Printf("Starting")
	start := time.Now()
	err := scene.Lowpoly.ReadOBJ(*lowName, false)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("Read lowpoly mesh in %s", time.Since(start))

	err = scene.Highpoly.ReadOBJ(*highName, true)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("Read highpoly mesh.. %s", time.Since(start))
	err = scene.Highpoly.ReadPLY(*highPLYName)
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("Readed PLY.. %s", time.Since(start))
	log.Printf("Started baking in %dx%d resolution", *size, *size)
	scene.Bake()
	//log.Printf("Baking...")
	//log.Printf("Finished baking: %s", time.Since(start))
	scene.BakedDiffuse.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_diff.png")
	scene.BakedID.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_id.png")
	scene.BakedNormal.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_nrm.png")
	scene.BakedObjectNormal.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_obj_nrm.png")
	log.Printf("Program finished in: %s", time.Since(start))
}
