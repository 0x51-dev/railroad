package svg

import "fmt"

type Line struct {
	RelativeEnd Point
}

func (e Line) SVG(start Point) (string, Point, Box) {
	end, box := Point{X: start.X + e.RelativeEnd.X, Y: start.Y + e.RelativeEnd.Y}, NewBox(start, e.RelativeEnd)
	next := end

	// The line is horizontal/vertical.
	if e.RelativeEnd.X == 0 || e.RelativeEnd.Y == 0 {
		return fmt.Sprintf(
			`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="black" stroke-linecap="round" />`,
			start.X, start.Y,
			start.X+e.RelativeEnd.X, start.Y+e.RelativeEnd.Y,
		), next, box
	}

	crX, crY := curveRadius, curveRadius
	if -2*curveRadius < e.RelativeEnd.X && e.RelativeEnd.X < 2*curveRadius {
		crX = 0
	}
	if -2*curveRadius < e.RelativeEnd.Y && e.RelativeEnd.Y < 2*curveRadius {
		crY = 0
	}

	switch {
	// -> NW
	case e.RelativeEnd.X < 0 && e.RelativeEnd.Y < 0:
		start, end = end, start
		e.RelativeEnd.X = -e.RelativeEnd.X
		e.RelativeEnd.Y = -e.RelativeEnd.Y
		fallthrough
	// -> SE
	case 0 < e.RelativeEnd.X && 0 < e.RelativeEnd.Y:
		return fmt.Sprintf(
			`<path d="M %.1f %.1f Q %0.1f %0.1f %0.1f %0.1f L %0.1f %0.1f Q %0.1f %0.1f %0.1f %0.1f L %0.1f %0.1f" fill="transparent" stroke="black" stroke-linecap="round" />`,
			start.X, start.Y, // Start top-left.

			start.X+crY, start.Y, // Curve down.
			start.X+crY, start.Y+crY,

			start.X+crY, start.Y+e.RelativeEnd.Y-crX, // Straight down.

			start.X+crY, end.Y, // Curve right.
			start.X+crY+crX, end.Y,

			end.X, end.Y, // End bottom-right.
		), next, box
	// -> SW
	case e.RelativeEnd.X < 0 && 0 < e.RelativeEnd.Y:
		start, end = end, start
		e.RelativeEnd.X = -e.RelativeEnd.X
		e.RelativeEnd.Y = -e.RelativeEnd.Y
		fallthrough
	// -> NE
	case 0 < e.RelativeEnd.X && e.RelativeEnd.Y < 0:
		return fmt.Sprintf(
			`<path d="M %.1f %.1f Q %0.1f %0.1f %0.1f %0.1f L %0.1f %0.1f Q %0.1f %0.1f %0.1f %0.1f L %0.1f %0.1f" fill="transparent" stroke="black" stroke-linecap="round" />`,
			start.X, start.Y, // Start bottom-left.

			start.X+crX, start.Y, // Curve up.
			start.X+crX, start.Y-crX,

			start.X+crX, start.Y+e.RelativeEnd.Y+crY, // Straight up.

			start.X+crY, start.Y+e.RelativeEnd.Y, // Curve right.
			start.X+2*crY, start.Y+e.RelativeEnd.Y,

			end.X, end.Y, // End top-right.
		), next, box
	}
	return `<g class="rr-line"></g>`, next, box
}
