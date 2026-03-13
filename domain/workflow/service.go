package workflow

import "context"

// Service 定义 q-workflow 领域层最小边界。
// 未来工作流状态机、调度策略、执行编排都应在此层沉淀。
type Service interface {
	Name() string
	Health(ctx context.Context) error
}

type service struct{}

func NewService() Service {
	return &service{}
}

func (s *service) Name() string {
	return "workflow"
}

func (s *service) Health(_ context.Context) error {
	return nil
}
