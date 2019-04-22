package schema

import (
	"reflect"
	"testing"

	"gin-http-server/model"

	"github.com/graphql-go/graphql"
	"github.com/graphql-go/graphql/testutil"
)

func TestMain(m *testing.M) {
	model.MainTest(m, "..")
}

type T struct {
	Query    string
	Schema   graphql.Schema
	Expected *graphql.Result
}

func testGraphql(test T, p graphql.Params, t *testing.T) {
	result := graphql.Do(p)
	if len(result.Errors) > 0 {
		t.Fatalf("wrong result, unexpected errors: %v", result.Errors)
	}
	if !reflect.DeepEqual(result, test.Expected) {
		t.Fatalf("wrong result, query: %v, graphql result diff: %v", test.Query, testutil.Diff(test.Expected, result))
	}
}
