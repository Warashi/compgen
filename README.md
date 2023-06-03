# Compgen
Compgen is a gqlgen plugin designed to simplify the generation of ComplexityRoot for gqlgen. The generated ComplexityRoot calculates complexity using the `@complexity(x: number, mul: [String!])` directive, and a configurable default fallback.

## Usage
1. Create a main.go file to use this plugin with gqlgen. Here's an example:
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
2. Use the generated ComplexityFunc as ComplexityRoot. For example:
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
    ): FooConnection! @complexity(x: 2, mul: ["first", "last"])
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
Execute the following query:
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
The complexity will be calculated as (3 + default + default + 2) * 5.

# Additional tools
A linter [relaycompmul](./linter/relaycompmul) is useful when compgen with [relay cursor connections spec](https://relay.dev/graphql/connections.htm).
This linter outout errors when fields that follow relay cursor connections specificaiton and do not have `@complexity` directive or lack `mul` arguments of `@complexity`.
