package main

import (
	"math"
	"syscall/js"
	"time"
)

var (
	earthAngle, moonAngle float64
	fullCircle            float64
	earthChangePerSecond  float64
	moonChangePerSecond   float64
)

// ----------------------------------------------------------------------------
func setupAnimation() {
	fullCircle = 2.0 * math.Pi
	earthChangePerSecond = fullCircle / 60.0
	moonChangePerSecond = fullCircle / 6.0
	startAnimation()
}

// ----------------------------------------------------------------------------
func render() {
	updateAngles()

	graphicsContext.GlobalCompositeOperation("destination-over")
	graphicsContext.ClearRect(0, 0, canvasWidth, canvasHeight)
	// predraw setups
	graphicsContext.FillStyle("rgb(0 0 0 / 40%)")
	graphicsContext.StrokeStyle("rgb(0 153 255 / 40%)")
	graphicsContext.Save()
	graphicsContext.Translate(150, 150)

	// draw earth
	graphicsContext.Rotate(earthAngle)
	graphicsContext.Translate(105, 0)
	graphicsContext.FillRect(0, -12, 40, 24)
	graphicsContext.DrawImage(earth, -12, -12)

	// moon
	graphicsContext.Save()
	graphicsContext.Rotate(moonAngle)
	graphicsContext.Translate(0, 28.5)
	graphicsContext.DrawImage(moon, -3.5, -3.5)
	graphicsContext.Restore()

	// earth orbit
	graphicsContext.Restore()
	graphicsContext.BeginPath()
	graphicsContext.Arc(150, 150, 105, 0, fullCircle)
	graphicsContext.Stroke()

	// finally, the sun
	graphicsContext.DrawImage(sun, 0, 0)
}

// ----------------------------------------------------------------------------
func updateAngles() {
	currentTime := time.Now()
	timeInSeconds := float64(currentTime.Second()) + float64(currentTime.Nanosecond())/1e9

	earthAngle = earthChangePerSecond * timeInSeconds
	moonAngle = moonChangePerSecond * timeInSeconds
}

// ----------------------------------------------------------------------------
// Allows JS to call into Wasm to refresh the effect.
func setRefreshEffectCallback() {
	js.Global().Set("refreshEffect", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		render()
		return nil
	}))
}
