package main

import (
	"log"
	"os"

	"github.com/benp98/gomesh"
	"github.com/benp98/gomesh/meshutils"
	"github.com/benp98/gomesh/obj"
)

func main() {
	// Create new file
	f, err := os.Create("cube.obj")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Create Mesh, generate Cube and Encode as OBJ
	mesh := &gomesh.Mesh{Name: "Cube"}
	meshutils.AddCube(mesh, gomesh.Vector3D{0, 0, 0}, 1)
	scene := gomesh.Scene{mesh}
	obj.Encode(f, scene)
}
