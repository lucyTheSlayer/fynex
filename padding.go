package fynex

import "fyne.io/fyne/v2"

type PaddingLayout struct {
	paddingTop    float32
	paddingRight  float32
	paddingBottom float32
	paddingLeft   float32
}

func NewPaddingLayout(paddingTop, paddingRight, paddingBottom, paddingLeft float32) *PaddingLayout {
	return &PaddingLayout{
		paddingTop:    paddingTop,
		paddingRight:  paddingRight,
		paddingBottom: paddingBottom,
		paddingLeft:   paddingLeft,
	}
}

func (p *PaddingLayout) MinSize(objects []fyne.CanvasObject) fyne.Size {
	w, h := p.paddingLeft+p.paddingRight, p.paddingTop+p.paddingBottom
	if len(objects) > 0 {
		childSize := objects[0].MinSize()
		w += childSize.Width
		h += childSize.Height
	}
	return fyne.NewSize(w, h)
}

func (p *PaddingLayout) Layout(objects []fyne.CanvasObject, containerSize fyne.Size) {
	if len(objects) == 0 {
		return
	}
	o := objects[0]
	o.Move(fyne.NewPos(p.paddingLeft, p.paddingTop))
	o.Resize(fyne.NewSize(containerSize.Width-p.paddingLeft-p.paddingRight, containerSize.Height-p.paddingTop-p.paddingBottom))
}
