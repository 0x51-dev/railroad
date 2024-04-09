package svg

import "fmt"

type SVGable interface {
	SVG(start Point) (svg string, end Point, box Box)
}

func DebugSVG(element SVGable, start Point) (string, Point, Box) {
	svgElement, end, box := element.SVG(start)
	return fmt.Sprintf(
		`<g class="debug">
<circle cx="%.1f" cy="%.1f" r="2" fill="transparent" stroke="green" stroke-width="0.5"/>
%s
<circle cx="%.1f" cy="%.1f" r="2" fill="transparent" stroke="green" stroke-width="0.5"/>
<rect x="%.1f" y="%.1f" width="%.1f" height="%.1f" fill="transparent" stroke="red" stroke-width="0.5"/>
</g>`,
		start.X, start.Y,
		svgElement,
		end.X, end.Y,
		box.Position.X, box.Position.Y, box.Size.X, box.Size.Y,
	), end, box
}
