// Code generated by goctl. DO NOT EDIT.
// Source: oms.proto

package server

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/internal/logic/orderreturnapplyservice"
	"github.com/feihua/zero-admin/rpc/oms/internal/svc"
	"github.com/feihua/zero-admin/rpc/oms/omsclient"
)

type OrderReturnApplyServiceServer struct {
	svcCtx *svc.ServiceContext
	omsclient.UnimplementedOrderReturnApplyServiceServer
}

func NewOrderReturnApplyServiceServer(svcCtx *svc.ServiceContext) *OrderReturnApplyServiceServer {
	return &OrderReturnApplyServiceServer{
		svcCtx: svcCtx,
	}
}

// 添加订单退货申请
func (s *OrderReturnApplyServiceServer) AddOrderReturnApply(ctx context.Context, in *omsclient.AddOrderReturnApplyReq) (*omsclient.AddOrderReturnApplyResp, error) {
	l := orderreturnapplyservicelogic.NewAddOrderReturnApplyLogic(ctx, s.svcCtx)
	return l.AddOrderReturnApply(in)
}

// 删除订单退货申请
func (s *OrderReturnApplyServiceServer) DeleteOrderReturnApply(ctx context.Context, in *omsclient.DeleteOrderReturnApplyReq) (*omsclient.DeleteOrderReturnApplyResp, error) {
	l := orderreturnapplyservicelogic.NewDeleteOrderReturnApplyLogic(ctx, s.svcCtx)
	return l.DeleteOrderReturnApply(in)
}

// 更新订单退货申请
func (s *OrderReturnApplyServiceServer) UpdateOrderReturnApply(ctx context.Context, in *omsclient.UpdateOrderReturnApplyReq) (*omsclient.UpdateOrderReturnApplyResp, error) {
	l := orderreturnapplyservicelogic.NewUpdateOrderReturnApplyLogic(ctx, s.svcCtx)
	return l.UpdateOrderReturnApply(in)
}

// 查询订单退货申请详情
func (s *OrderReturnApplyServiceServer) QueryOrderReturnApplyDetail(ctx context.Context, in *omsclient.QueryOrderReturnApplyDetailReq) (*omsclient.QueryOrderReturnApplyDetailResp, error) {
	l := orderreturnapplyservicelogic.NewQueryOrderReturnApplyDetailLogic(ctx, s.svcCtx)
	return l.QueryOrderReturnApplyDetail(in)
}

// 查询订单退货申请列表
func (s *OrderReturnApplyServiceServer) QueryOrderReturnApplyList(ctx context.Context, in *omsclient.QueryOrderReturnApplyListReq) (*omsclient.QueryOrderReturnApplyListResp, error) {
	l := orderreturnapplyservicelogic.NewQueryOrderReturnApplyListLogic(ctx, s.svcCtx)
	return l.QueryOrderReturnApplyList(in)
}
