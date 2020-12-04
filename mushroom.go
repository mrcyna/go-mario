package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type mushroom struct {
	tex     *sdl.Texture
	x       float64
	y       float64
	started bool
}

func newMushroom(renderer *sdl.Renderer) (m mushroom, err error) {
	mushroomImg, err := img.Load("images/mushroom.png")
	if err != nil {
		return mushroom{}, fmt.Errorf("loading mushroom image: %v", err)
	}
	defer mushroomImg.Free()

	m.tex, err = renderer.CreateTextureFromSurface(mushroomImg)
	if err != nil {
		return mushroom{}, fmt.Errorf("creating mushroom texture: %v", err)
	}

	m.x = 500
	m.y = 305

	return m, nil
}

func (m *mushroom) draw(renderer *sdl.Renderer) {
	renderer.Copy(m.tex,
		&sdl.Rect{X: 0, Y: 0, W: 100, H: 100},
		&sdl.Rect{X: int32(m.x), Y: int32(m.y), W: 40, H: 40},
	)
}

func (m *mushroom) update() {
	if m.started {
		return
	}

	m.started = true
	go func() {
		for {
			time.Sleep(3 * time.Millisecond)
			m.x = m.x - 1.1

			if m.x < -100 {
				m.x = 800
			}
		}
	}()
}
