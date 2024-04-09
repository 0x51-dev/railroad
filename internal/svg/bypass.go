package svg

import "fmt"

type Bypass struct {
	Element SVGable
}

func (b Bypass) SVG(start Point) (string, Point, Box) {
	wrappedStart := Point{X: start.X + curveRadius*2, Y: start.Y}
	wrappedSVG, wrappedEnd, wrappedBox := b.Element.SVG(wrappedStart)

	l := wrappedBox.Size.Y + (wrappedBox.Position.Y - start.Y) - curveRadius

	curveStartSVG, curveStart, curveStartBox := CurveInvS{L: l}.SVG(start)
	curveEndSVG, curveEnd, curveEndBox := CurveS{
		L: curveStart.Y - wrappedEnd.Y - 2*curveRadius,
	}.SVG(Point{
		X: curveStart.X + wrappedBox.Size.X,
		Y: curveStart.Y,
	})
	return fmt.Sprintf(
		`<g class="rr-bypass">
<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="black"/>
%s
<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="black"/>
%s
<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="black"/>
%s
</g>`,
		start.X, start.Y, wrappedStart.X, wrappedStart.Y,
		wrappedSVG,
		wrappedEnd.X, wrappedEnd.Y, curveEnd.X, curveEnd.Y,

		curveStartSVG,
		curveStart.X, curveStart.Y, curveStart.X+wrappedBox.Size.X, curveStart.Y,
		curveEndSVG,
	), curveEnd, curveStartBox.Combine(wrappedBox).Combine(curveEndBox)
}
