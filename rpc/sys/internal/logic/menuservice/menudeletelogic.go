package menuservicelogic

import (
	"context"
	"errors"
	"github.com/feihua/zero-admin/rpc/sys/gen/query"
	"github.com/feihua/zero-admin/rpc/sys/internal/svc"
	"github.com/feihua/zero-admin/rpc/sys/sysclient"
	"github.com/zeromicro/go-zero/core/logc"

	"github.com/zeromicro/go-zero/core/logx"
)

// MenuDeleteLogic 删除菜单
/*
Author: LiuFeiHua
Date: 2023/12/18 15:45
*/
type MenuDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMenuDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MenuDeleteLogic {
	return &MenuDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// MenuDelete 删除菜单
func (l *MenuDeleteLogic) MenuDelete(in *sysclient.MenuDeleteReq) (*sysclient.MenuDeleteResp, error) {
	q := query.SysMenu
	_, err := q.WithContext(l.ctx).Where(q.ID.In(in.Ids...)).Delete()
	if err != nil {
		logc.Errorf(l.ctx, "删除菜单失败,参数:%+v,异常:%s", in, err.Error())
		return nil, errors.New("删除菜单失败")
	}

	return &sysclient.MenuDeleteResp{}, nil
}
