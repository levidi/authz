package opa

import (
	"context"
	"errors"
	"fmt"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/open-policy-agent/opa/util"
)

func newOpa() (query rego.PreparedEvalQuery, errOpa error) {

	ctx := context.Background()

	module := `
				package security

				default allow = false
				
				allow {
					input.profile == data.rules[_].profile[_]
				}
		   	`

	data := `
				{
					"rules": [
						{
							"profile": ["administrator", "financial"]
						}
					]
				}
		    `
	var json map[string]interface{}

	err := util.UnmarshalJSON([]byte(data), &json)
	if err != nil {
		errOpa = err
	}

	store := inmem.NewFromObject(json)

	compiler, err := ast.CompileModules(map[string]string{
		"module.rego": module,
	})

	query, err = rego.New(
		rego.Compiler(compiler),
		rego.Query("data.security"),
		rego.Store(store),
	).PrepareForEval(ctx)
	return
}

// Eval evaluate policy
func Eval(input map[string]interface{}) bool {
	fmt.Println("input:", input)
	results, err := policyEval(context.Background(), input)

	mapResults := results[0].Expressions[0].Value.(map[string]interface{})
	fmt.Println("result of allow", mapResults["allow"].(bool))
	fmt.Println("err:", err)
	return mapResults["allow"].(bool)
}

func policyEval(ctx context.Context, input map[string]interface{}) (results rego.ResultSet, err error) {

	query, err := newOpa()
	results, err = query.Eval(ctx, rego.EvalInput(input))

	if err != nil && len(results) == 0 {
		err = errors.New("Handle undefined result")
	}
	return
}
