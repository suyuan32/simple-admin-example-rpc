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
    err := l.svcCtx.DB.Student.UpdateOneID(in.Id).
			SetNotEmptyName(in.Name).
			SetNotEmptyAge(int(in.Age)).
			SetNotEmptyAgeInt8(int8(in.AgeInt8)).
			SetNotEmptyAgeUint8(uint8(in.AgeUint8)).
			SetNotEmptyAgeInt16(int16(in.AgeInt16)).
			SetNotEmptyAgeUint16(uint16(in.AgeUint16)).
			SetNotEmptyAgeInt32(in.AgeInt32).
			SetNotEmptyAgeUint32(in.AgeUint32).
			SetNotEmptyAgeInt64(in.AgeInt64).
			SetNotEmptyAgeUint64(in.AgeUint64).
			SetNotEmptyAgeInt(int(in.AgeInt)).
			SetNotEmptyAgeUint(uint(in.AgeUint)).
			SetNotEmptyWeightFloat(in.WeightFloat).
			SetNotEmptyWeightFloat32(in.WeightFloat32).
			SetNotEmptyClassID(uuidx.ParseUUIDString(in.ClassId)).
			SetNotEmptyEnrollAt(time.Unix(in.EnrollAt, 0)).
			SetNotEmptyStatusBool(in.StatusBool).
			Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &example.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
