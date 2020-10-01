package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/profile"
	"github.com/rtropisz/gobaker/gobaker"

	bep "github.com/gen2brain/beeep"
)

var (
	size               = flag.Int("s", 1024, "size of the output images in pixels")
	lowName            = flag.String("lp", "", "path to lowpoly mesh")
	highName           = flag.String("hp", "", "path to highpoly mesh")
	highPLYName        = flag.String("ply", "", "path to highpoly PLY mesh")
	readID             = flag.Bool("id", true, "read ID map and use it in baking process")
	maxRearDistance    = flag.Float64("rearD", 3.0, "max rear distance")
	maxFrontalDistance = flag.Float64("frontD", 3.0, "max front distance")
	output             = flag.String("o", "", "output directory")

	cpuProfiling    = flag.Bool("cpuP", false, "turn on cpu profiling")
	memProfiling    = flag.Bool("memP", false, "turn on memory profiling")
	tracecProfiling = flag.Bool("traceP", false, "turn on trace profiling")
	useHalfCPU      = flag.Bool("useHalfCPU", false, "use half of available CPU cores, if set to false all you have")
)

func main() {

	flag.Parse()

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

	workers := runtime.NumCPU()
	if *useHalfCPU {
		workers = runtime.NumCPU() / 2
	}

	scene := gobaker.NewScene(*size, *readID, *maxFrontalDistance, *maxRearDistance)

	log.Printf("Starting")
	start := time.Now()

	err := scene.Lowpoly.ReadOBJ(*lowName)
	if err != nil {
		log.Fatal(err)
	}

	err = scene.Highpoly.ReadOBJ(*highName)
	if err != nil {
		log.Fatal(err)
	}

	if scene.ReadID {
		err = scene.Highpoly.ReadPLY(*highPLYName)
		if err != nil {
			log.Fatal(err)
		}
	}

	err = scene.ReadTexturesForHighpoly()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Started baking in %dx%d resolution", *size, *size)
	scene.Bake(workers)

	err = scene.BakedDiffuse.SaveImage(*output, strings.TrimSuffix(*lowName, ".obj")+"_diff.png")
	if err != nil {
		log.Fatal(err)
	}

	if scene.ReadID {
		err = scene.BakedID.SaveImage(*output, strings.TrimSuffix(*lowName, ".obj")+"_id.png")
		if err != nil {
			log.Fatal(err)
		}
	}

	// scene.BakedNormal.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_nrm.png")
	// scene.BakedObjectNormal.SaveImage(strings.TrimSuffix(*lowName, ".obj") + "_obj_nrm.png")
	log.Printf("Program finished in: %s", time.Since(start))
	fmt.Print("Press 'Enter' to continue...")

	bep.Alert("BAKER", "Baking finished!", "cooper.png")

	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
