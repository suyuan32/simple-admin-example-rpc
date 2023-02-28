package student

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
		return nil, dberrorhandler.DefaultEntError(err, in)
	}

	return &example.BaseIDResp{Id: result.ID, Msg: i18n.CreateSuccess}, nil
}
