package main

import (
	"log"
	"os"

	"github.com/benp98/gomesh"
)

func main() {
	// Open the file "mesh.obj" or throw an error
	file, err := os.Open("mesh.obj")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	// Decode the mesh data
	mesh, err := gomesh.DecodeOBJ(file)
	if err != nil {
		log.Fatalln(err)
	}

	// Print all vertices
	log.Println("Vertices:")
	for i, v := range mesh.Vertices {
		log.Printf("\tVertex: %4d %8.2f %8.2f %8.2f\n", i, v.X, v.Y, v.Z)
	}
	log.Println()

	// Print all faces
	log.Println("Faces:")
	for _, f := range mesh.Faces {
		log.Println("\tFace:")

		// Print vertex infos
		for _, vID := range f.Vertices {
			v := mesh.Vertices[vID]
			log.Printf("\t\tVertex: %8.2f %8.2f %8.2f\n", v.X, v.Y, v.Z)
		}
	}
}
