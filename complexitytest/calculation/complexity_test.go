//go:generate go run ../cmd/gqlgen.go -config gqlgen-default0.yaml -stub gql0/stub.go -complexity 0
//go:generate go run ../cmd/gqlgen.go -config gqlgen-default1.yaml -stub gql1/stub.go -complexity 1
//go:generate go run ../cmd/gqlgen.go -config gqlgen-default2.yaml -stub gql2/stub.go -complexity 2
package complexity_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/99designs/gqlgen/client"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/Warashi/compgen/complexitytest/calculation/gql0"
	"github.com/Warashi/compgen/complexitytest/calculation/gql1"
	"github.com/Warashi/compgen/complexitytest/calculation/gql2"
	"github.com/stretchr/testify/assert"
)

func TestComplexityDefault0(t *testing.T) {
	resolvers := &gql0.Stub{}

	cfg := gql0.Config{
		Resolvers:  resolvers,
		Complexity: gql0.ComplexityFuncs,
	}

	srv := handler.NewDefaultServer(gql0.NewExecutableSchema(cfg))
	srv.Use(extension.FixedComplexityLimit(-1))

	c := client.New(srv)

	tests := []struct {
		query      string
		complexity int
	}{
		{query: `query { foo(a: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 10},
		{query: `query { foo(a: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 20},
		{query: `query { foo(a: 5, b: 1) { edges { node { bar } }  } }`, complexity: 25},
		{query: `query { foo(a: 10, b: 1) { edges { node { bar } } } }`, complexity: 50},
		{query: `query { foo(b: 5) { pageInfo { hasNextPage } } }`, complexity: 10},
		{query: `query { foo(b: 10) { pageInfo { hasNextPage } } }`, complexity: 20},
		{query: `query { foo(b: 5) { edges { node { bar } }  } }`, complexity: 25},
		{query: `query { foo(b: 10) { edges { node { bar } } } }`, complexity: 50},
		{query: `query { foo(c: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 2},
		{query: `query { foo(c: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 2},
		{query: `query { foo(c: 5, b: 1) { edges { node { bar } }  } }`, complexity: 5},
		{query: `query { foo(c: 10, b: 1) { edges { node { bar } } } }`, complexity: 5},
		{query: `query { bar(a: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 0},
		{query: `query { bar(a: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 0},
		{query: `query { bar(a: 5, b: 1) { edges { node { bar } }  } }`, complexity: 15},
		{query: `query { bar(a: 10, b: 1) { edges { node { bar } } } }`, complexity: 30},
		{query: `query { bar(b: 5) { pageInfo { hasNextPage } } }`, complexity: 0},
		{query: `query { bar(b: 10) { pageInfo { hasNextPage } } }`, complexity: 0},
		{query: `query { bar(b: 5) { edges { node { bar } }  } }`, complexity: 15},
		{query: `query { bar(b: 10) { edges { node { bar } } } }`, complexity: 30},
		{query: `query { bar(c: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 0},
		{query: `query { bar(c: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 0},
		{query: `query { bar(c: 5, b: 1) { edges { node { bar } }  } }`, complexity: 3},
		{query: `query { bar(c: 10, b: 1) { edges { node { bar } } } }`, complexity: 3},
    {query: `query { baz(ids: ["a"]) { pageInfo { hasNextPage } } }`, complexity: 0},
		{query: `query { baz(ids: ["a","b"]) { pageInfo { hasNextPage } } }`, complexity: 0},
		{query: `query { baz(ids: ["a"]) { edges { node { bar } }  } }`, complexity: 3},
		{query: `query { baz(ids: ["a","b"]) { edges { node { bar } } } }`, complexity: 6},
    {query: `query { bazz(ids: ["a"]) { pageInfo { hasNextPage } } }`, complexity: 0},
		{query: `query { bazz(ids: ["a","b"]) { pageInfo { hasNextPage } } }`, complexity: 0},
		{query: `query { bazz(ids: ["a"]) { edges { node { bar } }  } }`, complexity: 3},
		{query: `query { bazz(ids: ["a","b"]) { edges { node { bar } } } }`, complexity: 6},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			err := c.Post(tt.query, new(struct{}))
			assert.EqualError(t, err, fmt.Sprintf(`[{"message":"operation has complexity %d, which exceeds the limit of -1","extensions":{"code":"COMPLEXITY_LIMIT_EXCEEDED"}}]`, tt.complexity))
		})
	}
}

func TestComplexityDefault1(t *testing.T) {
	resolvers := &gql1.Stub{}

	cfg := gql1.Config{
		Resolvers:  resolvers,
		Complexity: gql1.ComplexityFuncs,
	}

	srv := handler.NewDefaultServer(gql1.NewExecutableSchema(cfg))
	srv.Use(extension.FixedComplexityLimit(-1))

	c := client.New(srv)

	tests := []struct {
		query      string
		complexity int
	}{
		{query: `query { foo(a: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 20},
		{query: `query { foo(a: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 40},
		{query: `query { foo(a: 5, b: 1) { edges { node { bar } }  } }`, complexity: 35},
		{query: `query { foo(a: 10, b: 1) { edges { node { bar } } } }`, complexity: 70},
		{query: `query { foo(b: 5) { pageInfo { hasNextPage } } }`, complexity: 20},
		{query: `query { foo(b: 10) { pageInfo { hasNextPage } } }`, complexity: 40},
		{query: `query { foo(b: 5) { edges { node { bar } }  } }`, complexity: 35},
		{query: `query { foo(b: 10) { edges { node { bar } } } }`, complexity: 70},
		{query: `query { foo(c: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 4},
		{query: `query { foo(c: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 4},
		{query: `query { foo(c: 5, b: 1) { edges { node { bar } }  } }`, complexity: 7},
		{query: `query { foo(c: 10, b: 1) { edges { node { bar } } } }`, complexity: 7},
		{query: `query { bar(a: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 15},
		{query: `query { bar(a: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 30},
		{query: `query { bar(a: 5, b: 1) { edges { node { bar } }  } }`, complexity: 30},
		{query: `query { bar(a: 10, b: 1) { edges { node { bar } } } }`, complexity: 60},
		{query: `query { bar(b: 5) { pageInfo { hasNextPage } } }`, complexity: 15},
		{query: `query { bar(b: 10) { pageInfo { hasNextPage } } }`, complexity: 30},
		{query: `query { bar(b: 5) { edges { node { bar } }  } }`, complexity: 30},
		{query: `query { bar(b: 10) { edges { node { bar } } } }`, complexity: 60},
		{query: `query { bar(c: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 3},
		{query: `query { bar(c: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 3},
		{query: `query { bar(c: 5, b: 1) { edges { node { bar } }  } }`, complexity: 6},
		{query: `query { bar(c: 10, b: 1) { edges { node { bar } } } }`, complexity: 6},
    {query: `query { baz(ids: ["a"]) { pageInfo { hasNextPage } } }`, complexity: 3},
		{query: `query { baz(ids: ["a","b"]) { pageInfo { hasNextPage } } }`, complexity: 6},
		{query: `query { baz(ids: ["a"]) { edges { node { bar } }  } }`, complexity: 6},
		{query: `query { baz(ids: ["a","b"]) { edges { node { bar } } } }`, complexity: 12},
    {query: `query { bazz(ids: ["a"]) { pageInfo { hasNextPage } } }`, complexity: 3},
		{query: `query { bazz(ids: ["a","b"]) { pageInfo { hasNextPage } } }`, complexity: 6},
		{query: `query { bazz(ids: ["a"]) { edges { node { bar } }  } }`, complexity: 6},
		{query: `query { bazz(ids: ["a","b"]) { edges { node { bar } } } }`, complexity: 12},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			err := c.Post(tt.query, new(struct{}))
			assert.EqualError(t, err, fmt.Sprintf(`[{"message":"operation has complexity %d, which exceeds the limit of -1","extensions":{"code":"COMPLEXITY_LIMIT_EXCEEDED"}}]`, tt.complexity))
		})
	}
}

func TestComplexityDefault2(t *testing.T) {
	resolvers := &gql2.Stub{}

	cfg := gql2.Config{
		Resolvers:  resolvers,
		Complexity: gql2.ComplexityFuncs,
	}

	srv := handler.NewDefaultServer(gql2.NewExecutableSchema(cfg))
	srv.Use(extension.FixedComplexityLimit(-1))

	c := client.New(srv)

	tests := []struct {
		query      string
		complexity int
	}{
		{query: `query { foo(a: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 30},
		{query: `query { foo(a: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 60},
		{query: `query { foo(a: 5, b: 1) { edges { node { bar } }  } }`, complexity: 45},
		{query: `query { foo(a: 10, b: 1) { edges { node { bar } } } }`, complexity: 90},
		{query: `query { foo(b: 5) { pageInfo { hasNextPage } } }`, complexity: 30},
		{query: `query { foo(b: 10) { pageInfo { hasNextPage } } }`, complexity: 60},
		{query: `query { foo(b: 5) { edges { node { bar } }  } }`, complexity: 45},
		{query: `query { foo(b: 10) { edges { node { bar } } } }`, complexity: 90},
		{query: `query { foo(c: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 6},
		{query: `query { foo(c: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 6},
		{query: `query { foo(c: 5, b: 1) { edges { node { bar } }  } }`, complexity: 9},
		{query: `query { foo(c: 10, b: 1) { edges { node { bar } } } }`, complexity: 9},
		{query: `query { bar(a: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 30},
		{query: `query { bar(a: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 60},
		{query: `query { bar(a: 5, b: 1) { edges { node { bar } }  } }`, complexity: 45},
		{query: `query { bar(a: 10, b: 1) { edges { node { bar } } } }`, complexity: 90},
		{query: `query { bar(b: 5) { pageInfo { hasNextPage } } }`, complexity: 30},
		{query: `query { bar(b: 10) { pageInfo { hasNextPage } } }`, complexity: 60},
		{query: `query { bar(b: 5) { edges { node { bar } }  } }`, complexity: 45},
		{query: `query { bar(b: 10) { edges { node { bar } } } }`, complexity: 90},
		{query: `query { bar(c: 5, b: 1) { pageInfo { hasNextPage } } }`, complexity: 6},
		{query: `query { bar(c: 10, b: 1) { pageInfo { hasNextPage } } }`, complexity: 6},
		{query: `query { bar(c: 5, b: 1) { edges { node { bar } }  } }`, complexity: 9},
		{query: `query { bar(c: 10, b: 1) { edges { node { bar } } } }`, complexity: 9},
    {query: `query { baz(ids: ["a"]) { pageInfo { hasNextPage } } }`, complexity: 6},
		{query: `query { baz(ids: ["a","b"]) { pageInfo { hasNextPage } } }`, complexity: 12},
		{query: `query { baz(ids: ["a"]) { edges { node { bar } }  } }`, complexity: 9},
		{query: `query { baz(ids: ["a","b"]) { edges { node { bar } } } }`, complexity: 18},
    {query: `query { bazz(ids: ["a"]) { pageInfo { hasNextPage } } }`, complexity: 6},
		{query: `query { bazz(ids: ["a","b"]) { pageInfo { hasNextPage } } }`, complexity: 12},
		{query: `query { bazz(ids: ["a"]) { edges { node { bar } }  } }`, complexity: 9},
		{query: `query { bazz(ids: ["a","b"]) { edges { node { bar } } } }`, complexity: 18},
	}

	for i, tt := range tests {
		i, tt := i, tt
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			t.Parallel()

			err := c.Post(tt.query, new(struct{}))
			assert.EqualError(t, err, fmt.Sprintf(`[{"message":"operation has complexity %d, which exceeds the limit of -1","extensions":{"code":"COMPLEXITY_LIMIT_EXCEEDED"}}]`, tt.complexity))
		})
	}
}
