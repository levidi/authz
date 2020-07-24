package opa

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/open-policy-agent/opa/ast"
	"github.com/open-policy-agent/opa/rego"
	"github.com/open-policy-agent/opa/storage/inmem"
	"github.com/open-policy-agent/opa/util"
)

var mapOPA = sync.Map{}

type configPolicy struct {
	RegoRules         []string `json:"regoRules"`
	RegoQueryName     string   `json:"regoQueryName"`
	preparedEvalQuery rego.PreparedEvalQuery
}

func getFileInBytes(path string) (bytePolicy []byte) {
	bytePolicy, errPolicy := ioutil.ReadFile(path)
	if errPolicy != nil {
		log.Fatal(errPolicy)
	}
	return
}

func convertJSONBytesToMap(jsonBytes []byte) (jsonMap map[string]interface{}) {
	errData := util.UnmarshalJSON(jsonBytes, &jsonMap)
	if errData != nil {
		log.Fatal(errData)
	}
	return jsonMap
}

func new() (errOpa error) {
	ctx := context.Background()

	path := os.Getenv("PATH_CONFIG_MAP")

	bytePolicy := getFileInBytes(path + "/validate-profile.rego")
	byteData := getFileInBytes(path + "/data.json")
	byteConfigPolicy := getFileInBytes(path + "/config-policy.json")

	var config configPolicy
	errData := util.UnmarshalJSON(byteConfigPolicy, &config)
	if errData != nil {
		log.Fatal(errData)
	}
	policy := string(bytePolicy)

	compiler, err := ast.CompileModules(map[string]string{
		"module.rego": policy,
	})

	jsonMap := convertJSONBytesToMap(byteData)

	store := inmem.NewFromObject(jsonMap)
	preparedEvalQuery, err := rego.New(
		rego.Compiler(compiler),
		rego.Query("data."+config.RegoQueryName),
		rego.Store(store),
	).PrepareForEval(ctx)

	if err != nil {
		panic(err)
	}
	config.preparedEvalQuery = preparedEvalQuery
	mapOPA.Store("configPolicy", config)
	return
}

func getOPA() (config configPolicy) {
	new()
	value, ok := mapOPA.Load("configPolicy")
	if ok {
		return value.(configPolicy)
	}
	panic("configPolicy not found")
}

// Eval evaluate policy
func Eval(input map[string]interface{}) bool {
	allow, _ := policyEval(context.Background(), input)
	return allow
}

func policyEval(ctx context.Context, input map[string]interface{}) (allow bool, err error) {
	config := getOPA()
	results, err := config.preparedEvalQuery.Eval(ctx, rego.EvalInput(input))
	if err != nil && len(results) == 0 {
		err = errors.New("Handle undefined result")
	}

	mapResults := results[0].Expressions[0].Value.(map[string]interface{})
	for rule, t := range mapResults {
		switch v := t.(type) {
		case bool:
			allow = mapResults[rule].(bool)
		default:
			fmt.Printf("unexpected type %T", v)
		}

	}
	return
}
