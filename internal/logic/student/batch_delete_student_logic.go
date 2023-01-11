package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/ent"
	"github.com/suyuan32/simple-admin-example-rpc/ent/student"
	"github.com/suyuan32/simple-admin-example-rpc/example"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/zeromicro/go-zero/core/logx"
)

type BatchDeleteStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchDeleteStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchDeleteStudentLogic {
	return &BatchDeleteStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BatchDeleteStudentLogic) BatchDeleteStudent(in *example.IDsReq) (*example.BaseResp, error) {
	_, err := l.svcCtx.DB.Student.Delete().Where(student.IDIn(in.Ids...)).Exec(l.ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
	}

	return &example.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
