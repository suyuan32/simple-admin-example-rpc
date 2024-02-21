package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTeacherByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherByIdLogic {
	return &GetTeacherByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTeacherByIdLogic) GetTeacherById(in *example.IDReq) (*example.TeacherInfo, error) {
	result, err := l.svcCtx.DB.Teacher.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.TeacherInfo{
		Id:        &result.ID,
		CreatedAt: pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt: pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Name:      &result.Name,
		Age:       pointy.GetPointer(int32(result.Age)),
	}, nil
}
