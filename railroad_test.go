package railroad_test

import (
	"github.com/0x51-dev/railroad"
	"os"
	"testing"
)

func TestRailroad(t *testing.T) {
	for _, r := range []railroad.Rule{
		{
			Name: "explain-and",
			Expr: railroad.AndExpression{
				Expr: []railroad.Expression{
					railroad.T("A"),
					railroad.T("B"),
					railroad.T("C"),
				},
			},
		},
		{
			Name: "explain-optional",
			Expr: railroad.AndExpression{
				Expr: []railroad.Expression{
					railroad.OptionalExpression{
						Expr: railroad.T("A"),
					},
					railroad.T("B"),
				},
			},
		},
		{
			Name: "explain-or",
			Expr: railroad.OrExpression{
				Expr: []railroad.Expression{
					railroad.T("A"),
					railroad.T("B"),
					railroad.T("C"),
				},
			},
		},
		{
			Name: "explain-optional-or",
			Expr: railroad.OptionalExpression{
				Expr: railroad.OrExpression{
					Expr: []railroad.Expression{
						railroad.T("A"),
						railroad.T("B"),
						railroad.T("C"),
					},
				},
			},
		},

		{
			Name: "example1",
			Expr: railroad.OptionalExpression{
				Expr: railroad.AndExpression{
					Expr: []railroad.Expression{
						railroad.OptionalExpression{
							Expr: railroad.T("Hello,"),
						},
						railroad.T("World!"),
					},
				},
			},
		},
		{
			Name: "optional3",
			Expr: railroad.OptionalExpression{
				Expr: railroad.OptionalExpression{
					Expr: railroad.OptionalExpression{
						Expr: railroad.T("?"),
					},
				},
			},
		},
	} {
		t.Run(r.Name, func(t *testing.T) {
			f, err := os.OpenFile("testdata/"+r.Name+".svg", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
			if err != nil {
				t.Fatal(err)
			}

			_, err = f.WriteString(railroad.GenerateSVG(r))
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}
