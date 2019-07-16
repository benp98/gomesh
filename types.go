// Package gomesh provides functions for basic handling of 3D mesh data in Go. The mesh data can also be read from and written to OBJ files.
package gomesh

import "errors"

// ErrVertexNotFound is returned when the FindVertex() function does not find the vertex
var ErrVertexNotFound = errors.New("Vertex not found")

// ErrUnsupportedOBJ is returned when DecodeOBJ can not parse the data
var ErrUnsupportedOBJ = errors.New("Unsupported OBJ Format")

// ErrInvalidMesh is returned when a decoding function can not make a valid mesh from the data
var ErrInvalidMesh = errors.New("Invalid Mesh data")

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
