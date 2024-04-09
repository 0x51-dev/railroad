package svg

import "fmt"

type TextBox struct {
	Text           string
	RoundedBorders bool
}

func (t TextBox) SVG(start Point) (string, Point, Box) {
	offset := 10.0
	textWidth := fontSize / 2
	rounding := 1.0
	if t.RoundedBorders {
		rounding = textWidth
	}
	textLength := float64(len(t.Text)) * textWidth
	return fmt.Sprintf(
			`<g class="rr-text-box">
<rect x="%0.1f" y="%.1f" rx="%0.1f" ry="%0.1f" width="%.1f" height="%.1f" fill="white" stroke="black"/>
<text x="%.1f" y="%.1f" font-size="%.1f" textLength="%.1f" text-anchor="middle" dominant-baseline="middle">%s</text>
</g>`,
			start.X, start.Y-textWidth-offset/2,
			rounding, rounding,
			textLength+offset, fontSize+offset,

			start.X+textLength/2+offset/2, start.Y,
			fontSize, textLength,
			t.Text,
		), Point{X: start.X + textLength + offset, Y: start.Y}, Box{
			Position: Point{X: start.X, Y: start.Y - textWidth - offset/2},
			Size:     Point{X: textLength + offset, Y: fontSize + offset},
		}
}
