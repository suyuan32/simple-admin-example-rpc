package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-example-rpc/ent/teacher"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTeacherListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherListLogic {
	return &GetTeacherListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTeacherListLogic) GetTeacherList(in *example.TeacherListReq) (*example.TeacherListResp, error) {
	var predicates []predicate.Teacher
	if in.Name != nil {
		predicates = append(predicates, teacher.NameContains(*in.Name))
	}
	if in.Age != nil {
		predicates = append(predicates, teacher.AgeEQ(int16(*in.Age)))
	}
	result, err := l.svcCtx.DB.Teacher.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &example.TeacherListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &example.TeacherInfo{
			Id:        &v.ID,
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Name:      &v.Name,
			Age:       pointy.GetPointer(int32(v.Age)),
		})
	}

	return resp, nil
}
