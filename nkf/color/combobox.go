package color

import (
	"github.com/Kernle32DLL/nuklear-fridge/nkf"
	"github.com/golang-ui/nuklear/nk"
)

type ComboBox struct {
	children nkf.ChildrenContainer

	width     float32
	maxHeight float32

	color nk.Color

	clickListener func(comboBox *ComboBox)
}

func NewComboBox(maxWidth, maxHeight float32, initialColor nk.Color, clickListener func(comboBox *ComboBox)) *ComboBox {
	return &ComboBox{
		width:         maxWidth,
		maxHeight:     maxHeight,
		color:         initialColor,
		clickListener: clickListener,
	}
}

func NewComboBoxAsChild(parent nkf.RecursiveDrawable, maxWidth, maxHeight float32, initialColor nk.Color, clickListener func(comboBox *ComboBox)) *ComboBox {
	comboBox := NewComboBox(maxWidth, maxHeight, initialColor, clickListener)
	parent.AddChild(comboBox)

	return comboBox
}

func (comboBox *ComboBox) Render(ctx *nk.Context) {

	var width float32
	if comboBox.width < 0 {
		width = nk.NkWidgetWidth(ctx)
	}

	if nk.NkComboBeginColor(ctx, comboBox.color, nk.NkVec2(width, comboBox.maxHeight)) > 0 {
		comboBox.children.Render(ctx)

		nk.NkComboEnd(ctx)
	}
}

func (comboBox *ComboBox) Color() nk.Color {
	return comboBox.color
}

func (comboBox *ComboBox) SetColor(color nk.Color) {
	comboBox.color = color
}

func (comboBox *ComboBox) Width() float32 {
	return comboBox.width
}

func (comboBox *ComboBox) SetWidth(width float32) {
	comboBox.width = width
}

func (comboBox *ComboBox) MaxHeight() float32 {
	return comboBox.maxHeight
}

func (comboBox *ComboBox) SetMaxHeight(maxHeight float32) {
	comboBox.maxHeight = maxHeight
}

func (comboBox *ComboBox) AddChild(drawable nkf.Drawable) {
	comboBox.children.AddChild(drawable)
}

func (comboBox *ComboBox) ClearChildren() {
	comboBox.children.ClearChildren()
}

func (comboBox *ComboBox) ChildByIndex(index int) nkf.Drawable {
	return comboBox.children.ChildByIndex(index)
}

func (comboBox *ComboBox) IndexOfChild(drawable nkf.Drawable) int {
	return comboBox.children.IndexOfChild(drawable)
}
