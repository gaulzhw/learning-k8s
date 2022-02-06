//go:build tools
// +build tools

// This package imports things required by build scripts, to force `go mod` to see them as dependencies
package tools

import (
	_ "github.com/drone/envsubst/v2/cmd/envsubst"
	_ "github.com/joelanford/go-apidiff"
	_ "github.com/onsi/ginkgo/ginkgo"
	_ "gotest.tools/gotestsum"
	_ "k8s.io/code-generator/cmd/conversion-gen"
	_ "sigs.k8s.io/controller-runtime/tools/setup-envtest"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
	_ "sigs.k8s.io/promo-tools/v3/cmd/kpromo"
)
