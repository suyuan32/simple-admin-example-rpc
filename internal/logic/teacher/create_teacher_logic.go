package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

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

func (l *CreateTeacherLogic) CreateTeacher(in *example.TeacherInfo) (*example.BaseUUIDResp, error) {
	query := l.svcCtx.DB.Teacher.Create().
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

	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
