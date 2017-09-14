package property

import (
	"github.com/Kernle32DLL/nuklear-fridge/nkf"
	"github.com/golang-ui/nuklear/nk"
)

type Integer struct {
	name string

	min, value, max, step int32
	incPerPixel           float32

	changeListener func(integer *Integer)
}

func NewInteger(name string, min, initialValue, max, step int32, incPerPixel float32, changeListener func(integer *Integer)) *Integer {
	return &Integer{
		name:           name,
		min:            min,
		value:          initialValue,
		max:            max,
		step:           step,
		incPerPixel:    incPerPixel,
		changeListener: changeListener,
	}
}

func NewIntegerAsChild(parent nkf.RecursiveDrawable, name string, min, initialValue, max, step int32, incPerPixel float32, changeListener func(integer *Integer)) *Integer {
	integer := NewInteger(name, min, initialValue, max, step, incPerPixel, changeListener)
	parent.AddChild(integer)

	return integer
}

func (integer *Integer) Render(ctx *nk.Context) {
	prevValue := integer.value
	nk.NkPropertyInt(ctx, integer.name, integer.min, &integer.value, integer.max, integer.step, integer.incPerPixel)

	if prevValue != integer.value && integer.changeListener != nil {
		integer.changeListener(integer)
	}
}

func (integer *Integer) Name() string {
	return integer.name
}

func (integer *Integer) SetName(name string) {
	integer.name = name
}

func (integer *Integer) Value() int32 {
	return integer.value
}

func (integer *Integer) SetValue(value int32) {
	integer.value = value
}

func (integer *Integer) Max() int32 {
	return integer.max
}

func (integer *Integer) SetMax(max int32) {
	integer.max = max
}

func (integer *Integer) Min() int32 {
	return integer.min
}

func (integer *Integer) SetMin(min int32) {
	integer.min = min
}
