package teacher

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-example-rpc/example"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"

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

func (l *CreateTeacherLogic) CreateTeacher(in *example.TeacherInfo) (*example.BaseUUIDResp, error) {
	result, err := l.svcCtx.DB.Teacher.Create().
		SetName(in.Name).
		SetAge(int(in.Age)).
		SetAgeInt32(in.AgeInt32).
		SetAgeInt64(in.AgeInt64).
		SetAgeUint(uint(in.AgeUint)).
		SetAgeUint32(in.AgeUint32).
		SetAgeUint64(in.AgeUint64).
		SetWeightFloat(in.WeightFloat).
		SetWeightFloat32(in.WeightFloat32).
		SetClassID(uuidx.ParseUUIDString(in.ClassId)).
		SetEnrollAt(time.Unix(in.EnrollAt, 0)).
		SetStatusBool(in.StatusBool).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
