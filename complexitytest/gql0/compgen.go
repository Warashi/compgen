// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package gql0

var ComplexityFuncs ComplexityRoot = struct {
	Foo struct {
		Bar func(childComplexity int) int
	}

	FooConnection struct {
		Edges    func(childComplexity int) int
		PageInfo func(childComplexity int) int
	}

	FooEdge struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}

	PageInfo struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}

	Query struct {
		Foo func(childComplexity int, after *string, first *int, before *string, last *int) int
	}
}{

	Foo: struct {
		Bar func(childComplexity int) int
	}{
		Bar: func(childComplexity int) int {
			var complexity int

			complexity = childComplexity + 3

			return complexity
		},
	},

	FooConnection: struct {
		Edges    func(childComplexity int) int
		PageInfo func(childComplexity int) int
	}{
		Edges: func(childComplexity int) int {
			var complexity int

			complexity = childComplexity + 0

			return complexity
		},
		PageInfo: func(childComplexity int) int {
			var complexity int

			complexity = childComplexity + 0

			return complexity
		},
	},

	FooEdge: struct {
		Cursor func(childComplexity int) int
		Node   func(childComplexity int) int
	}{
		Cursor: func(childComplexity int) int {
			var complexity int

			complexity = childComplexity + 0

			return complexity
		},
		Node: func(childComplexity int) int {
			var complexity int

			complexity = childComplexity + 0

			return complexity
		},
	},

	PageInfo: struct {
		EndCursor       func(childComplexity int) int
		HasNextPage     func(childComplexity int) int
		HasPreviousPage func(childComplexity int) int
		StartCursor     func(childComplexity int) int
	}{
		EndCursor: func(childComplexity int) int {
			var complexity int

			complexity = childComplexity + 0

			return complexity
		},
		HasNextPage: func(childComplexity int) int {
			var complexity int

			complexity = childComplexity + 0

			return complexity
		},
		HasPreviousPage: func(childComplexity int) int {
			var complexity int

			complexity = childComplexity + 0

			return complexity
		},
		StartCursor: func(childComplexity int) int {
			var complexity int

			complexity = childComplexity + 0

			return complexity
		},
	},

	Query: struct {
		Foo func(childComplexity int, after *string, first *int, before *string, last *int) int
	}{
		Foo: func(childComplexity int, after *string, first *int, before *string, last *int) int {
			var complexity int

			complexity = childComplexity + 2

			if first != nil {
				complexity *= *first
			}

			if last != nil {
				complexity *= *last
			}

			return complexity
		},
	},
}
