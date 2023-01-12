package teacher

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

type CreateOrUpdateTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateTeacherLogic {
	return &CreateOrUpdateTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateTeacherLogic) CreateOrUpdateTeacher(in *example.TeacherInfo) (*example.BaseResp, error) {
	if in.Id == "" {
		err := l.svcCtx.DB.Teacher.Create().
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
			Exec(l.ctx)
		if err != nil {
			switch {
			case ent.IsConstraintError(err):
				logx.Errorw(err.Error(), logx.Field("detail", in))
				return nil, statuserr.NewInvalidArgumentError(i18n.CreateFailed)
			default:
				logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
				return nil, statuserr.NewInternalError(i18n.DatabaseError)
			}
		}

		return &example.BaseResp{Msg: i18n.CreateSuccess}, nil
	} else {
		err := l.svcCtx.DB.Teacher.UpdateOneID(uuidx.ParseUUIDString(in.Id)).
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
}
