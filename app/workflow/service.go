package workflow

import (
	"context"

	domainworkflow "github.com/richer421/q-workflow/domain/workflow"
)

// AppService 是 q-workflow 的应用层占位。
// 后续工作流编排入口统一从这里暴露给 HTTP/MCP。
type AppService struct {
	domain domainworkflow.Service
}

func NewAppService() *AppService {
	return &AppService{
		domain: domainworkflow.NewService(),
	}
}

func (s *AppService) Health(ctx context.Context) error {
	return s.domain.Health(ctx)
}
