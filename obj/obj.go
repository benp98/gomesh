// Package obj implements a very basic OBJ importer and exporter for gomesh
package obj

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/benp98/gomesh"
)

// ErrUnsupportedOBJ is returned when DecodeOBJ can not parse the data
var ErrUnsupportedOBJ = errors.New("unsupported OBJ Format")

// Define some regular expressions to match the commands
var objVertexExp = regexp.MustCompile(`^v (-?[0-9]+?\.[0-9]+?) (-?[0-9]+?\.[0-9]+?) (-?[0-9]+?\.[0-9]+?)$`)
var objFaceExp = regexp.MustCompile(`^f ([\-.0-9 ]+?)$`)

// Decode reads a Mesh from the given Reader of a Wavefront OBJ file
func Decode(reader io.Reader) (gomesh.Scene, error) {
	mesh := &gomesh.Mesh{}

	// Process the reader line by line
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()

		// Check the line against a set of regular expressions
		switch true {
		case objVertexExp.MatchString(line): // Line is a vertex
			vertexDataRaw := objVertexExp.FindStringSubmatch(line)[1:]
			if len(vertexDataRaw) != 3 {
				return nil, ErrUnsupportedOBJ
			}

			// Parse X
			x, err := strconv.ParseFloat(vertexDataRaw[0], 64)
			if err != nil {
				return nil, err
			}

			// Parse Y
			y, err := strconv.ParseFloat(vertexDataRaw[1], 64)
			if err != nil {
				return nil, err
			}

			// Parse Z
			z, err := strconv.ParseFloat(vertexDataRaw[2], 64)
			if err != nil {
				return nil, err
			}

			// Create Vertex
			vertex := &gomesh.Vertex{}
			vertex.Position = gomesh.Vector3D{x, y, z}

			// Add vertex data
			mesh.Vertices = append(mesh.Vertices, vertex)

		case objFaceExp.MatchString(line): // Line is a face
			faceDataRaw := objFaceExp.FindStringSubmatch(line)[1]
			faceData := strings.Fields(faceDataRaw)

			face := &gomesh.Face{}

			// Process Vertex ids
			for _, vStr := range faceData {
				vInt, err := strconv.Atoi(vStr)
				if err != nil {
					return nil, ErrUnsupportedOBJ
				}

				face.VertexIDs = append(face.VertexIDs, vInt-1) // OBJ Vertex counting starts with 1
			}

			mesh.Faces = append(mesh.Faces, face)
		default: // Something different, just ignore
		}
	}

	mesh.Name = "ImportMesh"
	scene := gomesh.Scene{mesh}

	// Validate the scene, return an error if invalid
	if !mesh.Validate() {
		return nil, gomesh.ErrInvalidScene
	}
	return scene, nil
}

// Encode writes the mesh to the writer in Wavefront OBJ format
func Encode(writer io.Writer, scene gomesh.Scene) error {
	fmt.Fprintln(writer, "# gomesh OBJ Encoder")
	fmt.Fprintln(writer, "# https://github.com/benp98/gomesh")

	vertexIDoffset := 0
	for _, mesh := range scene {
		// Write mesh name if set
		if mesh.Name != "" {
			fmt.Fprintf(writer, "o %s\n", mesh.Name)
		} else {
			fmt.Fprintln(writer, "o Gomesh")
		}

		// Write all vertices, line by line
		for _, vertex := range mesh.Vertices {
			fmt.Fprintf(writer, "v %f %f %f\n", vertex.Position[0], vertex.Position[1], vertex.Position[2])
		}

		// Create faces by connecting vertices
		for _, face := range mesh.Faces {
			if face.SmoothingGroup == gomesh.NoSmoothing {
				fmt.Fprintln(writer, "s off")
			} else {
				fmt.Fprintf(writer, "s %d\n", face.SmoothingGroup)
			}

			fmt.Fprint(writer, "f")
			for _, vertex := range face.VertexIDs {
				fmt.Fprintf(writer, " %d", vertex+vertexIDoffset+1) // OBJ Vertex counting starts with 1
			}
			fmt.Fprintln(writer)
		}

		vertexIDoffset += len(mesh.Vertices)
	}

	return nil
}
