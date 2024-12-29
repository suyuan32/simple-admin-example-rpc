package student

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/i18n"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateStudentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateStudentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateStudentLogic {
	return &CreateStudentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateStudentLogic) CreateStudent(in *example.StudentInfo) (*example.BaseUUIDResp, error) {
	query := l.svcCtx.DB.Student.Create().
		SetNotNilName(in.Name).
		SetNotNilAddress(in.Address).
		SetNotNilScore(in.Score).
		SetNotNilWeight(in.Weight).
		SetNotNilHealthy(in.Healthy).
		SetNotNilCode(in.Code)

	if in.Age != nil {
		query.SetNotNilAge(pointy.GetPointer(int16(*in.Age)))
	}

	result, err := query.Save(l.ctx)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	return &example.BaseUUIDResp{Id: result.ID.String(), Msg: i18n.CreateSuccess}, nil
}
