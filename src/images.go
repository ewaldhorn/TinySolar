package main

import (
	"syscall/js"

	"github.com/ewaldhorn/dommie/dom"
)

var sun, earth, moon js.Value

// ----------------------------------------------------------------------------
// load the image assets
func setupImages() {
	sun = dom.CreateImg("sun.png")
	earth = dom.CreateImg("earth.png")
	moon = dom.CreateImg("moon.png")
}
