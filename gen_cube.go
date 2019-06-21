package gomesh

// AddCube generates a cube with the size of {sx, sy, sz}
func (mesh *Mesh) AddCube(sx, sy, sz float64) {
	mesh.AddPlane(sx, sy, sz/2, Z)
	mesh.AddPlane(sx, sy, -sz/2, Z)
	mesh.AddPlane(sz, sy, sx/2, X)
	mesh.AddPlane(sz, sy, -sx/2, X)
	mesh.AddPlane(sx, sz, sy/2, Y)
	mesh.AddPlane(sx, sz, -sy/2, Y)
}
