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
	"errors"
	"time"
)

// ErrAnimationExists is an error returned if an Animation already exists in the animator instance.
var ErrAnimationExists = errors.New("animation already exists in animator instance")

// ErrAnimationNotFound is an error returned if an Animation called by string is not found in the animator instance.
var ErrAnimationNotFound = errors.New("animation was not found")

// Animator is an interface that acts as a handler for animations.
type Animator interface {
	// Add will add an animation to the map of the animator instance.
	//
	// If an animation already exists in the map this function will
	// return an ErrAnimationExists error.
	Add(animationName string, animation *Animation) error

	// Play will play the animation, iterating through each frame.
	//
	// This function should be called every frame in your update function.
	//
	// If the animation name passed is not found in the animator's map an
	// ErrAnimationNotFound error is returned.
	Play(animationName string) error

	// Pause will stop an animation from playing until the next call to Resume.
	Pause()

	// Resume will start an animation until the next call to Pause.
	//
	// Note: you must still call the `Play()` method to iterate through the
	// frames of an animation.
	Resume()

	// Reset will reset an animation back to it's startFrame
	Reset()

	// Animation returns the currently selected animation.
	//
	// If no animation has been set using the Play() method
	// a zero valued Animation will be returned.
	Animation() *Animation

	// Paused returns true if the animator's current animation is paused, otherwise it returns false.
	Paused() bool

	// AnimationTime returns a time.Duration representing the total time it takes to perform a full animation sequence.
	AnimationTime() time.Duration
}

// NewAnimator returns a new Animator interface.
func NewAnimator() Animator {
	return &animator{
		index: make(map[string]*Animation),
	}
}

type animator struct {
	index            map[string]*Animation
	currentAnimation *Animation
}

// Add will add an animation to the map of animations. If the animation name already exists
// an ErrAnimationExists error is returned.
func (a *animator) Add(animationName string, animation *Animation) error {

	if _, ok := a.index[animationName]; ok {
		return ErrAnimationExists
	}

	a.index[animationName] = animation

	return nil
}

// Play will start playing the animation specified by the animation name.
func (a *animator) Play(animationName string) error {
	if _, ok := a.index[animationName]; !ok {
		return ErrAnimationNotFound
	}

	a.currentAnimation = a.index[animationName]
	a.index[animationName].update()

	return nil
}

// Resume starts playing an animation until the next call to Pause.
func (a *animator) Resume() {
	a.currentAnimation.paused = false
}

// Pause will stop the animation until the next call to Resume.
func (a *animator) Pause() {
	a.currentAnimation.paused = true
}

// Animation returns the animator's current animation.
func (a *animator) Animation() *Animation {
	return a.currentAnimation
}

// IsPlaying returns true if an animation is in a playing state, otherwise it returns false.
func (a *animator) Paused() bool {
	return a.currentAnimation.paused
}

// Reset reset's the currently playing animation to it's start frame.
func (a *animator) Reset() {
	a.currentAnimation.currentFrameIndex = 0
}

// AnimationTime returns a time.Duration representing the total time it takes to perform a full animation sequence.
func (a *animator) AnimationTime() time.Duration {
	return time.Duration(a.currentAnimation.FrameCount * int(a.currentAnimation.Delay))
}
