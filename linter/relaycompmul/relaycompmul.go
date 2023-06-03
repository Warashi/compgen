package relaycompmul

import (
	"fmt"
	"strings"

	"github.com/gqlgo/gqlanalysis"
	"github.com/vektah/gqlparser/v2/ast"
)

var Analyzer = &gqlanalysis.Analyzer{
	Name: "relaycompmul",
	Doc:  "relaycompmul finds a relay cursor connection fields which not have @complexity directive or its mul argument.",
	Run:  run,
}

func isRelayCursorConnection(field *ast.FieldDefinition) bool {
	if !strings.HasSuffix(field.Type.Name(), "Connection") {
		return false
	}
	for _, argn := range []string{"after", "before"} {
		if field.Arguments.ForName(argn) == nil {
			return false
		}
	}
	for _, argn := range []string{"first", "last"} {
		arg := field.Arguments.ForName(argn)
		if arg == nil {
			return false
		}
		if arg.Type.Name() != "Int" {
			return false
		}
	}
	return true
}

func hasComplexity(field *ast.FieldDefinition) bool {
	return field.Directives.ForName("complexity") != nil
}

func hasMul(fn string, field *ast.FieldDefinition) bool {
	directive := field.Directives.ForName("complexity")
	if directive == nil {
		return false
	}

	mul := directive.Arguments.ForName("mul")
	if mul == nil {
		return false
	}

	set := make(map[string]struct{}, 10)
	for _, child := range mul.Value.Children {
		name := child.Value.Raw
		set[name] = struct{}{}
	}

	_, ok := set[fn]
	return ok
}

func run(pass *gqlanalysis.Pass) (interface{}, error) {
	if pass == nil || pass.Schema == nil || pass.Schema.Query == nil {
		return nil, nil
	}
	for _, field := range pass.Schema.Query.Fields {
		if !isRelayCursorConnection(field) {
			continue
		}
		if !hasComplexity(field) {
			pass.Report(&gqlanalysis.Diagnostic{
				Pos:     field.Type.Position,
				Message: fmt.Sprintf("field [%s] is Relay Cursor Connections field, but @complexity directive is not found", field.Name),
			})
			continue
		}

		first := hasMul("first", field)
		last := hasMul("last", field)

		const message = "field [%s] is Relay Cursor Connections field, but [%s] argument is not included in [mul] of @complexity"
		switch {
		case !first && !last:
			pass.Report(&gqlanalysis.Diagnostic{
				Pos:     field.Type.Position,
				Message: fmt.Sprintf(message, field.Name, "first, last"),
			})
		case !first:
			pass.Report(&gqlanalysis.Diagnostic{
				Pos:     field.Type.Position,
				Message: fmt.Sprintf(message, field.Name, "first"),
			})
		case !last:
			pass.Report(&gqlanalysis.Diagnostic{
				Pos:     field.Type.Position,
				Message: fmt.Sprintf(message, field.Name, "last"),
			})
		default:
			// no diagnostic
		}
	}
	return nil, nil
}
