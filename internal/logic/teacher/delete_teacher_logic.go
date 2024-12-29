package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/ent/teacher"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteTeacherLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteTeacherLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteTeacherLogic {
	return &DeleteTeacherLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteTeacherLogic) DeleteTeacher(in *example.IDsReq) (*example.BaseResp, error) {
	_, err := l.svcCtx.DB.Teacher.Delete().Where(teacher.IDIn(in.Ids...)).Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.BaseResp{Msg: errormsg.DeleteSuccess}, nil
}
