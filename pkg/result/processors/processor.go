package processors

import (
	"github.com/adits31/golangci-lint/pkg/result"
)

type Processor interface {
	Process(issues []result.Issue) ([]result.Issue, error)
	Name() string
	Finish()
}
