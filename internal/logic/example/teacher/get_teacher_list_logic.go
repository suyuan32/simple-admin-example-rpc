package teacher

import (
	"context"

	"github.com/suyuan32/simple-admin-example-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-example-rpc/ent/teacher"
	"github.com/suyuan32/simple-admin-example-rpc/example"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"

	"github.com/suyuan32/simple-admin-core/pkg/i18n"
	"github.com/suyuan32/simple-admin-core/pkg/statuserr"
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

func (l *GetTeacherListLogic) GetTeacherList(in *example.TeacherPageReq) (*example.TeacherListResp, error) {
	var predicates []predicate.Teacher
	if in.Name != "" {
		predicates = append(predicates, teacher.NameContains(in.Name))
	}
	result, err := l.svcCtx.DB.Teacher.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		logx.Error(err.Error())
		return nil, statuserr.NewInternalError(i18n.DatabaseError)
	}

	resp := &example.TeacherListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &example.TeacherInfo{
			Id:            v.ID.String(),
			CreatedAt:     v.CreatedAt.UnixMilli(),
			Name:          v.Name,
			Age:           int64(v.Age),
			AgeInt32:      v.AgeInt32,
			AgeInt64:      v.AgeInt64,
			AgeUint:       uint64(v.AgeUint),
			AgeUint32:     v.AgeUint32,
			AgeUint64:     v.AgeUint64,
			WeightFloat:   v.WeightFloat,
			WeightFloat32: v.WeightFloat32,
			ClassId:       v.ClassID.String(),
			EnrollAt:      v.EnrollAt.UnixMilli(),
			StatusBool:    v.StatusBool,
		})
	}

	return resp, nil
}
