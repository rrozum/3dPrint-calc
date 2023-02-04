package desktop

import (
	"fyne.io/fyne/v2"
)

const (
	minWidth  float32 = 200
	minHeight float32 = 40
	outer     float32 = 16 // расстояние по бокам
	rowSpace  float32 = 8  // растояние между строками
)

type ProjectLayout struct {
	rows int
}

func (l *ProjectLayout) Layout(objs []fyne.CanvasObject, s fyne.Size) {
	l.rows = len(objs)
	rowWidth := s.Width - outer*2
	minWidthOuter := minWidth + outer*2
	if rowWidth < minWidthOuter {
		rowWidth = minWidthOuter
	}

	cellSize := fyne.NewSize(rowWidth, minHeight)

	offset := 0
	pos := fyne.Position{X: outer, Y: outer}
	for _, o := range objs {
		o.Resize(cellSize)
		o.Move(pos)

		offset++
		if offset >= l.rows {
			offset = 0
		}
		pos.Y += minHeight + rowSpace
	}
}

func (l *ProjectLayout) MinSize(cells []fyne.CanvasObject) fyne.Size {
	height := (minHeight + rowSpace*2) * float32(len(cells))

	return fyne.NewSize(minWidth+outer*2, height)
}
