package logic

import (
	"context"
	"github.com/feihua/zero-admin/api/admin/internal/common/errorx"
	"github.com/feihua/zero-admin/rpc/sms/smsclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/feihua/zero-admin/api/admin/internal/svc"
	"github.com/feihua/zero-admin/api/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FlashPromotionSessionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFlashPromotionSessionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) FlashPromotionSessionListLogic {
	return FlashPromotionSessionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FlashPromotionSessionListLogic) FlashPromotionSessionList(req types.ListFlashPromotionSessionReq) (*types.ListFlashPromotionSessionResp, error) {
	resp, err := l.svcCtx.FlashPromotionSessionService.FlashPromotionSessionList(l.ctx, &smsclient.FlashPromotionSessionListReq{
		Current:  req.Current,
		PageSize: req.PageSize,
	})

	if err != nil {
		logc.Errorf(l.ctx, "参数: %+v,查询限时购场次表列表异常:%s", req, err.Error())
		return nil, errorx.NewDefaultError("查询限时购场次表失败")
	}

	var list []*types.ListFlashPromotionSessionData

	for _, item := range resp.List {
		list = append(list, &types.ListFlashPromotionSessionData{
			Id:         item.Id,
			Name:       item.Name,
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			Status:     item.Status,
			CreateTime: item.CreateTime,
		})
	}

	return &types.ListFlashPromotionSessionResp{
		Current:  req.Current,
		Data:     list,
		PageSize: req.PageSize,
		Success:  true,
		Total:    resp.Total,
		Code:     "000000",
		Message:  "查询限时购场次表成功",
	}, nil
}
