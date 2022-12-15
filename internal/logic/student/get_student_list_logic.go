package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-example-rpc/ent/student"
	"github.com/suyuan32/simple-admin-example-rpc/example"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentListLogic {
	return &GetStudentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentListLogic) GetStudentList(in *example.StudentPageReq) (*example.StudentListResp, error) {
	var predicates []predicate.Student
	if in.Name != "" {
		predicates = append(predicates, student.NameContains(in.Name))
	}
	result, err := l.svcCtx.DB.Student.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &example.StudentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &example.StudentInfo{
			Id:        v.ID,
			CreatedAt: v.CreatedAt.UnixMilli(),
			Name:      v.Name,
			Age:       int64(v.Age),
		})
	}

	return resp, nil
}
