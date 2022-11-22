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
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

// Origin repesents the x, y drawing point of the sprite.
//
// This means if the x, y location of the origin is TopLeft
// the sprite's image will be drawn to the destination image
// from the TopLeft corner of the image.
//
// If the origin is centered, the sprite will be drawn to the
// destination image from the center of the sprite's image.
//
//	TopLeft ──┐
//	 Origin   │
//	          │     ┌── Center
//	          │     │    Origin
//	          │     │
//	          o─────┼──────┐
//	          │     │      │
//	          │     │      │
//	          │     o      │
//	          │            │
//	          │            │
//	          └────────────┘
type Origin int

const (
	// Topleft origin of a sprite image.
	TopLeft Origin = iota
	// Center origin of a sprite image.
	Center
)

// Vector represents a mathematical vector
// used for positioning and movement.
type Vector struct {
	X, Y float64
}

// Sprite represents any visible image on the game screen.
type Sprite struct {
	// Texture is the sprite's image.
	Texture *ebiten.Image
	// Position is the sprites location on screen.
	Position *Vector
	// Animator can be used to add animations to your sprite.
	Animator Animator
	// UpdateFunc handles your sprites update logic. This function should be
	// called in Ebitengine's Game Upadate() method.
	UpdateFunc func() error
	// Origin repesents the origin point of the sprite.
	Origin Origin
	// Color will fill the sprite with a give color.
	// This is mostly used for prototyping.
	Color color.Color
	// Scale represents the scale of the sprite.
	Scale Vector
	// flippedHorizontal determines if the sprite has been flipped horizontally.
	flippedHorizontal bool
	// flippedVertical determines if the sprite has been flipped vertically.
	flippedVertical bool
}

// Draw draws the sprite at the sprites current location.
func (s *Sprite) Draw(dst *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	if s.Color != nil {
		s.Texture.Fill(s.Color)
	}

	if s.Origin == Center {
		if s.Animator != nil && s.Animator.Animation() != nil {
			op.GeoM.Translate(-float64(s.Animator.Animation().FrameWidth()/2), -float64(s.Animator.Animation().FrameHeight()/2))
		} else {
			op.GeoM.Translate(-float64(s.Texture.Bounds().Dx()/2), -float64(s.Texture.Bounds().Dy()/2))
		}
	}

	if s.Scale.X == 0 && s.Scale.Y == 0 {
		s.Scale.X = 1
		s.Scale.Y = 1
	}
	op.GeoM.Scale(s.Scale.X, s.Scale.Y)

	op.GeoM.Translate(s.Position.X, s.Position.Y)

	var tex *ebiten.Image
	if s.Animator != nil && s.Animator.Animation() != nil {
		tex = s.Texture.SubImage(s.Animator.Animation().Frame()).(*ebiten.Image)
	} else {
		tex = s.Texture
	}

	dst.DrawImage(tex, op)
}

// Update performs the sprite's logical updates.
func (s *Sprite) Update() {
	s.UpdateFunc()
}

// Scale will set the scale of the sprite texture.
func (s *Sprite) SetScale(x, y float64) {
	s.Scale.X = x
	s.Scale.Y = y
}

// FlipHorizontal will flip a sprite horizontally if flipped is set to true.
func (s *Sprite) FlipHorizontal(flipped bool) {
	if flipped {
		if !s.flippedHorizontal {
			s.Scale.X = -s.Scale.X
		}
		s.flippedHorizontal = true
	} else {
		s.Scale.X = math.Abs(s.Scale.X)
		s.flippedHorizontal = false
	}
}

// FlipVertical will flip a sprite vertically if flipped is set to true.
func (s *Sprite) FlipVertical(flipped bool) {
	if flipped {
		if !s.flippedVertical {
			s.Scale.Y = -s.Scale.Y
		}
		s.flippedVertical = true
	} else {
		s.Scale.Y = math.Abs(s.Scale.Y)
		s.flippedVertical = false
	}
}
