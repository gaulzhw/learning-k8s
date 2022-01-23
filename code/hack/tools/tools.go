//go:build tools
// +build tools

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "github.com/go-bindata/go-bindata/v3"
	_ "github.com/golangci/golangci-lint/cmd/golangci-lint"
	_ "github.com/onsi/ginkgo/ginkgo/v2"
	_ "k8s.io/code-generator/cmd/conversion-gen"
	_ "sigs.k8s.io/apiserver-builder-alpha/cmd/apiserver-boot"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
