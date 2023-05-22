package _go

import _ "github.com/ogen-go/ogen"

//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target ./gen/ -package gen --clean ./../openapi.yaml
