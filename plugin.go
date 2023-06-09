package compgen

import (
	_ "embed"
	"errors"
	"fmt"
	"path/filepath"
	"text/template"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
	"github.com/vektah/gqlparser/v2/ast"
)

//go:embed templates/compgen.gotpl
var tmpl string

var (
	ErrMulFieldIsWrongType = errors.New("mul field is neither Int nor List")
	ErrMulFieldIsNotExist  = errors.New("mul field is not exist")
)

var (
	_ plugin.Plugin        = (*Plugin)(nil)
	_ plugin.CodeGenerator = (*Plugin)(nil)
	_ plugin.ConfigMutator = (*Plugin)(nil)
)

type Plugin struct {
	Filename          string
	DefaultComplexity int
}

type Option func(*Plugin)

// WithDefaultComplexity sets default complexity. default is `1`.
func WithDefaultComplexity(x int) Option {
	return func(p *Plugin) {
		p.DefaultComplexity = x
	}
}

// WithFilename sets generated filename. default is `compgen.go`.
func WithFilename(fn string) Option {
	return func(p *Plugin) {
		p.Filename = fn
	}
}

func New(opts ...Option) *Plugin {
	p := Plugin{
		Filename:          "compgen.go",
		DefaultComplexity: 1,
	}

	for _, o := range opts {
		o(&p)
	}

	return &p
}

// Name implements plugin.Plugin
func (*Plugin) Name() string {
	return "compgen"
}

type Inputs struct {
	Plugin     *Plugin
	GQLGenData *codegen.Data
}

func IsInt(typ *ast.Type) bool {
	return typ.NamedType == "Int"
}

func IsList(typ *ast.Type) bool {
	return typ.NamedType == ""
}

// GenerateCode implements plugin.CodeGenerator
func (p *Plugin) GenerateCode(cfg *codegen.Data) error {
	for _, object := range cfg.Objects {
		for _, fields := range object.UniqueFields() {
			field := fields[0]
			directive := field.FieldDefinition.Directives.ForName("complexity")
			if directive == nil {
				continue
			}
			mul := directive.Arguments.ForName("mul")
			if mul == nil {
				continue
			}

			for _, child := range mul.Value.Children {
				name := child.Value.Raw
				arg := field.Arguments.ForName(name)
				if arg == nil {
					return fmt.Errorf("argument `%s` is used by @complexity's mul argument, but its not exist: %w", name, ErrMulFieldIsNotExist)
				}
				if !IsInt(arg.Type) && !IsList(arg.Type) {
					return fmt.Errorf("argument `%s` is used by @complexity's mul argument, but its type is neither Int nor List: %w", name, ErrMulFieldIsWrongType)
				}
			}
		}
	}
	return templates.Render(templates.Options{
		PackageName: cfg.Config.Exec.Package,
		Template:    tmpl,
		Filename:    filepath.Join(filepath.Dir(cfg.Config.Exec.Filename), p.Filename),
		Funcs: template.FuncMap{
			"IsInt":  IsInt,
			"IsList": IsList,
		},
		GeneratedHeader: true,
		Packages:        cfg.Config.Packages,
		Data: Inputs{
			Plugin:     p,
			GQLGenData: cfg,
		},
	})
}

// MutateConfig implements plugin.ConfigMutator
func (*Plugin) MutateConfig(cfg *config.Config) error {
	cfg.Directives["complexity"] = config.DirectiveConfig{
		SkipRuntime: true,
	}
	return nil
}
