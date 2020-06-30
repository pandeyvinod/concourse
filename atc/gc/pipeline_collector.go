package gc

import "context"

type pipelineCollector struct{}

func NewPipelineCollector() *pipelineCollector {
	return &pipelineCollector{}
}

func (pc *pipelineCollector) Run(ctx context.Context) error {
	return nil
}
