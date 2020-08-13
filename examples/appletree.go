package main

import (
	"log"
	"time"

	"github.com/rtropisz/gobaker"
)

const (
	s = 1024
)

func main() {
	scene := gobaker.NewScene(s)
	log.Printf("Starting")
	start := time.Now()
	log.Printf("Reading lowpoly mesh.. %s", time.Since(start))
	scene.Lowpoly.ReadOBJ("examples/AppleTree_lowpoly", false)
	log.Printf("Reading highpoly mesh.. %s", time.Since(start))
	scene.Highpoly.ReadOBJ("examples/AppleTree", true)
	log.Printf("Baking")
	scene.Bake()
	log.Printf("Finished baking took %s", time.Since(start))
	scene.BakedDiffuse.SaveImage("outAppleTree_diff.png")
	scene.BakedNormal.SaveImage("outAppleTree_norm.png")
	log.Printf("Finished saving images took %s", time.Since(start))
}
