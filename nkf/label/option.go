package label

import (
	"github.com/Kernle32DLL/nuklear-fridge/nkf"
	"github.com/golang-ui/nuklear/nk"
)

type Option struct {
	content       string
	active        bool
	clickListener func(option *Option)
}

func NewOption(content string, active bool, clickListener func(option *Option)) *Option {
	return &Option{
		content:       content,
		active:        active,
		clickListener: clickListener,
	}
}

func NewOptionAsChild(parent nkf.RecursiveDrawable, content string, active bool, clickListener func(option *Option)) *Option {
	option := NewOption(content, active, clickListener)
	parent.AddChild(option)

	return option
}

func (option *Option) Render(ctx *nk.Context) {
	var activeFlag int32 = 0
	if option.active {
		activeFlag = 1
	}
	if clicked := nk.NkOptionLabel(ctx, option.content, activeFlag); clicked != activeFlag && option.clickListener != nil {
		option.clickListener(option)
	}
}

func (option *Option) Content() string {
	return option.content
}

func (option *Option) SetContent(content string) {
	option.content = content
}

func (option *Option) Active() bool {
	return option.active
}

func (option *Option) SetActive(active bool) {
	option.active = active
}
