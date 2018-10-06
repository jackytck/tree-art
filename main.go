package main

import (
	"math"

	"github.com/fogleman/ln/ln"
)

func main() {
	cube("cube")
	hole(xxyy, "hole")
	sphere("sphere", false)
	sphere("outline-sphere", true)
	cylinder("cylinder", false)
	cylinder("outline-cylinder", true)
}

func cube(out string) {
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
	paths.WriteToPNG(out+".png", width, height)

	// save the paths as an svg
	paths.WriteToSVG(out+".svg", width, height)
}

func xxyy(x, y float64) float64 {
	return -1 / (x*x + y*y)
}

func cosxy(x, y float64) float64 {
	return math.Cos(x*y) * (x*x - y*y)
}

func hole(f func(x, y float64) float64, out string) {
	scene := ln.Scene{}

	// add func
	box := ln.Box{
		Min: ln.Vector{X: -2, Y: -2, Z: -4},
		Max: ln.Vector{X: 2, Y: 2, Z: 2},
	}
	scene.Add(ln.NewFunction(f, box, ln.Below))

	eye := ln.Vector{X: 3, Y: 0, Z: 3}
	center := ln.Vector{X: 1.1, Y: 0, Z: 0}
	up := ln.Vector{X: 0, Y: 0, Z: 1}
	width := 1024.0
	height := 1024.0

	paths := scene.Render(eye, center, up, width, height, 50, 0.1, 100, 0.01)
	paths.WriteToPNG(out+".png", width, height)
	paths.WriteToSVG(out+".svg", width, height)
}

func sphere(out string, outline bool) {
	scene := ln.Scene{}

	eye := ln.Vector{X: 3, Y: 0, Z: 3}
	center := ln.Vector{X: 1.1, Y: 0, Z: 0}
	up := ln.Vector{X: 0, Y: 0, Z: 1}
	radius := 0.333
	width := 1024.0
	height := 1024.0

	if outline {
		o := ln.NewOutlineSphere(eye, up, center, radius)
		scene.Add(o)
	} else {
		s := ln.NewSphere(center, radius)
		scene.Add(s)
	}

	paths := scene.Render(eye, center, up, width, height, 50, 0.1, 100, 0.01)
	paths.WriteToPNG(out+".png", width, height)
	paths.WriteToSVG(out+".svg", width, height)
}

func cylinder(out string, outline bool) {
	scene := ln.Scene{}

	eye := ln.Vector{X: 3, Y: 0, Z: 3}
	center := ln.Vector{X: 1.1, Y: 0, Z: 0}
	up := ln.Vector{X: 0, Y: 0, Z: 1}
	radius := 0.333
	width := 1024.0
	height := 1024.0

	if outline {
		o := ln.NewOutlineCylinder(eye, up, radius, -3, 3)
		scene.Add(o)
	} else {
		s := ln.NewCylinder(radius, -3, 3)
		scene.Add(s)
	}

	paths := scene.Render(eye, center, up, width, height, 50, 0.1, 100, 0.01)
	paths.WriteToPNG(out+".png", width, height)
	paths.WriteToSVG(out+".svg", width, height)
}
