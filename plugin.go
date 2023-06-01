package compgen

import (
	_ "embed"
	"path/filepath"

	"github.com/99designs/gqlgen/codegen"
	"github.com/99designs/gqlgen/codegen/config"
	_ "github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/codegen/templates"
	"github.com/99designs/gqlgen/plugin"
)

//go:embed templates/compgen.gotpl
var tmpl string

var (
	_ plugin.Plugin        = (*Plugin)(nil)
	_ plugin.CodeGenerator = (*Plugin)(nil)
	_ plugin.ConfigMutator = (*Plugin)(nil)
)

type Plugin struct {
	DefaultComplexity int
}

type Option func(*Plugin)

func WithDefaultComplexity(x int) Option {
	return func(p *Plugin) {
		p.DefaultComplexity = x
	}
}

func New(opts ...Option) *Plugin {
	p := Plugin{
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

// GenerateCode implements plugin.CodeGenerator
func (p *Plugin) GenerateCode(cfg *codegen.Data) error {
	return templates.Render(templates.Options{
		PackageName:     cfg.Config.Exec.Package,
		Template:        tmpl,
		Filename:        filepath.Join(filepath.Dir(cfg.Config.Exec.Filename), "compgen.go"),
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
