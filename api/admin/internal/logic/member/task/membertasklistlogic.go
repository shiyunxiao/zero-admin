package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/ums/umsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MemberTaskListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMemberTaskListLogic(ctx context.Context, svcCtx *svc.ServiceContext) MemberTaskListLogic {
	return MemberTaskListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MemberTaskListLogic) MemberTaskList(req types.ListMemberTaskReq) (*types.ListMemberTaskResp, error) {
	resp, err := l.svcCtx.MemberTaskService.MemberTaskList(l.ctx, &umsclient.MemberTaskListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询会员任务列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询会员任务失败")
	}

	var list []*types.ListMemberTaskData

	for _, item := range resp.List {
		list = append(list, &types.ListMemberTaskData{
			Id:           item.Id,
			Name:         item.TaskName,
			Growth:       item.TaskGrowth,
			Intergration: item.TaskIntegral,
			Type:         item.TaskType,
		})
	}

	return &types.ListMemberTaskResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询会员任务成功",
	}, nil
}
