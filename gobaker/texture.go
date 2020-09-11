package gobaker

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math"
	"os"
	"path/filepath"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/ftrvxmtrx/tga"
)

// Texture defines image texture
type Texture struct {
	Image *image.NRGBA
	w, h  int
}

// Material of the mesh
type Material struct {
	Diffuse *Texture
	Normal  *Texture
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
func (t *Texture) SaveImage(dir, f string) {
	defer duration(track("Saving file " + f + "took"))

	a := strings.Split(toSlash(f), "/")
	f = a[len(a)-1]

	img := imaging.FlipV(t.Image)
	outDiff, err := os.Create(filepath.Join(dir, f))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	switch f[len(f)-3:] {
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
	case "tga":
		err = tga.Encode(outDiff, img)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

}

// LoadTexture loads texture object from png or jpg file with name 'n'
func LoadTexture(pathToFile string) (*Texture, error) {
	if pathToFile == "" {
		return nil, fmt.Errorf("Cannot open file. Path is not set")
	}
	var img image.Image

	// Read image from file that already exists
	f, err := os.Open(pathToFile)
	if err != nil {
		return &Texture{nil, 0, 0}, nil
	}
	defer f.Close()

	switch pathToFile[len(pathToFile)-3:] {
	case "png":
		img, err = png.Decode(f)
		if err != nil {
			return nil, err
		}
	case "jpg":
		img, err = jpeg.Decode(f)
		if err != nil {
			return nil, err
		}
	case "tga":
		img, err = tga.Decode(f)
		if err != nil {
			return nil, err
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
		},
		nil
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
