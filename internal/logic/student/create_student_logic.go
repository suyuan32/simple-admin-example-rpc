package student

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
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
	result, err := l.svcCtx.DB.Student.Create().
		SetNotNilName(in.Name).
		SetNotNilAge(pointy.GetPointer(int(*in.Age))).
		SetNotNilAgeInt8(pointy.GetPointer(int8(*in.AgeInt8))).
		SetNotNilAgeUint8(pointy.GetPointer(uint8(*in.AgeUint8))).
		SetNotNilAgeInt16(pointy.GetPointer(int16(*in.AgeInt16))).
		SetNotNilAgeUint16(pointy.GetPointer(uint16(*in.AgeUint16))).
		SetNotNilAgeInt32(in.AgeInt32).
		SetNotNilAgeUint32(in.AgeUint32).
		SetNotNilAgeInt64(in.AgeInt64).
		SetNotNilAgeUint64(in.AgeUint64).
		SetNotNilAgeInt(pointy.GetPointer(int(*in.AgeInt))).
		SetNotNilAgeUint(pointy.GetPointer(uint(*in.AgeUint))).
		SetNotNilWeightFloat(in.WeightFloat).
		SetNotNilWeightFloat32(in.WeightFloat32).
		SetNotNilClassID(uuidx.ParseUUIDStringToPointer(*in.ClassId)).
		SetNotNilEnrollAt(pointy.GetPointer(time.Unix(*in.EnrollAt, 0))).
		SetNotNilStatusBool(in.StatusBool).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.BaseIDResp{Id: result.ID, Msg: errormsg.CreateSuccess}, nil
}
