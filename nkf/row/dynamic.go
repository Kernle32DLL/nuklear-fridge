package row

import (
	"github.com/Kernle32DLL/nuklear-fridge/nkf"
	"github.com/golang-ui/nuklear/nk"
)

type Dynamic struct {
	children nkf.ChildrenContainer

	height float32
	cols   int32
}

func NewDynamic(height float32, cols int32) *Dynamic {
	return &Dynamic{
		height: height,
		cols:   cols,
	}
}

func NewDynamicAsChild(parent nkf.RecursiveDrawable, height float32, cols int32) *Dynamic {
	dynamic := NewDynamic(height, cols)
	parent.AddChild(dynamic)

	return dynamic
}

func (dynamic *Dynamic) Render(ctx *nk.Context) {
	nk.NkLayoutRowDynamic(ctx, dynamic.height, dynamic.cols)
	dynamic.children.Render(ctx)
}

func (dynamic *Dynamic) AddChild(drawable nkf.Drawable) {
	dynamic.children.AddChild(drawable)
}

func (dynamic *Dynamic) ClearChildren() {
	dynamic.children.ClearChildren()
}

func (dynamic *Dynamic) ChildByIndex(index int) nkf.Drawable {
	return dynamic.children.ChildByIndex(index)
}

func (dynamic *Dynamic) IndexOfChild(drawable nkf.Drawable) int {
	return dynamic.children.IndexOfChild(drawable)
}
