package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetTeacherByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTeacherByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTeacherByIdLogic {
	return &GetTeacherByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTeacherByIdLogic) GetTeacherById(in *example.UUIDReq) (*example.TeacherInfo, error) {
	result, err := l.svcCtx.DB.Teacher.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.TeacherInfo{
		Id:            result.ID.String(),
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
