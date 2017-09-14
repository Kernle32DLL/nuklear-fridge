package color

import (
	"github.com/Kernle32DLL/nuklear-fridge/nkf"
	"github.com/golang-ui/nuklear/nk"
)

type Picker struct {
	color  nk.Color
	format nk.ColorFormat

	changeListener func(picker *Picker)
}

func NewPicker(initialColor nk.Color, format nk.ColorFormat, changeListener func(picker *Picker)) *Picker {
	return &Picker{
		color:          initialColor,
		format:         format,
		changeListener: changeListener,
	}
}

func NewPickerAsChild(parent nkf.RecursiveDrawable, initialColor nk.Color, format nk.ColorFormat, changeListener func(picker *Picker)) *Picker {
	picker := NewPicker(initialColor, format, changeListener)
	parent.AddChild(picker)

	return picker
}

func (picker *Picker) Render(ctx *nk.Context) {
	prevValue := picker.color
	picker.color = nk.NkColorPicker(ctx, picker.color, picker.format)

	if prevValue != picker.color && picker.changeListener != nil {
		picker.changeListener(picker)
	}
}

func (picker *Picker) Color() nk.Color {
	return picker.color
}

func (picker *Picker) SetColor(color nk.Color) {
	picker.color = color
}

func (picker *Picker) ColorFormat() nk.ColorFormat {
	return picker.format
}

func (picker *Picker) SetColorFormat(colorFormat nk.ColorFormat) {
	picker.format = colorFormat
}
