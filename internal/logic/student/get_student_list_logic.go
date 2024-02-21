package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-example-rpc/ent/student"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
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

func (l *GetStudentListLogic) GetStudentList(in *example.StudentListReq) (*example.StudentListResp, error) {
	var predicates []predicate.Student
	if in.Name != nil {
		predicates = append(predicates, student.NameContains(*in.Name))
	}
	if in.Address != nil {
		predicates = append(predicates, student.AddressContains(*in.Address))
	}
	result, err := l.svcCtx.DB.Student.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &example.StudentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &example.StudentInfo{
			Id:        pointy.GetPointer(v.ID.String()),
			CreatedAt: pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt: pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Name:      &v.Name,
			Age:       pointy.GetPointer(int32(v.Age)),
			Address:   &v.Address,
		})
	}

	return resp, nil
}
