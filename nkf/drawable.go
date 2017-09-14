package nkf

import "github.com/golang-ui/nuklear/nk"

type Drawable interface {
	Render(ctx *nk.Context)
}

type RecursiveDrawable interface {
	Render(ctx *nk.Context)
	AddChild(drawable Drawable)
	ClearChildren()
	ChildByIndex(index int) Drawable
	IndexOfChild(drawable Drawable) int
}

type ChildrenContainer struct {
	children []Drawable
}

func (childrenContainer *ChildrenContainer) Render(ctx *nk.Context) {
	for _, drawable := range childrenContainer.children {
		drawable.Render(ctx)
	}
}

func (childrenContainer *ChildrenContainer) AddChild(drawable Drawable) {
	childrenContainer.children = append(childrenContainer.children, drawable)
}

func (childrenContainer *ChildrenContainer) ClearChildren() {
	childrenContainer.children = []Drawable{}
}

func (childrenContainer *ChildrenContainer) ChildByIndex(index int) Drawable {
	return childrenContainer.children[index]
}
func (childrenContainer *ChildrenContainer) IndexOfChild(drawable Drawable) int {
	for p, v := range childrenContainer.children {
		if v == drawable {
			return p
		}
	}
	return -1
}
