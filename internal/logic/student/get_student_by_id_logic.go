package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
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
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.StudentInfo{
		Id:            &result.ID,
		CreatedAt:     pointy.GetPointer(result.CreatedAt.Unix()),
		UpdatedAt:     pointy.GetPointer(result.UpdatedAt.Unix()),
		Name:          &result.Name,
		Age:           pointy.GetPointer(int64(result.Age)),
		AgeInt8:       pointy.GetPointer(int32(result.AgeInt8)),
		AgeUint8:      pointy.GetPointer(uint32(result.AgeUint8)),
		AgeInt16:      pointy.GetPointer(int32(result.AgeInt16)),
		AgeUint16:     pointy.GetPointer(uint32(result.AgeUint16)),
		AgeInt32:      &result.AgeInt32,
		AgeUint32:     &result.AgeUint32,
		AgeInt64:      &result.AgeInt64,
		AgeUint64:     &result.AgeUint64,
		AgeInt:        pointy.GetPointer(int64(result.AgeInt)),
		AgeUint:       pointy.GetPointer(uint64(result.AgeUint)),
		WeightFloat:   &result.WeightFloat,
		WeightFloat32: &result.WeightFloat32,
		ClassId:       pointy.GetPointer(result.ClassID.String()),
		EnrollAt:      pointy.GetPointer(result.EnrollAt.UnixMilli()),
		StatusBool:    &result.StatusBool,
	}, nil
}
