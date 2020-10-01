package gobaker

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"io"
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
	Name    string
	Diffuse *Texture
	//Normal  *Texture
	ID *Texture
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
func (t *Texture) SaveImage(dir, f string) error {
	a := strings.Split(toSlash(f), "/")
	f = a[len(a)-1]
	fpath := filepath.Join(dir, f)

	var outDiff *os.File

	outDiff, err := os.Create(fpath)
	if err != nil {
		return err
	}

	if dir == "" {
		defer duration(track("Saving file .\\" + fpath + " took"))
	} else {
		defer duration(track("Saving file " + fpath + " took"))
	}

	defer outDiff.Close()

	img := imaging.FlipV(t.Image)

	switch f[len(f)-3:] {
	case "png":
		err = png.Encode(outDiff, img)
		if err != nil {
			return err
		}
		break
	case "jpg":
		err = jpeg.Encode(outDiff, img, &jpeg.Options{Quality: 80})
		if err != nil {
			return err
		}
		break
	case "tga":
		err = tga.Encode(outDiff, img)
		if err != nil {
			return err
		}
		break
	}
	return err
}

// LoadTexture loads texture object from png or jpg file with name 'n'
func LoadTexture(pathToFile string) (*Texture, error) {
	if pathToFile == "" {
		return nil, fmt.Errorf("Cannot open file. Path is not set")
	}

	fpath, err := filepath.Abs(pathToFile)
	if err != nil {
		return nil, fmt.Errorf("Cannot open file. %v", err)
	}

	var img *image.NRGBA

	if _, err := os.Stat(fpath + ".png"); err == nil {
		defer duration(track("Reading " + fpath + ".PNG took"))
		if img, err = decodeTexture(fpath, ".png", png.Decode); err != nil {
			return nil, err
		}
		return &Texture{img, img.Bounds().Max.X, img.Bounds().Max.Y}, nil
	}

	// Not working, some textures are read incorrectly
	//
	// if _, err := os.Stat(fpath + ".tga"); err == nil {
	// 	log.Println("Reading " + fpath + ".TGA")
	// 	if img, err = decodeTexture(fpath, ".tga", tga.Decode); err != nil {
	// 		return nil, err
	// 	}
	// 	return &Texture{img, img.Bounds().Max.X, img.Bounds().Max.Y}, nil
	// }

	if _, err := os.Stat(fpath + ".jpg"); err == nil {
		defer duration(track("Reading " + fpath + ".JPG took"))
		if img, err = decodeTexture(fpath, ".jpg", jpeg.Decode); err != nil {
			return nil, err
		}
		return &Texture{img, img.Bounds().Max.X, img.Bounds().Max.Y}, nil
	}

	return nil, err
}

type decode func(io.Reader) (image.Image, error)

func decodeTexture(path, ext string, fn decode) (out *image.NRGBA, err error) {
	path = path + ext
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("Cannot open file. %v", err)
	}

	defer f.Close()

	var img image.Image

	img, err = fn(f)
	if err != nil {
		return nil, fmt.Errorf("Cannot decode ."+ext+" file. %v", err)
	}

	out = image.NewNRGBA(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			out.Set(x, y, color.NRGBAModel.Convert(img.At(x, y)).(color.NRGBA))
		}
	}
	return imaging.FlipV(out), nil
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
