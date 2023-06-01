package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/plugin/stubgen"
	"github.com/Warashi/compgen"
)

func main() {
	stub := flag.String("stub", "", "name of stub file to generate")
	cfgPath := flag.String("config", "", "path to config file (use default if omitted)")
	defaultComplexity := flag.Int("complexity", 0, "default complexity (use 0 if omitted)")
	flag.Parse()

	log.SetOutput(io.Discard)

	start := graphql.Now()

	var cfg *config.Config
	var err error
	if cfgPath != nil && *cfgPath != "" {
		cfg, err = config.LoadConfig(*cfgPath)
	} else {
		cfg, err = config.LoadConfigFromDefaultLocations()
	}
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	var options []api.Option
	if *stub != "" {
		options = append(options,
			api.AddPlugin(stubgen.New(*stub, "Stub")),
			api.AddPlugin(compgen.New(compgen.WithDefaultComplexity(*defaultComplexity))),
		)
	}

	err = api.Generate(cfg, options...)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}

	fmt.Printf("Generated %s in %4.2fs\n", cfg.Exec.ImportPath(), time.Since(start).Seconds())
}
