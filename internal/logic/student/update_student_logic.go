package student

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

type UpdateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentLogic {
	return &UpdateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStudentLogic) UpdateStudent(in *example.StudentInfo) (*example.BaseResp, error) {
	query := l.svcCtx.DB.Student.UpdateOneID(*in.Id).
		SetNotNilName(in.Name).
		SetNotNilAgeInt32(in.AgeInt32).
		SetNotNilAgeUint32(in.AgeUint32).
		SetNotNilAgeInt64(in.AgeInt64).
		SetNotNilAgeUint64(in.AgeUint64).
		SetNotNilWeightFloat(in.WeightFloat).
		SetNotNilWeightFloat32(in.WeightFloat32).
		SetNotNilClassID(uuidx.ParseUUIDStringToPointer(in.ClassId)).
		SetNotNilEnrollAt(pointy.GetTimePointer(in.EnrollAt, 0)).
		SetNotNilStatusBool(in.StatusBool)

	if in.Age != nil {
		query.SetNotNilAge(pointy.GetPointer(int(*in.Age)))
	}
	if in.AgeInt8 != nil {
		query.SetNotNilAgeInt8(pointy.GetPointer(int8(*in.AgeInt8)))
	}
	if in.AgeUint8 != nil {
		query.SetNotNilAgeUint8(pointy.GetPointer(uint8(*in.AgeUint8)))
	}
	if in.AgeInt16 != nil {
		query.SetNotNilAgeInt16(pointy.GetPointer(int16(*in.AgeInt16)))
	}
	if in.AgeUint16 != nil {
		query.SetNotNilAgeUint16(pointy.GetPointer(uint16(*in.AgeUint16)))
	}
	if in.AgeInt != nil {
		query.SetNotNilAgeInt(pointy.GetPointer(int(*in.AgeInt)))
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
