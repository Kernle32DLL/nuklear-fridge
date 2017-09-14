package label

import (
	"github.com/Kernle32DLL/nuklear-fridge/nkf"
	"github.com/golang-ui/nuklear/nk"
)

type ComboBox struct {
	children nkf.ChildrenContainer

	width     float32
	maxHeight float32

	content string

	clickListener func(comboBox *ComboBox)
}

func NewComboBox(maxWidth, maxHeight float32, initialText string, clickListener func(comboBox *ComboBox)) *ComboBox {
	return &ComboBox{
		width:         maxWidth,
		maxHeight:     maxHeight,
		content:       initialText,
		clickListener: clickListener,
	}
}

func NewComboBoxAsChild(parent nkf.RecursiveDrawable, maxWidth, maxHeight float32, initialText string, clickListener func(comboBox *ComboBox)) *ComboBox {
	comboBox := NewComboBox(maxWidth, maxHeight, initialText, clickListener)
	parent.AddChild(comboBox)

	return comboBox
}

func (comboBox *ComboBox) Render(ctx *nk.Context) {

	var width float32
	if comboBox.width < 0 {
		width = nk.NkWidgetWidth(ctx)
	}

	if nk.NkComboBeginLabel(ctx, comboBox.content, nk.NkVec2(width, comboBox.maxHeight)) > 0 {
		comboBox.children.Render(ctx)

		nk.NkComboEnd(ctx)
	}
}

func (comboBox *ComboBox) Content() string {
	return comboBox.content
}

func (comboBox *ComboBox) SetContent(content string) {
	comboBox.content = content
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
