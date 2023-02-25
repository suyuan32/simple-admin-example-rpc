package student

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-example-rpc/ent"
	"github.com/suyuan32/simple-admin-example-rpc/example"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"
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
		SetNotEmptyAgeInt32(in.AgeInt32).
		SetNotEmptyAgeInt64(in.AgeInt64).
		SetNotEmptyAgeUint(uint(in.AgeUint)).
		SetNotEmptyAgeUint32(in.AgeUint32).
		SetNotEmptyAgeUint64(in.AgeUint64).
		SetNotEmptyWeightFloat(in.WeightFloat).
		SetNotEmptyWeightFloat32(in.WeightFloat32).
		SetNotEmptyClassID(uuidx.ParseUUIDString(in.ClassId)).
		SetNotEmptyEnrollAt(time.Unix(in.EnrollAt, 0)).
		SetNotEmptyStatusBool(in.StatusBool).
		Exec(l.ctx)

	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &example.BaseResp{Msg: i18n.UpdateSuccess}, nil
}
