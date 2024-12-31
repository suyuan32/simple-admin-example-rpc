package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/msg/errormsg"
	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudentLogic {
	return &UpdateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateStudentLogic) UpdateStudent(in *example.StudentInfo) (*example.BaseResp, error) {
	query := l.svcCtx.DB.Student.UpdateOneID(uuidx.ParseUUIDString(*in.Id)).
		SetNotNilName(in.Name).
		SetNotNilAddress(in.Address).
		SetNotNilScore(in.Score).
		SetNotNilWeight(in.Weight).
		SetNotNilHealthy(in.Healthy).
		SetNotNilCode(in.Code).
		SetNotNilIdentifyID(in.IdentifyId)

	if in.Status != nil {
		query.SetNotNilStatus(pointy.GetPointer(uint8(*in.Status)))
	}
	if in.Age != nil {
		query.SetNotNilAge(pointy.GetPointer(int16(*in.Age)))
	}
	if in.Height != nil {
		query.SetNotNilHeight(pointy.GetPointer(int(*in.Height)))
	}

	err := query.Exec(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.BaseResp{Msg: errormsg.UpdateSuccess}, nil
}
