package gomesh

// AddPlane generates a plane with the size of {w,h} with offset z. The function panics if the axis parameter is invalid.
func (mesh *Mesh) AddPlane(w, h, offset float64, axis int) {
	vertices := make([]int, 0)

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
