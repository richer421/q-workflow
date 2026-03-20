package workflow

import (
	"context"

	domainworkflow "github.com/richer421/q-workflow/domain/workflow"
)

// app 是 q-workflow 的应用层占位。
// 后续工作流编排入口统一从这里暴露给 HTTP/MCP。
type app struct{}

var App = new(app)

var workflowDomain = domainworkflow.NewService()

func (s *app) Health(ctx context.Context) error {
	return workflowDomain.Health(ctx)
}
