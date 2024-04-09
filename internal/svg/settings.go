package svg

var (
	fontSize       = 8.0
	pointRadius    = 1.0
	interNodeSpace = 10.0
	curveRadius    = 5.0
	strokeWidth    = 1.0
)

// SetFontSize sets the font size of the text.
func SetFontSize(size float64) {
	fontSize = size
}

// SetPointRadius sets the pointRadius of a point.
func SetPointRadius(r float64) {
	pointRadius = r
}

// SetInterNodeSpace sets the interNodeSpace of the nodes.
func SetInterNodeSpace(s float64) {
	interNodeSpace = s
}

// SetCurveRadius sets the curveRadius of the curve.
func SetCurveRadius(r float64) {
	curveRadius = r
}

// SetStrokeWidth sets the stroke width of the curve.
func SetStrokeWidth(w float64) {
	strokeWidth = w
}
