package main

import (
	"fmt"
	"log"
	"os"

	"github.com/benp98/gomesh"
)

func main() {
	// Open the file "mesh.obj" or throw an error
	file, err := os.Open("mesh.obj")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Decode the mesh data
	mesh, err := gomesh.DecodeOBJ(file)
	if err != nil {
		log.Fatal(err)
	}

	// Print all vertices
	fmt.Println("Vertices:")
	for i, v := range mesh.Vertices {
		fmt.Printf("\tVertex: %4d %8.2f %8.2f %8.2f\n", i, v.X, v.Y, v.Z)
	}
	fmt.Println()

	// Print all faces
	fmt.Println("Faces:")
	for _, f := range mesh.Faces {
		fmt.Println("\tFace:")

		// Print vertex infos
		for _, vID := range f.Vertices {
			v := mesh.Vertices[vID]
			fmt.Printf("\t\tVertex: %8.2f %8.2f %8.2f\n", v.X, v.Y, v.Z)
		}
	}
}
