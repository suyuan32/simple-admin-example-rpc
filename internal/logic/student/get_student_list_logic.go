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
	result, err := l.svcCtx.DB.Student.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &example.StudentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &example.StudentInfo{
			Id:            &v.ID,
			CreatedAt:     pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:     pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Name:          &v.Name,
			Age:           pointy.GetPointer(int64(v.Age)),
			AgeInt8:       pointy.GetPointer(int32(v.AgeInt8)),
			AgeUint8:      pointy.GetPointer(uint32(v.AgeUint8)),
			AgeInt16:      pointy.GetPointer(int32(v.AgeInt16)),
			AgeUint16:     pointy.GetPointer(uint32(v.AgeUint16)),
			AgeInt32:      &v.AgeInt32,
			AgeUint32:     &v.AgeUint32,
			AgeInt64:      &v.AgeInt64,
			AgeUint64:     &v.AgeUint64,
			AgeInt:        pointy.GetPointer(int64(v.AgeInt)),
			AgeUint:       pointy.GetPointer(uint64(v.AgeUint)),
			WeightFloat:   &v.WeightFloat,
			WeightFloat32: &v.WeightFloat32,
			ClassId:       pointy.GetPointer(v.ClassID.String()),
			EnrollAt:      pointy.GetPointer(v.EnrollAt.UnixMilli()),
			StatusBool:    &v.StatusBool,
		})
	}

	return resp, nil
}
