package schema

import (
	"context"
	"testing"

	"gin-http-server/model"

	"github.com/graphql-go/graphql"
)

func TestQueryUser(t *testing.T) {
	model.PrepareTestDatabase()
	t.Run("single user", func(t *testing.T) {
		test := T{
			Query: `
query getUser (
    $id: Int!
) {
  getUser(id: $id) {
    name
  }
}
	  `,
			Schema: Schema,
			Expected: &graphql.Result{
				Data: map[string]interface{}{
					"getUser": map[string]interface{}{
						"name": "appleboy",
					},
				},
			},
		}
		params := graphql.Params{
			Schema:        test.Schema,
			RequestString: test.Query,
			Context:       context.TODO(),
			VariableValues: map[string]interface{}{
				"id": 1,
			},
		}
		testGraphql(test, params, t)
	})
}
