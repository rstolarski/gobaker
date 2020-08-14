package main

import (
	"log"
	"strings"
	"time"

	"github.com/rtropisz/gobaker"
)

const (
	size        = 512
	lowName     = "./AppleTree_lowpoly.obj"
	highName    = "./AppleTree.obj"
	highPlYName = "./AppleTree.ply"
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
	scene.Bake()
	scene.BakedDiffuse.SaveImage(strings.TrimSuffix(lowName, ".obj") + "_diff.png")
	scene.BakedID.SaveImage(strings.TrimSuffix(lowName, ".obj") + "_id.png")
	log.Printf("Program finished in: %s", time.Since(start))
}
