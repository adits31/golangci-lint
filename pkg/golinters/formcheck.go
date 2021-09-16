package golinters

import (
	"bytes"
	"go/ast"
	"go/printer"
	"go/token"

	"github.com/adits31/golangci-lint/pkg/golinters/goanalysis"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "addlint",
	Doc:  "reports integer additions",
	Run:  run,
}

var disallowedFunctionTypes = []string{"SelectWorkflow", "SelectWorkflowRow", "SelectTask", "SelectTaskRow"}

const formCheckName = "formcheck"

func NewFormcheck() *goanalysis.Linter {
	return goanalysis.NewLinter(
		formCheckName,
		"checks function and package usages of deprecated form methods complexity",
		[]*analysis.Analyzer{
			Analyzer,
		},
		nil,
	).WithLoadMode(goanalysis.LoadModeTypesInfo)
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			ce, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			fs, ok := ce.Fun.(*ast.SelectorExpr)
			if !ok {
				return true
			}

			functionIsDisallowed := false
			for _, dft := range disallowedFunctionTypes {
				if isIdent(fs.Sel, dft) {
					functionIsDisallowed = true
				}
			}

			if !functionIsDisallowed {
				return true
			}

			pass.Reportf(ce.Pos(), "disallowed form function usage found %q",
				render(pass.Fset, ce))
			return true
		})
	}

	return nil, nil
}

// render returns the pretty-print of the given node
func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}

func isIdent(expr ast.Expr, ident string) bool {
	id, ok := expr.(*ast.Ident)
	return ok && id.Name == ident
}
