package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/example"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateStudentLogic {
	return &CreateOrUpdateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Student management
func (l *CreateOrUpdateStudentLogic) CreateOrUpdateStudent(in *example.StudentInfo) (*example.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &example.BaseResp{}, nil
}
