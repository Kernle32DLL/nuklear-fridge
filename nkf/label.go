package nkf

import "github.com/golang-ui/nuklear/nk"

type Label struct {
	content   string
	alignment nk.Flags
}

func NewLabel(content string, alignment nk.Flags) *Label {
	return &Label{
		content:   content,
		alignment: alignment,
	}
}

func NewLabelAsChild(parent RecursiveDrawable, content string, alignment nk.Flags) *Label {
	label := NewLabel(content, alignment)
	parent.AddChild(label)

	return label
}

func (label *Label) Render(ctx *nk.Context) {
	nk.NkLabel(ctx, label.content, label.alignment)
}

func (label *Label) Content() string {
	return label.content
}

func (label *Label) SetContent(content string) {
	label.content = content
}
