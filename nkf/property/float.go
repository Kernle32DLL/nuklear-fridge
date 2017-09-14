package property

import (
	"github.com/Kernle32DLL/nuklear-fridge/nkf"
	"github.com/golang-ui/nuklear/nk"
)

type Float struct {
	name string

	min, value, max, step float32
	incPerPixel           float32

	changeListener func(float *Float)
}

func NewFloat(name string, min, initialValue, max, step float32, incPerPixel float32, changeListener func(float *Float)) *Float {
	return &Float{
		name:           name,
		min:            min,
		value:          initialValue,
		max:            max,
		step:           step,
		incPerPixel:    incPerPixel,
		changeListener: changeListener,
	}
}

func NewFloatAsChild(parent nkf.RecursiveDrawable, name string, min, initialValue, max, step float32, incPerPixel float32, changeListener func(float *Float)) *Float {
	float := NewFloat(name, min, initialValue, max, step, incPerPixel, changeListener)
	parent.AddChild(float)

	return float
}

func (float *Float) Render(ctx *nk.Context) {
	prevValue := float.value
	nk.NkPropertyFloat(ctx, float.name, float.min, &float.value, float.max, float.step, float.incPerPixel)

	if prevValue != float.value && float.changeListener != nil {
		float.changeListener(float)
	}
}

func (float *Float) Name() string {
	return float.name
}

func (float *Float) SetName(name string) {
	float.name = name
}

func (float *Float) Value() float32 {
	return float.value
}

func (float *Float) SetValue(value float32) {
	float.value = value
}

func (float *Float) Max() float32 {
	return float.max
}

func (float *Float) SetMax(max float32) {
	float.max = max
}

func (float *Float) Min() float32 {
	return float.min
}

func (float *Float) SetMin(min float32) {
	float.min = min
}
