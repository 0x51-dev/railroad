package railroad

import (
	"fmt"
	"github.com/0x51-dev/railroad/internal/svg"
)

func GenerateSVG(r Rule) string {
	margin := 10.0
	ruleSVG, _, box := r.SVG(svg.Point{})
	return fmt.Sprintf(
		`<svg viewBox="%.1f %.1f %.1f %.1f" xmlns="http://www.w3.org/2000/svg">%s
</svg>`,
		box.Position.X-margin, box.Position.Y-margin,
		box.Size.X+2*margin, box.Size.Y+2*margin,
		ruleSVG,
	)
}

type AndExpression struct {
	Expr []Expression
}

func (AndExpression) expression() {}

func (a AndExpression) SVG(start svg.Point) (string, svg.Point, svg.Box) {
	var combinedSVG string
	startBox := svg.Box{Position: start}
	for idx, e := range a.Expr {
		exprSVG, exprEnd, exprBox := e.SVG(start)
		combinedSVG += exprSVG
		start = exprEnd
		startBox = startBox.Combine(exprBox)
		if idx < len(a.Expr)-1 {
			combinedSVG += fmt.Sprintf(
				`<line x1="%.1f" y1="%.1f" x2="%.1f" y2="%.1f" stroke="black" />`,
				start.X, start.Y, start.X+10, start.Y,
			)
			start.X += 10
		}
	}
	return fmt.Sprintf(`
<g class="rr-and">%s
</g>`,
		combinedSVG,
	), start, startBox
}

type Expression interface {
	svg.SVGable

	expression()
}

type OrExpression struct {
	Expr []Expression
}

func (OrExpression) expression() {}

func (o OrExpression) SVG(start svg.Point) (string, svg.Point, svg.Box) {
	var combinedSVG string
	var end svg.Point
	var combinedBox svg.Box
	var height float64
	for _, e := range o.Expr {
		startLineSVG, startLineEnd, startLineBox := svg.Line{RelativeEnd: svg.Point{X: 10, Y: height}}.SVG(start)
		wrappedSVG, wrappedEnd, wrappedBox := e.SVG(startLineEnd)
		endLineSVG, endLineEnd, endLineBox := svg.Line{RelativeEnd: svg.Point{X: 10, Y: -height}}.SVG(wrappedEnd)
		combinedSVG += startLineSVG + wrappedSVG + endLineSVG
		end = endLineEnd
		combinedBox = combinedBox.Combine(startLineBox).Combine(wrappedBox).Combine(endLineBox)
		height += wrappedBox.Size.Y + svg.GetCurveRadius()
	}
	return combinedSVG, end, combinedBox
}

type OptionalExpression struct {
	Expr Expression
}

func (OptionalExpression) expression() {}

func (o OptionalExpression) SVG(start svg.Point) (string, svg.Point, svg.Box) {
	return svg.Bypass{Element: o.Expr}.SVG(start)
}

type Rule struct {
	Name string
	Expr Expression
}

func (r Rule) SVG(start svg.Point) (string, svg.Point, svg.Box) {
	startSVG, startEnd, startBox := svg.Start{}.SVG(start)
	ruleSVG, ruleEnd, ruleBox := r.Expr.SVG(startEnd)
	endSVG, endEnd, endBox := svg.End{}.SVG(ruleEnd)
	return fmt.Sprintf(`
<g class="rr-rule" id="%s">%s
%s
%s
</g>`,
		r.Name, startSVG, ruleSVG, endSVG,
	), endEnd, startBox.Combine(ruleBox).Combine(endBox)
}

type Term struct {
	Value string
}

func T(value string) Term {
	return Term{Value: value}
}

func (Term) expression() {}

func (t Term) SVG(start svg.Point) (string, svg.Point, svg.Box) {
	termSVG, termEnd, termBox := svg.TextBox{Text: t.Value}.SVG(start)
	return fmt.Sprintf(`
<g class="rr-term">%s
</g>`,
		termSVG,
	), termEnd, termBox
}
