//args: -Egoimports
//config: linters-settings.goimports.local-prefixes=github.com/adits31/golangci-lint
package goimports

import (
	"fmt"

	"github.com/adits31/golangci-lint/pkg/config"
	"github.com/pkg/errors"
)

func GoimportsLocalTest() {
	fmt.Print("x")
	_ = config.Config{}
	_ = errors.New("")
}
