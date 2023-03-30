package student

import (
	"context"

    "github.com/suyuan32/simple-admin-example-rpc/ent/student"
    "github.com/suyuan32/simple-admin-example-rpc/internal/svc"
    "github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
    "github.com/suyuan32/simple-admin-example-rpc/types/example"

    "github.com/suyuan32/simple-admin-common/i18n"
    "github.com/zeromicro/go-zero/core/logx"
)

type DeleteStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteStudentLogic {
	return &DeleteStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteStudentLogic) DeleteStudent(in *example.IDsReq) (*example.BaseResp, error) {
	_, err := l.svcCtx.DB.Student.Delete().Where(student.IDIn(in.Ids...)).Exec(l.ctx)

    if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

    return &example.BaseResp{Msg: i18n.DeleteSuccess}, nil
}
