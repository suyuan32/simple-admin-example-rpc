package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateTeacherLogic {
	return &UpdateTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateTeacherLogic) UpdateTeacher(in *example.TeacherInfo) (*example.BaseResp, error) {
	query := l.svcCtx.DB.Teacher.UpdateOneID(*in.Id).
		SetNotNilName(in.Name)

	if in.Age != nil {
		query.SetNotNilAge(pointy.GetPointer(int16(*in.Age)))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
