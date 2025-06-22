package metrics

import (
	"fmt"

	"github.com/mcdaigle1/slaslosli/internal/metrics/prometheus"
)

func NewCPU(kind string) (CPU, error) {
	switch kind {
	case "prometheus":
		return prometheus.PromCPU{}, nil
	default:
		return nil, fmt.Errorf("unknown monitoring kind: %s", kind)
	}
}
