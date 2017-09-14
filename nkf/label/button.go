package label

import (
	"github.com/Kernle32DLL/nuklear-fridge/nkf"
	"github.com/golang-ui/nuklear/nk"
)

type Button struct {
	content       string
	clickListener func(button *Button)
}

func NewButton(content string, clickListener func(button *Button)) *Button {
	return &Button{
		content:       content,
		clickListener: clickListener,
	}
}

func NewButtonAsChild(parent nkf.RecursiveDrawable, content string, clickListener func(button *Button)) *Button {
	buttonLabel := NewButton(content, clickListener)
	parent.AddChild(buttonLabel)

	return buttonLabel
}

func (button *Button) Render(ctx *nk.Context) {
	if clicked := nk.NkButtonLabel(ctx, button.content); clicked > 0 && button.clickListener != nil {
		button.clickListener(button)
	}
}

func (button *Button) Content() string {
	return button.content
}

func (button *Button) SetContent(content string) {
	button.content = content
}

func (button *Button) SetClickListener(clickListener func(buttonLabel *Button)) {
	button.clickListener = clickListener
}
