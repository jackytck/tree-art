package main

import "github.com/fogleman/ln/ln"

func main() {
	// create a scene and add a single cube
	scene := ln.Scene{}
	scene.Add(ln.NewCube(ln.Vector{X: -1, Y: -1, Z: -1}, ln.Vector{X: 1, Y: 1, Z: 1}))

	// define camera parameters
	eye := ln.Vector{X: 4, Y: 3, Z: 2}    // camera position
	center := ln.Vector{X: 0, Y: 0, Z: 0} // camera looks at
	up := ln.Vector{X: 0, Y: 0, Z: 1}     // up direction

	// define rendering parameters
	width := 1024.0  // rendered width
	height := 1024.0 // rendered height
	fovy := 50.0     // vertical field of view, degrees
	znear := 0.1     // near z plane
	zfar := 10.0     // far z plane
	step := 0.01     // how finely to chop the paths for visibility testing

	// compute 2D paths that depict the 3D scene
	paths := scene.Render(eye, center, up, width, height, fovy, znear, zfar, step)

	// render the paths in an image
	paths.WriteToPNG("out.png", width, height)

	// save the paths as an svg
	paths.WriteToSVG("out.svg", width, height)
}
