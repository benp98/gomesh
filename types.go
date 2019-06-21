package gomesh

import "errors"

// ErrVertexNotFound is returned when the FindVertex() function does not find the vertex
var ErrVertexNotFound = errors.New("Vertex not found")

// Axis definitions
const (
	X = iota
	Y = iota
	Z = iota
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
