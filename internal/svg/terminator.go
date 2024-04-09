package svg

import (
	"fmt"
)

// Start is a terminator that starts a path, i.e. "o-".
type Start struct{}

func (Start) SVG(start Point) (string, Point, Box) {
	return fmt.Sprintf(
			`<circle cx="%.1f" cy="%.1f" r="%.1f" fill="black"/>
<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="black"/>`,
			start.X, start.Y, // circle
			pointRadius,

			start.X+pointRadius/2, start.Y, // line
			start.X+pointRadius/2+interNodeSpace, start.Y,
		), Point{X: start.X + interNodeSpace, Y: start.Y}, Box{
			Position: Point{X: start.X - pointRadius, Y: start.Y - pointRadius},
			Size:     Point{X: pointRadius + interNodeSpace, Y: pointRadius * 2},
		}
}

// End is a terminator that ends a path, i.e. "-o".
type End struct{}

func (End) SVG(start Point) (string, Point, Box) {
	return fmt.Sprintf(
			`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="black"/>
<circle cx="%.1f" cy="%.1f" r="%.1f" fill="black"/>`,
			start.X, start.Y, // line
			start.X+interNodeSpace, start.Y,

			start.X+interNodeSpace, start.Y, // circle
			pointRadius,
		), Point{X: start.X + interNodeSpace, Y: start.Y}, Box{
			Position: Point{X: start.X, Y: start.Y - pointRadius},
			Size:     Point{X: interNodeSpace + pointRadius, Y: pointRadius * 2},
		}
}
