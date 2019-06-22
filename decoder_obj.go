package gomesh

import (
	"bufio"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// Define some regular expressions to match the commands
var objVertexExp = regexp.MustCompile(`^v (-?[0-9]+?\.[0-9]+?) (-?[0-9]+?\.[0-9]+?) (-?[0-9]+?\.[0-9]+?)$`)
var objFaceExp = regexp.MustCompile(`^f ([\-.0-9 ]+?)$`)

// DecodeOBJ reads a Mesh from the given Reader of a Wavefront OBJ file
func DecodeOBJ(reader io.Reader) (*Mesh, error) {
	mesh := new(Mesh)

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
			vertex := new(Point)
			vertex.X = x
			vertex.Y = y
			vertex.Z = z

			// Add vertex data
			mesh.Vertices = append(mesh.Vertices, *vertex)

		case objFaceExp.MatchString(line): // Line is a face
			faceDataRaw := objFaceExp.FindStringSubmatch(line)[1]
			faceData := strings.Fields(faceDataRaw)

			face := new(Face)

			// Process Vertex ids
			for _, vStr := range faceData {
				vInt, err := strconv.Atoi(vStr)
				if err != nil {
					return nil, ErrUnsupportedOBJ
				}

				face.Vertices = append(face.Vertices, vInt-1) // OBJ Vertex counting starts with 1
			}

			mesh.Faces = append(mesh.Faces, *face)
		default: // Something different, just ignore
		}
	}

	// Validate the mesh, return an error if invalid
	if !mesh.Validate() {
		return nil, ErrInvalidMesh
	}

	return mesh, nil
}
