package student

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/i18n"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"

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
		SetName(in.Name).
		SetAge(int(in.Age)).
		SetAgeInt8(int8(in.AgeInt8)).
		SetAgeUint8(uint8(in.AgeUint8)).
		SetAgeInt16(int16(in.AgeInt16)).
		SetAgeUint16(uint16(in.AgeUint16)).
		SetAgeInt32(in.AgeInt32).
		SetAgeUint32(in.AgeUint32).
		SetAgeInt64(in.AgeInt64).
		SetAgeUint64(in.AgeUint64).
		SetAgeInt(int(in.AgeInt)).
		SetAgeUint(uint(in.AgeUint)).
		SetWeightFloat(in.WeightFloat).
		SetWeightFloat32(in.WeightFloat32).
		SetClassID(uuidx.ParseUUIDString(in.ClassId)).
		SetEnrollAt(time.Unix(in.EnrollAt, 0)).
		SetStatusBool(in.StatusBool).
		Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
