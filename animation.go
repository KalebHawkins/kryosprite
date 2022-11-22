// Copyright 2022 Kaleb Hawkins

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kryosprite

import (
	"image"
	"time"
)

// Direction repestents the Direction in which the animation is read from a spritesheet.
type Direction int

const (
	// Horizontal, when provided to an Animation's Direction, will read the spritesheet from left to right.
	Horizontal Direction = iota
	// Vertical, when provided to an Animation's Direction, will read the spritesheet from top to bottom.
	Vertical
)

// Animation defines a spritesheet's animation.
type Animation struct {
	// StartFrame is the frame's starting x and y location, width and height.
	StartFrame Frame
	// FrameCount the number of frames for a particular animation.
	FrameCount int
	// Delay is the time it takes for each frame of the animation to pass.
	Delay time.Duration
	// Direction is the Direction the spritesheet should be read.
	// This should be either Horizontal or Vertical.
	Direction Direction

	// currentFrame is the animations currently selected frame.
	currentFrame Frame
	// currentFrameIndex repesents the index of the current frame.
	currentFrameIndex int
	// paused is a boolean representing the pause state of the animation.
	paused bool
	// prevFrame repesents the time between the frame iterations.
	prevFrame time.Time
	// deltaTime repesents the time delta between frames.
	deltaTime float64
}

// update iterates through each frame of a sprite sheet.
func (a *Animation) update() {
	if a.paused {
		return
	}

	a.deltaTime += float64(time.Since(a.prevFrame))
	a.prevFrame = time.Now()

	switch a.Direction {
	case Horizontal:
		a.currentFrame.X = a.currentFrameIndex * a.StartFrame.Width
		a.currentFrame.Y = a.StartFrame.Y
		a.currentFrame.Width = a.currentFrame.X + a.StartFrame.Width
		a.currentFrame.Height = a.currentFrame.Y + a.StartFrame.Height
	case Vertical:
		a.currentFrame.X = a.StartFrame.Width
		a.currentFrame.Y = a.currentFrameIndex * a.StartFrame.Y
		a.currentFrame.Width = a.currentFrame.X + a.StartFrame.Width
		a.currentFrame.Height = a.currentFrame.Y + a.StartFrame.Height
	}

	if a.deltaTime >= float64(a.Delay) {
		a.deltaTime = 0

		a.currentFrameIndex++
		if a.currentFrameIndex >= a.FrameCount {
			a.currentFrameIndex = 0
		}
	}
}

// Frame returns a rect repesenting the current frame of the animation.
func (a *Animation) Frame() image.Rectangle {
	return image.Rectangle{
		Min: image.Point{
			X: a.currentFrame.X,
			Y: a.currentFrame.Y,
		},
		Max: image.Point{
			X: a.currentFrame.Width,
			Y: a.currentFrame.Height,
		},
	}
}

// FrameWidth returns the width of the frame.
func (a *Animation) FrameWidth() int {
	return a.currentFrame.Width - a.currentFrame.X
}

// FrameHeight returns the height of the frame.
func (a *Animation) FrameHeight() int {
	return a.currentFrame.Height - a.currentFrame.Y
}
