package nkf

import "github.com/golang-ui/nuklear/nk"

type Fridge struct {
	windows []*Window

	aa               bool
	maxVertexBuffer  int
	maxElementBuffer int
}

func NewFridge(aa bool, maxVertexBuffer int, maxElementBuffer int) *Fridge {
	return &Fridge{
		aa:               aa,
		maxVertexBuffer:  maxVertexBuffer,
		maxElementBuffer: maxElementBuffer,
	}
}

func (fridge *Fridge) Render(ctx *nk.Context) {
	nk.NkPlatformNewFrame()

	for _, window := range fridge.windows {
		window.Render(ctx)
	}

	var aa nk.AntiAliasing
	if fridge.aa {
		aa = nk.AntiAliasingOn
	} else {
		aa = nk.AntiAliasingOff
	}

	nk.NkPlatformRender(aa, fridge.maxVertexBuffer, fridge.maxElementBuffer)
}

func (fridge *Fridge) AddWindow(window *Window) {
	fridge.windows = append(fridge.windows, window)
}

func (fridge *Fridge) WindowByIndex(index int) *Window {
	return fridge.windows[index]
}
func (fridge *Fridge) IndexOfWindow(window Window) int {
	for p, v := range fridge.windows {
		if v.identifier == window.identifier {
			return p
		}
	}
	return -1
}
