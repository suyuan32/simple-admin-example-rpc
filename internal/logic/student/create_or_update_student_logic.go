package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/ent"
	"github.com/suyuan32/simple-admin-example-rpc/example"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateOrUpdateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateOrUpdateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateOrUpdateStudentLogic {
	return &CreateOrUpdateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateOrUpdateStudentLogic) CreateOrUpdateStudent(in *example.StudentInfo) (*example.BaseResp, error) {
	if in.Id == 0 {
		err := l.svcCtx.DB.Student.Create().
			SetName(in.Name).
			SetAge(int(in.Age)).
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
		err := l.svcCtx.DB.Student.UpdateOneID(in.Id).
			SetName(in.Name).
			SetAge(int(in.Age)).
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
