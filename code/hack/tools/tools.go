//go:build tools
// +build tools

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "github.com/go-bindata/go-bindata"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/onsi/ginkgo/ginkgo"
	_ "k8s.io/code-generator/cmd/conversion-gen"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
