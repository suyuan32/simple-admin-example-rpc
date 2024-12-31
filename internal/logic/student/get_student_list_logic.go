package student

import (
	"context"
	"time"

	"github.com/suyuan32/simple-admin-example-rpc/ent/predicate"
	"github.com/suyuan32/simple-admin-example-rpc/ent/student"
	"github.com/suyuan32/simple-admin-example-rpc/internal/svc"
	"github.com/suyuan32/simple-admin-example-rpc/internal/utils/dberrorhandler"
	"github.com/suyuan32/simple-admin-example-rpc/types/example"

	"github.com/suyuan32/simple-admin-common/utils/pointy"
	"github.com/suyuan32/simple-admin-common/utils/uuidx"
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
	if in.Status != nil {
		predicates = append(predicates, student.StatusEQ(uint8(*in.Status)))
	}
	if in.Name != nil {
		predicates = append(predicates, student.NameContains(*in.Name))
	}
	if in.Age != nil {
		predicates = append(predicates, student.AgeEQ(int16(*in.Age)))
	}
	if in.Address != nil {
		predicates = append(predicates, student.AddressContains(*in.Address))
	}
	if in.Score != nil {
		predicates = append(predicates, student.ScoreEQ(*in.Score))
	}
	if in.Weight != nil {
		predicates = append(predicates, student.WeightEQ(*in.Weight))
	}
	if in.Healthy != nil {
		predicates = append(predicates, student.HealthyEQ(*in.Healthy))
	}
	if in.Code != nil {
		predicates = append(predicates, student.CodeEQ(*in.Code))
	}
	if in.IdentifyId != nil {
		predicates = append(predicates, student.IdentifyIDContains(*in.IdentifyId))
	}
	if in.Height != nil {
		predicates = append(predicates, student.HeightEQ(int(*in.Height)))
	}
	if in.ExpiredAt != nil {
		predicates = append(predicates, student.ExpiredAtGT(time.UnixMilli(*in.ExpiredAt)))
	}
	if in.StudentNumber != nil {
		predicates = append(predicates, student.StudentNumberEQ(uuidx.ParseUUIDString(*in.StudentNumber)))
	}
	result, err := l.svcCtx.DB.Student.Query().Where(predicates...).Page(l.ctx, in.Page, in.PageSize)

	if err != nil {
		return nil, dberrorhandler.DefaultEntError(l.Logger, err, in)
	}

	resp := &example.StudentListResp{}
	resp.Total = result.PageDetails.Total

	for _, v := range result.List {
		resp.Data = append(resp.Data, &example.StudentInfo{
			Id:            pointy.GetPointer(v.ID.String()),
			CreatedAt:     pointy.GetPointer(v.CreatedAt.UnixMilli()),
			UpdatedAt:     pointy.GetPointer(v.UpdatedAt.UnixMilli()),
			Status:        pointy.GetPointer(uint32(v.Status)),
			Name:          &v.Name,
			Age:           pointy.GetPointer(int32(v.Age)),
			Address:       &v.Address,
			Score:         &v.Score,
			Weight:        &v.Weight,
			Healthy:       &v.Healthy,
			Code:          &v.Code,
			IdentifyId:    &v.IdentifyID,
			Height:        pointy.GetPointer(int64(v.Height)),
			ExpiredAt:     pointy.GetUnixMilliPointer(v.ExpiredAt.UnixMilli()),
			StudentNumber: pointy.GetPointer(v.StudentNumber.String()),
		})
	}

	return resp, nil
}
