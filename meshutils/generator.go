// Package allows to manipulate mesh data for gomesh
package meshutils

import "github.com/benp98/gomesh"

const (
	AxisX = iota
	AxisY
	AxisZ
)

// AddPlane adds a plane to the mesh with the given offset, size and orientation
func AddPlane(mesh *gomesh.Mesh, offset gomesh.Vector3D, size, axis float64, orientation bool) {
	vertices := make([]int, 0)

	// Add vertices
	for x := -1; x <= 1; x += 2 {
		for y := -1; y <= 1; y += 2 {
			switch axis {
			case AxisX:
				vertices = append(vertices, mesh.AddVertex(&gomesh.Vertex{
					Position: gomesh.Vector3D{0, size * float64(x), size * float64(y)}.AddVector3D(offset),
				}))
			case AxisY:
				vertices = append(vertices, mesh.AddVertex(&gomesh.Vertex{
					Position: gomesh.Vector3D{size * float64(x), 0, size * float64(y)}.AddVector3D(offset),
				}))
			case AxisZ:
				vertices = append(vertices, mesh.AddVertex(&gomesh.Vertex{
					Position: gomesh.Vector3D{size * float64(x), size * float64(y), 0}.AddVector3D(offset),
				}))
			default:
				panic("Invalid axis specification")
			}
		}
	}

	// Create face
	if orientation {
		mesh.Faces = append(mesh.Faces, &gomesh.Face{VertexIDs: []int{vertices[2], vertices[3], vertices[1], vertices[0]}})
	} else {
		mesh.Faces = append(mesh.Faces, &gomesh.Face{VertexIDs: []int{vertices[0], vertices[1], vertices[3], vertices[2]}})
	}
}

// AddCube adds a cube to the mesh with the given offset and size
func AddCube(mesh *gomesh.Mesh, offset gomesh.Vector3D, size float64) {
	AddPlane(mesh, offset.AddVector3D(gomesh.Vector3D{0, 0, size}), size, AxisZ, true)
	AddPlane(mesh, offset.AddVector3D(gomesh.Vector3D{0, 0, -size}), size, AxisZ, false)
	AddPlane(mesh, offset.AddVector3D(gomesh.Vector3D{size, 0, 0}), size, AxisX, true)
	AddPlane(mesh, offset.AddVector3D(gomesh.Vector3D{-size, 0, 0}), size, AxisX, false)
	AddPlane(mesh, offset.AddVector3D(gomesh.Vector3D{0, size, 0}), size, AxisY, false)
	AddPlane(mesh, offset.AddVector3D(gomesh.Vector3D{0, -size, 0}), size, AxisY, true)
}
