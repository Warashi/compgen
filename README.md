# Compgen
Compgen is a gqlgen plugin.
Compgen generates ComplexityRoot of gqlgen.
Generated ComplexityRoot calculates complexity with directive `@complexity(x: number)` and relay paging specification, and configuarable default fallback.

## Example
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
