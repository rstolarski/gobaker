package gobaker

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"os"

	"github.com/disintegration/imaging"
)

// Texture defines image texture
type Texture struct {
	Image *image.NRGBA
	w, h  int
}

// Material of the mesh
type Material struct {
	Diffuse *Texture
	ID      *Texture
}

// NewTexture creates new texture with set size
func NewTexture(size int) *Texture {
	t := &Texture{
		image.NewNRGBA(image.Rect(0, 0, size, size)),
		size,
		size,
	}
	return t
}

// SaveImage saves Texture's image with a given name 'n'
func (t *Texture) SaveImage(n string) {
	img := imaging.FlipV(t.Image)
	outDiff, err := os.Create("./" + n)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch n[len(n)-3:] {
	case "png":
		err = png.Encode(outDiff, img)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "jpg":
		err = jpeg.Encode(outDiff, img, &jpeg.Options{Quality: 80})
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
}

// LoadTexture loads texture object from png or jpg file with name 'n'
func LoadTexture(n string) *Texture {
	var img image.Image

	// Read image from file that already exists
	f, err := os.Open(n)
	if err != nil {
		return &Texture{nil, 0, 0}
	}
	defer f.Close()

	switch n[len(n)-3:] {
	case "png":
		img, err = png.Decode(f)
		if err != nil {
			return nil
		}
	case "jpg":
		img, err = jpeg.Decode(f)
		if err != nil {
			return nil
		}
	}

	// Converting img to NRGBA format
	out := image.NewNRGBA(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			out.Set(x, y, color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA))
		}
	}

	return &Texture{
		imaging.FlipV(out),
		out.Bounds().Max.X,
		out.Bounds().Max.Y,
	}
}

// SamplePixel return color of a pixel in u and v coordinates on image
func (t *Texture) SamplePixel(u, v float64) color.NRGBA {
	indX := math.Mod(u*float64(t.w-1), float64(t.w))
	indY := math.Mod(v*float64(t.h-1), float64(t.h))
	if indX < 0 {
		indX = float64(t.w) + indX
	}
	if indY < 0 {
		indY = float64(t.h) + indY
	}

	return t.Image.NRGBAAt(int(indX), int(indY))
}
