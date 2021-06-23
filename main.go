package main

import (
	"context"

	"github.com/open-policy-agent/golang-opa-wasm/opa"
	"github.com/open-policy-agent/opa/rego"
)

func main() {
	panic("this is only a benchmarking program. See benchmark_main_test.go for details")
}

func wasmEval(ctx context.Context, rego *opa.OPA, input interface{}) {
	_, err := rego.Eval(ctx, &input)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Policy result: %v\n", result)
}

func opalibEval(ctx context.Context, pq *rego.PreparedEvalQuery, input interface{}) {
	_, err := pq.Eval(ctx, rego.EvalInput(input))
	if err != nil {
		panic(err)
	}
	//fmt.Printf("Policy result: %v\n", result)
}
