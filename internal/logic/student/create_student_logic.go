package student

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

type CreateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStudentLogic {
	return &CreateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStudentLogic) CreateStudent(in *example.StudentInfo) (*example.BaseIDResp, error) {
	query := l.svcCtx.DB.Student.Create().
		SetNotNilName(in.Name).
		SetNotNilAgeInt32(in.AgeInt32).
		SetNotNilAgeInt64(in.AgeInt64).
		SetNotNilAgeUint32(in.AgeUint32).
		SetNotNilAgeUint64(in.AgeUint64).
		SetNotNilWeightFloat(in.WeightFloat).
		SetNotNilWeightFloat32(in.WeightFloat32).
		SetNotNilClassID(uuidx.ParseUUIDStringToPointer(in.ClassId)).
		SetNotNilEnrollAt(pointy.GetTimeMilliPointer(in.EnrollAt)).
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

	return &example.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
