# Compgen
Compgen is a gqlgen plugin.
Compgen generates ComplexityRoot of gqlgen.
Generated ComplexityRoot calculates complexity with directive `@complexity(x: number)` and relay paging specification, and configuarable default fallback.

## Usage
First, write main.go as gqlgen to use this plugin. For example:
```go
package main

import (
	"fmt"
	"os"

	"github.com/99designs/gqlgen/api"
	"github.com/99designs/gqlgen/codegen/config"
	"github.com/Warashi/compgen"
)

func main() {
	cfg, err := config.LoadConfigFromDefaultLocations()
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load config", err.Error())
		os.Exit(2)
	}

	if err := api.Generate(cfg, api.AddPlugin(compgen.New(compgen.WithDefaultComplexity(1)))); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(3)
	}
}
```

Then, use generated ComplexityFunc as ComplexityRoot. For example:
```go
cfg := gql.Config{
  Resolvers: resolvers, 
  Complexity: gql.ComplexityFuncs,
}

srv := handler.NewDefaultServer(gql.NewExecutableSchema(cfg))
srv.Use(extension.FixedComplexityLimit(1000))
```

## Calculation Example
### Schema
```graphql
type Query {
    foo(
      after: String,
      first: Int,
      before: String,
      last: Int,
    ): FooConnection! @complexity(x: 2)
}

type PageInfo {
    hasNextPage: Boolean!
    hasPreviousPage: Boolean!
    startCursor: String
    endCursor: String
}

type FooConnection {
    edges: [FooEdge!]!
    pageInfo: PageInfo!
}

type FooEdge {
    cursor: String!
    node: Foo!
}

type Foo {
  bar: String! @complexity(x: 3)
}
```

execute below query 
```graphql
query {
  foo(first: 5) {
    edges {
      node {
        bar
      }
    } 
  }
}
```

complexity becomes `(3 + default + default + 2) * 5`
