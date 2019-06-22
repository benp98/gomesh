package gomesh

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
