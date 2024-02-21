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

type CreateTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTeacherLogic {
	return &CreateTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTeacherLogic) CreateTeacher(in *example.TeacherInfo) (*example.BaseIDResp, error) {
	query := l.svcCtx.DB.Teacher.Create().
		SetNotNilName(in.Name)

	if in.Age != nil {
		query.SetNotNilAge(pointy.GetPointer(int16(*in.Age)))
	}

	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
