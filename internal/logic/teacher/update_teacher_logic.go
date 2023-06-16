package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
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
	query := l.svcCtx.DB.Teacher.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
		SetNotNilName(in.Name).
		SetNotNilAgeInt32(in.AgeInt32).
		SetNotNilAgeInt64(in.AgeInt64).
		SetNotNilAgeUint32(in.AgeUint32).
		SetNotNilAgeUint64(in.AgeUint64).
		SetNotNilWeightFloat(in.WeightFloat).
		SetNotNilWeightFloat32(in.WeightFloat32).
		SetNotNilClassID(uuidx.ParseUUIDStringToPointer(in.ClassId)).
		SetNotNilEnrollAt(pointy.GetTimePointer(in.EnrollAt, 0)).
		SetNotNilStatusBool(in.StatusBool)

	if in.Age != nil {
		query.SetNotNilAge(pointy.GetPointer(int(*in.Age)))
	}
	if in.AgeUint != nil {
		query.SetNotNilAgeUint(pointy.GetPointer(uint(*in.AgeUint)))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.BaseResp{Msg: errormsg.UpdateSuccess}, nil
}
