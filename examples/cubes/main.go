package main

import (
	"fmt"
	"log"
	"os"

	"github.com/benp98/gomesh"
	"github.com/benp98/gomesh/meshutils"
	"github.com/benp98/gomesh/obj"
)

func main() {
	// Create new file
	f, err := os.Create("cubes.obj")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scene := make(gomesh.Scene, 0)

	// Create Mesh, generate Cube and Encode as OBJ
	i := 1
	for x := -1; x <= 1; x++ {
		for y := -1; y <= 1; y++ {
			for z := -1; z <= 1; z++ {
				mesh := &gomesh.Mesh{Name: fmt.Sprintf("Cube-%d", i)}
				meshutils.AddCube(mesh, gomesh.Vector3D{float64(x) * 2, float64(y) * 2, float64(z) * 2}, 0.5)
				scene = append(scene, mesh)
				i++
			}
		}
	}

	obj.Encode(f, scene)
}
