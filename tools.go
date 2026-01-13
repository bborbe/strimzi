// Copyright (c) 2024 Benjamin Borbe All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build tools
// +build tools

package tools

import (
	_ "cloud.google.com/go/compute/metadata"
	_ "github.com/actgardner/gogen-avro/v9/cmd/gogen-avro"
	_ "github.com/google/addlicense"
	_ "github.com/incu6us/goimports-reviser/v3"
	_ "github.com/kisielk/errcheck"
	_ "github.com/maxbrunsfeld/counterfeiter/v6"
	_ "github.com/shoenig/go-modtool"
	_ "golang.org/x/lint/golint"
	_ "golang.org/x/vuln/cmd/govulncheck"
	_ "k8s.io/code-generator"
	_ "k8s.io/code-generator/cmd/validation-gen"
)
