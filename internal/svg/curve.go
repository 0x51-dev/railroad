package svg

import "fmt"

type CurveInvS struct {
	L float64
}

func (c CurveInvS) SVG(start Point) (string, Point, Box) {
	r, ry, x, y, l := curveRadius, curveRadius, start.X, start.Y, c.L
	return fmt.Sprintf(
			`<path d="M %.1f %.1f Q %.1f %.1f %.1f %.1f L %.1f %.1f Q %.1f %.1f %.1f %.1f" fill="transparent" stroke="black" stroke-width="%.1f" />`,
			x, y,

			x+r, y,
			x+r, y+ry,

			x+r, y+ry+l,

			x+r, y+2*ry+l,
			x+2*r, y+2*ry+l,

			strokeWidth,
		), Point{X: x + 2*r, Y: y + 2*ry + l}, Box{
			Position: Point{X: x, Y: y},
			Size:     Point{X: 2 * r, Y: 2*ry + l},
		}
}

type CurveS struct {
	L float64
}

func (c CurveS) SVG(start Point) (string, Point, Box) {
	r, ry, x, y, l, ly := curveRadius, -curveRadius, start.X, start.Y, -c.L, c.L
	return fmt.Sprintf(
			`<path d="M %.1f %.1f Q %.1f %.1f %.1f %.1f L %.1f %.1f Q %.1f %.1f %.1f %.1f" fill="transparent" stroke="black" stroke-width="%.1f" />`,
			x, y,

			x+r, y,
			x+r, y+ry,

			x+r, y+ry+l,

			x+r, y+2*ry+l,
			x+2*r, y+2*ry+l,

			strokeWidth,
		), Point{X: x + 2*r, Y: y + 2*ry + l}, Box{
			Position: Point{X: x, Y: y + 2*ry + l},
			Size:     Point{X: 2 * r, Y: 2*r + ly},
		}
}
