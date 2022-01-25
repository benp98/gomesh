// Package gomesh provides functions for basic handling of 3D mesh data in Go.
package gomesh

import "errors"

// ErrVertexNotFound is returned when the FindVertex() function does not find the vertex
var ErrVertexNotFound = errors.New("vertex not found")

// ErrInvalidScene is returned when a decoding function can not make a valid scene from the source data
var ErrInvalidScene = errors.New("invalid Scene data")

// Special smoothing group which disables smoothing
const NoSmoothing = 0

// Vertex holds the position and normal information of a single vertex
type Vertex struct {
	Position Vector3D
	Normal   Vector3D
}

// Face contains a list of points (as a slice of mesh vertex ids) which describe the face's 3D structure
type Face struct {
	SmoothingGroup int
	VertexIDs      []int
}

// Mesh is composed of one or multiple faces
type Mesh struct {
	Name     string
	Vertices []*Vertex
	Faces    []*Face
}

// FindVertex searches for a vertex at the given position
func (mesh Mesh) FindVertex(other *Vertex, tolerance float64) (int, error) {
	for k, v := range mesh.Vertices {
		if v.Position.AlmostEqual(other.Position, tolerance) && v.Normal.AlmostEqual(other.Normal, tolerance) {
			return k, nil
		}
	}

	return 0, ErrVertexNotFound
}

// AddVertex adds a vertex to the list of vertices and returns the index
func (mesh *Mesh) AddVertex(v *Vertex) int {
	id := len(mesh.Vertices)

	mesh.Vertices = append(mesh.Vertices, v)

	return id
}

// Validate performs basic validation of the Mesh data
func (mesh Mesh) Validate() bool {
	// Iterate over all faces
	for _, f := range mesh.Faces {
		// Check if all vertex ids are in bounds
		for _, v := range f.VertexIDs {
			if v > len(mesh.Vertices)-1 || v < 0 {
				// Vertex ids are not in bounds
				return false
			}
		}
	}

	// Everything okay
	return true
}

// Scene is a collection of meshes
type Scene []*Mesh

// Validate performs basic validation of the Scene data and its meshes
func (scene Scene) Validate() bool {
	meshNames := make(map[string]bool)

	for _, mesh := range scene {
		// There must be only one mesh with a given name
		if _, exists := meshNames[mesh.Name]; exists {
			return false
		}
		meshNames[mesh.Name] = true

		if !mesh.Validate() {
			return false
		}
	}

	// Everything ok
	return true
}
