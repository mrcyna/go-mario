package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

type background struct {
	tex *sdl.Texture
}

func newBackground(renderer *sdl.Renderer) (b background, err error) {
	backgroundImg, err := img.Load("images/background.png")
	if err != nil {
		return background{}, fmt.Errorf("loading background image: %v", err)
	}
	defer backgroundImg.Free()

	b.tex, err = renderer.CreateTextureFromSurface(backgroundImg)
	if err != nil {
		return background{}, fmt.Errorf("creating background texture: %v", err)
	}

	return b, nil
}

func (b *background) draw(renderer *sdl.Renderer) {
	renderer.Copy(b.tex,
		&sdl.Rect{X: 0, Y: 0, W: 800, H: 400},
		&sdl.Rect{X: 0, Y: 0, W: 800, H: 400},
	)
}
