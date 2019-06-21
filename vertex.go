package gomesh

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
	if err == ErrVertexNotFound {
		id = len(mesh.Vertices)
		mesh.Vertices = append(mesh.Vertices, Point{x, y, z})
	}

	return id
}
