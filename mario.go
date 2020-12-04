package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type mario struct {
	tex     *sdl.Texture
	x       float64
	y       float64
	jumping bool
}

func newMario(renderer *sdl.Renderer) (m mario, err error) {
	marioImg, err := img.Load("images/mario.png")
	if err != nil {
		return mario{}, fmt.Errorf("loading mario image: %v", err)
	}
	defer marioImg.Free()

	m.tex, err = renderer.CreateTextureFromSurface(marioImg)
	if err != nil {
		return mario{}, fmt.Errorf("creating mario texture: %v", err)
	}

	m.x = 50
	m.y = 250
	return m, nil
}

func (m *mario) draw(renderer *sdl.Renderer) {
	renderer.Copy(m.tex,
		&sdl.Rect{X: 0, Y: 0, W: 102, H: 101},
		&sdl.Rect{X: int32(m.x), Y: int32(m.y), W: 102, H: 101},
	)
}

func (m *mario) update() {
	keys := sdl.GetKeyboardState()

	if keys[sdl.SCANCODE_SPACE] == 1 && !m.jumping {
		m.jump()
	}
}

func (m *mario) jump() {
	m.jumping = true

	go func() {
		for i := 0; i < 130; i++ {
			time.Sleep(2 * time.Millisecond)
			m.y = m.y - 1
		}

		go func() {
			for i := 0; i < 130; i++ {
				time.Sleep(3 * time.Millisecond)
				m.y = m.y + 1
			}

			m.jumping = false
		}()
	}()
}
