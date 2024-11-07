package fynex

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/driver/mobile"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

type NumericalEntry struct {
	widget.Entry
	extraWidth float32
}

func NewNumericalEntry(extraWidth float32) *NumericalEntry {
	entry := &NumericalEntry{
		extraWidth: extraWidth,
	}
	entry.ExtendBaseWidget(entry)
	return entry
}

func NewNumericalEntryWithData(extraWidth float32, data binding.String) *NumericalEntry {
	entry := NewNumericalEntry(extraWidth)
	entry.Bind(data)
	return entry
}

func (e *NumericalEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' || r == ',' {
		e.Entry.TypedRune(r)
	}
}

func (e *NumericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}

	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func (e *NumericalEntry) Keyboard() mobile.KeyboardType {
	return mobile.NumberKeyboard
}

func (e *NumericalEntry) MinSize() fyne.Size {
	return fyne.NewSize(e.Entry.MinSize().Width+e.extraWidth, e.Entry.MinSize().Height)
}
