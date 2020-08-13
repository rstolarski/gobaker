package main

import (
	"log"
	"strings"
	"time"

	"github.com/rtropisz/gobaker"
)

const (
	size     = 1024
	lowName  = "./AppleTree_lowpoly.obj"
	highName = "./AppleTree.obj"
)

func main() {
	scene := gobaker.NewScene(size)
	log.Printf("Starting")
	start := time.Now()
	scene.Lowpoly.ReadOBJ(lowName, false)
	log.Printf("Reading lowpoly mesh.. %s", time.Since(start))
	scene.Highpoly.ReadOBJ(highName, true)
	log.Printf("Reading highpoly mesh.. %s", time.Since(start))
	scene.Bake()
	log.Printf("Baking...")
	log.Printf("Finished baking took %s", time.Since(start))
	scene.BakedDiffuse.SaveImage(strings.TrimSuffix(lowName, ".obj") + "_diff.png")
	scene.BakedNormal.SaveImage(strings.TrimSuffix(lowName, ".obj") + "_norm.png")
	log.Printf("Finished saving images took %s", time.Since(start))
}
