package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/ent"
	"github.com/suyuan32/simple-admin-example-rpc/example"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/msg/logmsg"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
	"github.com/suyuan32/simple-admin-core/pkg/uuidx"
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
		switch {
		case ent.IsNotFound(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.TargetNotFound)
		case ent.IsConstraintError(err):
			logx.Errorw(err.Error(), logx.Field("detail", in))
			return nil, statuserr.NewInvalidArgumentError(i18n.UpdateFailed)
		default:
			logx.Errorw(logmsg.DatabaseError, logx.Field("detail", err.Error()))
			return nil, statuserr.NewInternalError(i18n.DatabaseError)
		}
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
