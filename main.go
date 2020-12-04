package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 800
	screenHeight = 400
)

func main() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		fmt.Println("initializing SDL:", err)
		return
	}

	window, err := sdl.CreateWindow(
		"Mario!",
		sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight,
		sdl.WINDOW_OPENGL)
	if err != nil {
		fmt.Println("initializing window:", err)
		return
	}
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Println("initializing renderer:", err)
		return
	}
	defer renderer.Destroy()

	background, err := newBackground(renderer)
	if err != nil {
		fmt.Println("create new background:", err)
		return
	}

	mushroom, err := newMushroom(renderer)
	if err != nil {
		fmt.Println("create new mushroom:", err)
		return
	}

	mario, err := newMario(renderer)
	if err != nil {
		fmt.Println("create new mario:", err)
		return
	}

	var hit int
	var meter int64

	go func() {
		for {
			meter++
			time.Sleep(5 * time.Millisecond)
			if mushroom.x+50 > mario.x && mushroom.x-50 < mario.x && !mario.jumping {
				hit++
			}

			window.SetTitle(fmt.Sprintf("Mario Game | Meter: %d | Hit: %d", meter, hit))
		}
	}()

	for {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				return
			}
		}
		renderer.SetDrawColor(204, 204, 204, 255)
		renderer.Clear()

		background.draw(renderer)

		mushroom.draw(renderer)
		mushroom.update()

		mario.draw(renderer)
		mario.update()

		renderer.Present()
	}
}
