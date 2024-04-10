package svg

import (
	"fmt"
)

type Bypass struct {
	Element SVGable
}

func (b Bypass) SVG(start Point) (string, Point, Box) {
	_, _, wrappedBox := b.Element.SVG(start)
	l := start.Y - wrappedBox.Position.Y + curveRadius

	curveStartSVG, curveStart, curveStartBox := Line{RelativeEnd: Point{X: 2 * curveRadius, Y: l}}.SVG(start)
	wrappedSVG, wrappedEnd, wrappedBox := b.Element.SVG(curveStart)
	curveEndSVG, curveEnd, curveEndBox := Line{RelativeEnd: Point{X: 2 * curveRadius, Y: -l}}.SVG(wrappedEnd)
	return fmt.Sprintf(`
<g class="rr-bypass">
<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="black"/>%s
%s
%s
</g>`,
		start.X, start.Y, curveEnd.X, curveEnd.Y,
		curveStartSVG,
		wrappedSVG,
		curveEndSVG,
	), curveEnd, curveStartBox.Combine(wrappedBox).Combine(curveEndBox)
}
