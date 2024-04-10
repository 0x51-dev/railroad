package svg

import (
	"fmt"
	"os"
	"testing"
)

func TestCurve(t *testing.T) {
	if os.Getenv("GENERATE_SVG") != "" {
		t.Run("lines", func(t *testing.T) {
			for _, size := range []float64{1, 10} {
				var combinedSVG string
				start := Point{X: 0, Y: 0}
				combinedBox := Box{Position: start}
				for _, e := range []SVGable{
					Line{RelativeEnd: Point{Y: -10 * size}},
					Line{RelativeEnd: Point{X: 2 * size, Y: -5 * size}},
					Line{RelativeEnd: Point{X: 5 * size, Y: -5 * size}},
					Line{RelativeEnd: Point{X: 5 * size, Y: -2 * size}},
					Line{RelativeEnd: Point{X: 10 * size}},
					Line{RelativeEnd: Point{X: 5 * size, Y: 2 * size}},
					Line{RelativeEnd: Point{X: 5 * size, Y: 5 * size}},
					Line{RelativeEnd: Point{X: 2 * size, Y: 5 * size}},
					Line{RelativeEnd: Point{Y: 10 * size}},
					Line{RelativeEnd: Point{X: -2 * size, Y: 5 * size}},
					Line{RelativeEnd: Point{X: -5 * size, Y: 5 * size}},
					Line{RelativeEnd: Point{X: -5 * size, Y: 2 * size}},
					Line{RelativeEnd: Point{X: -10 * size}},
					Line{RelativeEnd: Point{X: -5 * size, Y: -2 * size}},
					Line{RelativeEnd: Point{X: -5 * size, Y: -5 * size}},
					Line{RelativeEnd: Point{X: -2 * size, Y: -5 * size}},
				} {
					svg, end, box := DebugSVG(e, start)
					// svg, end, box := e.SVG(start)
					combinedSVG += svg
					start = end
					combinedBox = combinedBox.Combine(box)
				}
				_ = os.WriteFile(fmt.Sprintf("testdata/lines%.0f.svg", size), []byte(
					fmt.Sprintf(
						`<svg viewBox="%.1f %.1f %.1f %.1f" xmlns="http://www.w3.org/2000/svg">%s
</svg>`,
						combinedBox.Position.X-10, combinedBox.Position.Y-10,
						combinedBox.Size.X+20, combinedBox.Size.Y+20,
						combinedSVG,
					),
				), 0644)
			}
		})

		t.Run("combined", func(t *testing.T) {
			var combinedSVG string
			start := Point{X: 0, Y: 0}
			combinedBox := Box{Position: start}
			for _, e := range []SVGable{
				Start{},
				Line{RelativeEnd: Point{X: 10, Y: -10}},
				Line{RelativeEnd: Point{X: 10, Y: 10}},
				Bypass{
					Element: TextBox{Text: "Hello,"},
				},
				Line{RelativeEnd: Point{X: 10, Y: -10}},
				TextBox{Text: "World!", RoundedBorders: true},
				Line{RelativeEnd: Point{X: 10}},
				Line{RelativeEnd: Point{X: 10, Y: 10}},
				End{},
			} {
				svg, end, box := DebugSVG(e, start)
				// svg, end, box := e.SVG(start)
				combinedSVG += svg
				start = end
				combinedBox = combinedBox.Combine(box)
			}
			_ = os.WriteFile("testdata/total.svg", []byte(
				fmt.Sprintf(
					`<svg viewBox="%.1f %.1f %.1f %.1f" xmlns="http://www.w3.org/2000/svg">%s
</svg>`,
					combinedBox.Position.X-10, combinedBox.Position.Y-10,
					combinedBox.Size.X+20, combinedBox.Size.Y+20,
					combinedSVG,
				),
			), 0644)
		})
	}
}
