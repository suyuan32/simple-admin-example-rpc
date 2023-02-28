package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/example"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudentByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudentByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudentByIdLogic {
	return &GetStudentByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetStudentByIdLogic) GetStudentById(in *example.IDReq) (*example.StudentInfo, error) {
	result, err := l.svcCtx.DB.Student.Get(l.ctx, in.Id)
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(err, in)
	}

	return &example.StudentInfo{
		Id:            result.ID,
		CreatedAt:     result.CreatedAt.UnixMilli(),
		UpdatedAt:     result.UpdatedAt.UnixMilli(),
		Name:          result.Name,
		Age:           int64(result.Age),
		AgeInt32:      result.AgeInt32,
		AgeInt64:      result.AgeInt64,
		AgeUint:       uint64(result.AgeUint),
		AgeUint32:     result.AgeUint32,
		AgeUint64:     result.AgeUint64,
		WeightFloat:   result.WeightFloat,
		WeightFloat32: result.WeightFloat32,
		ClassId:       result.ClassID.String(),
		EnrollAt:      result.EnrollAt.UnixMilli(),
		StatusBool:    result.StatusBool,
	}, nil
}
