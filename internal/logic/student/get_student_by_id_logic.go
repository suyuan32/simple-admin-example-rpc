package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
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

func (l *GetStudentByIdLogic) GetStudentById(in *example.UUIDReq) (*example.StudentInfo, error) {
	result, err := l.svcCtx.DB.Student.Get(l.ctx, uuidx.ParseUUIDString(in.Id))
	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.StudentInfo{
		Id:         pointy.GetPointer(result.ID.String()),
		CreatedAt:  pointy.GetPointer(result.CreatedAt.UnixMilli()),
		UpdatedAt:  pointy.GetPointer(result.UpdatedAt.UnixMilli()),
		Status:     pointy.GetPointer(uint32(result.Status)),
		Name:       &result.Name,
		Age:        pointy.GetPointer(int32(result.Age)),
		Address:    &result.Address,
		Score:      &result.Score,
		Weight:     &result.Weight,
		Healthy:    &result.Healthy,
		Code:       &result.Code,
		IdentifyId: &result.IdentifyID,
		Height:     pointy.GetPointer(int64(result.Height)),
	}, nil
}
