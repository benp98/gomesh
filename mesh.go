// Package gomesh provides functions for basic handling of 3D mesh data in Go.
package gomesh

import "errors"

// ErrVertexNotFound is returned when the FindVertex() function does not find the vertex
var ErrVertexNotFound = errors.New("vertex not found")

// ErrInvalidMesh is returned when a decoding function can not make a valid mesh from the data
var ErrInvalidMesh = errors.New("invalid Mesh data")

// Axis definitions
const (
	X = iota
	Y
	Z
)

// Point is a location definition in 3D space
type Point struct {
	X float64
	Y float64
	Z float64
}

// Face contains a list of points (as a slice of mesh vertex ids) which describe the face's 3D structure
type Face struct {
	Vertices []int
}

// Mesh is composed of one or multiple faces
type Mesh struct {
	Vertices []Point
	Faces    []Face
}

// FindVertex searches for a vertex at the given position
func (mesh Mesh) FindVertex(x, y, z float64) (int, error) {
	for k, v := range mesh.Vertices {
		if v.X == x && v.Y == y && v.Z == z {
			return k, nil
		}
	}

	return 0, ErrVertexNotFound
}

// AddVertex creates a vertex at the given position if there is not already a vertex. Returns the vertex id
func (mesh *Mesh) AddVertex(x, y, z float64) int {
	id, err := mesh.FindVertex(x, y, z)
	if errors.Is(err, ErrVertexNotFound) {
		id = len(mesh.Vertices)
		mesh.Vertices = append(mesh.Vertices, Point{x, y, z})
	}

	return id
}

// Validate validates the Mesh data
func (mesh Mesh) Validate() bool {
	// Iterate over all faces
	for _, f := range mesh.Faces {
		// Check if all vertex ids are in bounds
		for _, v := range f.Vertices {
			if v > len(mesh.Vertices)-1 || v < 0 {
				// Vertex ids are not in bounds
				return false
			}
		}
	}

	// Everything okay
	return true
}

// AddPlane generates a plane with the size of {w,h} with offset z. The function panics if the axis parameter is invalid.
func (mesh *Mesh) AddPlane(w, h, offset float64, axis int) {
	var vertices []int

	// Add vertices
	for x := -w / 2; x <= w/2; x += w {
		for y := -h / 2; y <= h/2; y += h {
			switch axis {
			case X:
				vertices = append(vertices, mesh.AddVertex(offset, y, x))
			case Y:
				vertices = append(vertices, mesh.AddVertex(x, offset, y))
			case Z:
				vertices = append(vertices, mesh.AddVertex(x, y, offset))
			default:
				panic("Invalid axis specification")
			}
		}
	}

	// Make faces
	mesh.Faces = append(mesh.Faces, Face{[]int{vertices[0], vertices[1], vertices[3], vertices[2]}})
}

// AddCube generates a cube with the size of {sx, sy, sz}
func (mesh *Mesh) AddCube(sx, sy, sz float64) {
	mesh.AddPlane(sx, sy, sz/2, Z)
	mesh.AddPlane(sx, sy, -sz/2, Z)
	mesh.AddPlane(sz, sy, sx/2, X)
	mesh.AddPlane(sz, sy, -sx/2, X)
	mesh.AddPlane(sx, sz, sy/2, Y)
	mesh.AddPlane(sx, sz, -sy/2, Y)
}
