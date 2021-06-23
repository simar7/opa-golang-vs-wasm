package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/open-policy-agent/opa/rego"

	"github.com/open-policy-agent/golang-opa-wasm/opa"
)

func BenchmarkWasm(b *testing.B) {
	b.ReportAllocs()

	ctx := context.Background()
	policy, err := ioutil.ReadFile("example.wasm")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	rego, err := opa.New().WithPolicyBytes(policy).Init()
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	defer rego.Close()

	// Evaluate the policy once.
	var input interface{} = map[string]interface{}{
		"foo": true,
		"bar": false,
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		wasmEval(ctx, rego, input)
	}
}

func BenchmarkOPAGoLibrary(b *testing.B) {
	b.ReportAllocs()
	policy, err := ioutil.ReadFile("example.rego")
	if err != nil {
		fmt.Printf("error: %v\n", err)
		return
	}

	// Evaluate the policy once.
	var input interface{} = map[string]interface{}{
		"foo": true,
		"bar": false,
	}

	ctx := context.Background()

	pq, err := rego.New(
		rego.Query("x = data.example.allow"),
		rego.Module("example-1.rego", string(policy)),
	).PrepareForEval(ctx)
	if err != nil {
		panic(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		opalibEval(ctx, &pq, input)
	}
}
