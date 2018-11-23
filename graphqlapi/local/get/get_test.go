// These tests verify that the parameters to the resolver are properly extracted from a GraphQL query.
package local_get

import (
	"github.com/creativesoftwarefdn/weaviate/database/schema"
	"github.com/creativesoftwarefdn/weaviate/database/schema/kind"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/common_resolver"
	"github.com/creativesoftwarefdn/weaviate/graphqlapi/local/common_filters"
	test_helper "github.com/creativesoftwarefdn/weaviate/graphqlapi/test/helper"
	"testing"
)

func TestSimpleFieldParamsOK(t *testing.T) {
	t.Parallel()

	resolver := newMockResolver()

	expectedParams := &LocalGetClassParams{
		Kind:       kind.ACTION_KIND,
		ClassName:  "SomeAction",
		Properties: []SelectProperty{{Name: "intField"}},
	}

	resolver.On("LocalGetClass", expectedParams).
		Return(test_helper.EmptyListThunk(), nil).Once()

	resolver.AssertResolve(t, "{ Get { Actions { SomeAction { intField } } } }")
}

func TestExtractIntField(t *testing.T) {
	t.Parallel()

	resolver := newMockResolver()

	expectedParams := &LocalGetClassParams{
		Kind:       kind.ACTION_KIND,
		ClassName:  "SomeAction",
		Properties: []SelectProperty{{Name: "intField"}},
	}

	resolver.On("LocalGetClass", expectedParams).
		Return(test_helper.EmptyListThunk(), nil).Once()

	query := "{ Get { Actions { SomeAction { intField } } } }"
	resolver.AssertResolve(t, query)
}

func TestExtractPagination(t *testing.T) {
	t.Parallel()

	resolver := newMockResolver()

	expectedParams := &LocalGetClassParams{
		Kind:       kind.ACTION_KIND,
		ClassName:  "SomeAction",
		Properties: []SelectProperty{{Name: "intField"}},
		Pagination: &common_resolver.Pagination{
			First: 10,
			After: 20,
		},
	}

	resolver.On("LocalGetClass", expectedParams).
		Return(test_helper.EmptyListThunk(), nil).Once()

	query := "{ Get { Actions { SomeAction(first:10, after: 20) { intField } } } }"
	resolver.AssertResolve(t, query)
}

func TestExtractFilterToplevelField(t *testing.T) {
	t.Parallel()

	resolver := newMockResolver()

	expectedParams := &LocalGetClassParams{
		Kind:      kind.ACTION_KIND,
		ClassName: "SomeAction",
		Filters: &common_filters.LocalFilter{Root: &common_filters.Clause{
			Operator: common_filters.OperatorEqual,
			On: &common_filters.Path{
				Class:    schema.AssertValidClassName("SomeAction"),
				Property: schema.AssertValidPropertyName("intField"),
			},
			Value: &common_filters.Value{
				Value: 42,
				Type:  schema.DataTypeInt,
			},
		}},
		Properties: []SelectProperty{{Name: "intField"}},
	}

	resolver.On("LocalGetClass", expectedParams).
		Return(test_helper.EmptyListThunk(), nil).Once()

	query := `{ Get(where: { path: ["SomeAction", "intField"], operator: Equal, valueInt: 42}) { Actions { SomeAction { intField } } } }`
	resolver.AssertResolve(t, query)
}

func TestExtractFilterNestedField(t *testing.T) {
	t.Parallel()

	resolver := newMockResolver()

	expectedParams := &LocalGetClassParams{
		Kind:      kind.ACTION_KIND,
		ClassName: "SomeAction",
		Filters: &common_filters.LocalFilter{Root: &common_filters.Clause{
			Operator: common_filters.OperatorEqual,
			On: &common_filters.Path{
				Class:    schema.AssertValidClassName("SomeAction"),
				Property: schema.AssertValidPropertyName("hasAction"),
				Child: &common_filters.Path{
					Class:    schema.AssertValidClassName("SomeAction"),
					Property: schema.AssertValidPropertyName("intField"),
				},
			},
			Value: &common_filters.Value{
				Value: 42,
				Type:  schema.DataTypeInt,
			},
		}},
		Properties: []SelectProperty{{Name: "intField"}},
	}

	resolver.On("LocalGetClass", expectedParams).
		Return(test_helper.EmptyListThunk(), nil).Once()

	query := `{ Get(where: { path: ["SomeAction", "HasAction", "SomeAction", "intField"], operator: Equal, valueInt: 42}) { Actions { SomeAction { intField } } } }`
	resolver.AssertResolve(t, query)
}
