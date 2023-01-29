package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/example"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDeleteStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteStudentLogic {
	return &BatchDeleteStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchDeleteStudentLogic) BatchDeleteStudent(in *example.IDsReq) (*example.BaseResp, error) {
	// todo: add your logic here and delete this line

	return &example.BaseResp{}, nil
}
