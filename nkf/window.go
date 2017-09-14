package nkf

import (
	"github.com/golang-ui/nuklear/nk"
	"github.com/satori/go.uuid"
)

type windowState struct {
	x, y, w, h float32

	xChanged, yChanged bool
	wChanged, hChanged bool

	bounds nk.Rect
}

type Window struct {
	children ChildrenContainer
	state    windowState

	bgColor nk.Color

	title      string
	identifier string
	flags      nk.Flags
}

func NewWindow(x, y, w, h float32, title string, flags nk.Flags) *Window {
	return &Window{
		title:      title,
		identifier: uuid.NewV4().String(),
		flags:      flags,
		state: windowState{
			x:      x,
			y:      y,
			w:      w,
			h:      h,
			bounds: nk.NkRect(x, y, w, h),
		},
		bgColor: nk.NkRgba(28, 48, 62, 128),
	}
}

func (window Window) Identifier() string {
	return window.identifier
}

func (window Window) Position() (x float32, y float32) {
	return window.state.x, window.state.y
}

func (window *Window) SetPosition(x, y float32) {
	window.state.xChanged = window.state.x != x
	window.state.yChanged = window.state.y != y

	window.state.x = x
	window.state.y = y
}

func (window Window) Size() (w float32, h float32) {
	return window.state.w, window.state.h
}

func (window *Window) SetSize(w, h float32) {
	window.state.wChanged = window.state.w != w
	window.state.hChanged = window.state.h != h

	window.state.w = w
	window.state.h = h
}

func (window *Window) Render(ctx *nk.Context) {
	windowState := &window.state
	flags := window.flags

	// Hack to prevent visual artifacts while changing the window bounds both by the user and programmatically
	if windowState.xChanged || windowState.yChanged {
		flags = flags &^ nk.WindowMovable
	}
	if windowState.wChanged || windowState.hChanged {
		flags = flags &^ nk.WindowScalable
	}

	// Ensure that the bounds of the window are correctly set on programmatic changes
	if windowState.xChanged || windowState.yChanged || windowState.wChanged || windowState.hChanged {
		windowState.bounds = nk.NkRect(windowState.x, windowState.y, windowState.w, windowState.h)
		nk.NkWindowSetBounds(ctx, window.title, windowState.bounds)
	}

	update := nk.NkBeginTitled(ctx, window.identifier, window.title, windowState.bounds, flags)
	windowState.handleUserStateUpdate(ctx)

	if update > 0 {
		window.children.Render(ctx)
	}

	nk.NkEnd(ctx)
}

func (windowState *windowState) handleUserStateUpdate(ctx *nk.Context) {
	windowState.handleUserStatePositionUpdate(ctx)
	windowState.handleUserStateSizeUpdate(ctx)
}

func (windowState *windowState) handleUserStatePositionUpdate(ctx *nk.Context) {
	// Check for the current window position (takes into account user-movement of the window)
	userPosition := nk.NkWindowGetPosition(ctx)

	if !windowState.xChanged {
		windowState.x = userPosition.X()
	}

	if !windowState.yChanged {
		windowState.y = userPosition.Y()
	}

	// Reflect updated state
	windowState.xChanged = false
	windowState.yChanged = false
}

func (windowState *windowState) handleUserStateSizeUpdate(ctx *nk.Context) {
	// Check for the current window size (takes into account user-sizing of the window)
	userSize := nk.NkWindowGetSize(ctx)

	if !windowState.wChanged {
		windowState.w = userSize.X()
	}

	if !windowState.hChanged {
		windowState.h = userSize.Y()
	}

	// Reflect updated state
	windowState.wChanged = false
	windowState.hChanged = false
}

func (window *Window) AddChild(drawable Drawable) {
	window.children.AddChild(drawable)
}

func (window *Window) ClearChildren() {
	window.children.ClearChildren()
}

func (window *Window) ChildByIndex(index int) Drawable {
	return window.children.ChildByIndex(index)
}

func (window *Window) IndexOfChild(drawable Drawable) int {
	return window.children.IndexOfChild(drawable)
}
