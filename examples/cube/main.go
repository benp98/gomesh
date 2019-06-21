package main

import (
	"log"
	"os"

	"github.com/benp98/gomesh"
)

func main() {
	// Create new file
	f, err := os.Create("cube.obj")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	// Create Mesh, generate Cube and Encode as OBJ
	mesh := gomesh.Mesh{}
	mesh.AddCube(1.0, 1.0, 1.0)
	mesh.EncodeOBJ(f)
}
