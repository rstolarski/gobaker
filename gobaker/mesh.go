package gobaker

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

// Mesh describes a 3D object
type Mesh struct {
	Triangles []Triangle
	Materials []Material
}

// ReadPLY adds to Mesh object vertex color values read from the PLY file, based on pathToFile
func (m *Mesh) ReadPLY(pathToFile string) error {
	if pathToFile == "" {
		return fmt.Errorf("Cannot open file. Path is not set")
	}

	inFile, err := os.Open(pathToFile)
	if err != nil {
		return fmt.Errorf(
			"Cannot open file. %v",
			err,
		)
	}
	defer inFile.Close()
	defer duration(track("Reading " + pathToFile + " took"))

	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	read := false

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		if fields[0] == "end_header" {
			read = true
		}
		if read == false || len(fields) < 5 { // use only lines with vertex definition
			continue
		}

		// find vertices based on vertex position
		vPos, err := ParseVector(fields[0], fields[2], fields[1])
		if err != nil {
			return fmt.Errorf(
				"string cannot be converted to float64, %v",
				err,
			)
		}
		vAlpha, err := strconv.ParseFloat(fields[len(fields)-1], 64)
		if err != nil {
			return fmt.Errorf(
				"string cannot be converted to float64, %v",
				err,
			)
		}
		vAlpha /= 255.0

		for i := 0; i < len(m.Triangles); i++ {
			if m.Triangles[i].V0.v.CompareVectors(vPos, 0.0000001) {
				m.Triangles[i].V0.SetVertexAlpha(vAlpha)

			}
			if m.Triangles[i].V1.v.CompareVectors(vPos, 0.0000001) {
				m.Triangles[i].V1.SetVertexAlpha(vAlpha)

			}
			if m.Triangles[i].V2.v.CompareVectors(vPos, 0.0000001) {
				m.Triangles[i].V2.SetVertexAlpha(vAlpha)

			}
		}
	}
	return nil
}

// ReadOBJ return Mesh object read from the OBJ file, based on pathToFile
// It needs a Material slice in order to add
// material to each triangle in the mesh
func (m *Mesh) ReadOBJ(pathToFile string, readMaterials bool) error {
	if pathToFile == "" {
		return fmt.Errorf("Cannot open file. Path is not set")
	}

	inFile, err := os.Open(pathToFile)
	if err != nil {
		return fmt.Errorf(
			"Cannot open file. %v",
			err,
		)
	}

	defer duration(track("Reading " + pathToFile + " took"))

	vertices := make([]Vector, 0)
	normals := make([]Vector, 0)
	textures := make([]Vector, 0)

	reader := bufio.NewReader(inFile)
	var line string

	for {
		line, err = reader.ReadString('\n')

		if err != nil {
			break
		}
		fields := strings.Fields(line)
		if len(fields) < 2 {
			continue
		}
		key := fields[0]
		args := fields[1:]

		switch key {
		case "v":
			v, err := ParseVector(args[0], args[1], args[2])
			if err != nil {
				return fmt.Errorf(
					"string %v cannot be converted to float64, %v",
					args,
					err,
				)
			}
			vertices = append(vertices, v)
		case "vn":
			vn, err := ParseVector(args[0], args[1], args[2])
			if err != nil {
				return fmt.Errorf(
					"string %v cannot be converted to float64, %v",
					args,
					err,
				)
			}
			normals = append(normals, vn.Normalize())
		case "vt":
			vt, err := ParseVector(args[0], args[1], "1.0")
			if err != nil {
				return fmt.Errorf(
					"string %v cannot be converted to float64, %v",
					args,
					err,
				)
			}
			textures = append(textures, vt)
		case "usemtl":
			if !readMaterials {
				break
			}
			// Get material name without prefix
			matName := strings.TrimPrefix(args[0], "MI_")
			f := toSlash(pathToFile)      // Convert pathToFile to proper slashes
			fSep := strings.Split(f, "/") // Split pathToFile to chunks
			fSep = fSep[:len(fSep)-1]     // Delete file name from slice

			f = strings.Join(fSep, "/")          // Join remaining element into path
			matName = path.Join(f, "T_"+matName) // Add directory path to material name with prefix 'T_'

			//var texDiffuse, texNormal, texID Texture
			var mat Material
			mat.Diffuse, err = LoadTexture(matName + "_diff.png")
			if err != nil {
				return nil
			}
			mat.Normal, err = LoadTexture(matName + "_nrm.png")
			if err != nil {
				return nil
			}
			mat.ID, err = LoadTexture(matName + "_id.png")
			if err != nil {
				return nil
			}

			m.Materials = append(m.Materials, mat)
		case "f":
			size := len(args)
			points := make([]Vertex, 3)

			for i := 0; i < size; i++ {
				f := strings.Split(args[i], "/")
				var v Vector
				var n Vector
				var t Vector

				if len(fields) > 0 {
					vReadIndes, err := index(f[0], len(vertices))
					if err != nil {
						return err
					}
					v = vertices[vReadIndes]

					tReadIndes, err := index(f[1], len(textures))
					if err != nil {
						return err

					}
					t = textures[tReadIndes]

					nReadIndes, err := index(f[2], len(normals))
					if err != nil {
						return err

					}
					n = normals[nReadIndes]
				}
				points[i] = Vertex{v, t, n, 0}
			}
			if !readMaterials {
				m.addTriangle(points[0], points[1], points[2], nil)
			} else {
				m.addTriangle(points[0], points[1], points[2], &m.Materials[len(m.Materials)-1])
			}
		}
	}
	return nil
}

func toSlash(pathToFile string) string {
	Separator := os.PathSeparator
	if Separator == '/' {
		return pathToFile

	}
	return strings.ReplaceAll(pathToFile, string(Separator), "/")

}

// String implements Stringer interface.
// It displays information about each triangle in the mesh
func (m Mesh) String() string {
	var s string
	s += "Mesh: \n"

	for i := 0; i < len(m.Triangles); i++ {
		s += fmt.Sprintf("%v\n", m.Triangles[i])
	}
	return s
}

// addTriangle adds new triangle to a mesh
func (m *Mesh) addTriangle(v0, v1, v2 Vertex, material *Material) {
	triangle := Triangle{
		V0:       v0,
		V1:       v1,
		V2:       v2,
		Material: material,
	}
	m.Triangles = append(m.Triangles, triangle)
}

// index returns index value of face's vertex/texture/normal coordinates, based on string
func index(s string, size int) (n int, err error) {
	i, err := strconv.ParseInt(s, 0, 0)
	if err != nil {
		return 0, err
	}
	if n < 0 || n > size-1 {
		return 0, fmt.Errorf("ReadIndes out of bounds: %v (size: %v, string: %v)", n, size, s)
	}
	n = size + int(i)
	if i > 0 {
		n = int(i - 1)
	}

	return n, nil
}
