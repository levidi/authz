package opa

import (
	"encoding/json"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Setenv("PATH_CONFIG_MAP", "../configMap")
	code := m.Run()
	os.Exit(code)
}

func TestEval(t *testing.T) {
	const msg = "Allow %v"

	const body = `{"name": "Levi Di Tomazzo", "age": 34, "profile": "administratorx"}`
	input := make(map[string]interface{})
	errUn := json.Unmarshal([]byte(body), &input)
	if errUn != nil {
		t.Errorf("err %v", errUn)
	}

	isAllow := Eval(input)

	if isAllow {
		t.Logf(msg, isAllow)
	} else {
		t.Errorf(msg, isAllow)
	}
}
