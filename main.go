//go:generate go tool oapi-codegen -config ./etc/openapi/server.yaml ./etc/openapi/spec.yaml
//go:generate go tool sqlc generate -f ./etc/db/sqlc.yaml

package main

import (
	"context"
	"os"

	"github.com/jsnfwlr/o11y"

	"github.com/jsnfwlr/vexil/internal/cmd"
)

func main() {
	cfg, err := o11y.LoadConfig()
	if err != nil {
		panic(err)
	}

	ctx, o, err := o11y.Initialise(context.Background(), cfg, os.Stdout)
	if err != nil {
		panic(err)
	}

	defer o.Close()

	cmd.Execute(ctx)
}
