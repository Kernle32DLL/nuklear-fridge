package row

import (
	"github.com/Kernle32DLL/nuklear-fridge/nkf"
	"github.com/golang-ui/nuklear/nk"
)

type Static struct {
	children nkf.ChildrenContainer

	height float32

	itemWidth int32
	cols      int32
}

func NewStatic(height float32, itemWidth, cols int32) *Static {
	return &Static{
		height:    height,
		itemWidth: itemWidth,
		cols:      cols,
	}
}

func NewStaticAsChild(parent nkf.RecursiveDrawable, height float32, itemWidth, cols int32) *Static {
	static := NewStatic(height, itemWidth, cols)

	parent.AddChild(static)

	return static
}

func (static *Static) Render(ctx *nk.Context) {
	nk.NkLayoutRowStatic(ctx, static.height, static.itemWidth, static.cols)
	static.children.Render(ctx)
}

func (static *Static) AddChild(drawable nkf.Drawable) {
	static.children.AddChild(drawable)
}

func (static *Static) ClearChildren() {
	static.children.ClearChildren()
}

func (static *Static) ChildByIndex(index int) nkf.Drawable {
	return static.children.ChildByIndex(index)
}

func (static *Static) IndexOfChild(drawable nkf.Drawable) int {
	return static.children.IndexOfChild(drawable)
}
