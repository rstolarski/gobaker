package main

import (
	"log"
	"runtime"
	"strings"
	"time"

	"github.com/rtropisz/gobaker/gobaker"
)

const (
	size        = 512
	lowName     = "./AppleTree_lowpoly.obj"
	highName    = "./AppleTree.obj"
	highPlYName = "./AppleTree.ply"
	output      = "./"
)

func main() {
	scene := gobaker.NewScene(size)
	log.Printf("Starting")
	start := time.Now()

	err := scene.Lowpoly.ReadOBJ(lowName, false)
	if err != nil {
		log.Fatal(err)
	}

	err = scene.Highpoly.ReadOBJ(highName, true)
	if err != nil {
		log.Fatal(err)
	}

	err = scene.Highpoly.ReadPLY(highPlYName)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Started baking in %dx%d resolution", size, size)
	scene.Bake(runtime.NumCPU())

	err = scene.BakedDiffuse.SaveImage(output, strings.TrimSuffix(lowName, ".obj")+"_diff.png")
	if err != nil {
		log.Fatal(err)
	}

	err = scene.BakedID.SaveImage(output, strings.TrimSuffix(lowName, ".obj")+"_id.png")
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Program finished in: %s", time.Since(start))
}
