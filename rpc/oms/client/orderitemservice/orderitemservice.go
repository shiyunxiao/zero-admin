// Code generated by goctl. DO NOT EDIT.
// Source: oms.proto

package orderitemservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/oms/omsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddCartItemReq                    = omsclient.AddCartItemReq
	AddCartItemResp                   = omsclient.AddCartItemResp
	AddCompanyAddressReq              = omsclient.AddCompanyAddressReq
	AddCompanyAddressResp             = omsclient.AddCompanyAddressResp
	AddOrderItemReq                   = omsclient.AddOrderItemReq
	AddOrderItemResp                  = omsclient.AddOrderItemResp
	AddOrderOperateHistoryReq         = omsclient.AddOrderOperateHistoryReq
	AddOrderOperateHistoryResp        = omsclient.AddOrderOperateHistoryResp
	AddOrderReturnApplyReq            = omsclient.AddOrderReturnApplyReq
	AddOrderReturnApplyResp           = omsclient.AddOrderReturnApplyResp
	AddOrderReturnReasonReq           = omsclient.AddOrderReturnReasonReq
	AddOrderReturnReasonResp          = omsclient.AddOrderReturnReasonResp
	AddOrderSettingReq                = omsclient.AddOrderSettingReq
	AddOrderSettingResp               = omsclient.AddOrderSettingResp
	CartItemListData                  = omsclient.CartItemListData
	CloseOrderReq                     = omsclient.CloseOrderReq
	CloseOrderResp                    = omsclient.CloseOrderResp
	CompanyAddressListData            = omsclient.CompanyAddressListData
	DeleteCartItemReq                 = omsclient.DeleteCartItemReq
	DeleteCartItemResp                = omsclient.DeleteCartItemResp
	DeleteCompanyAddressReq           = omsclient.DeleteCompanyAddressReq
	DeleteCompanyAddressResp          = omsclient.DeleteCompanyAddressResp
	DeleteOrderOperateHistoryReq      = omsclient.DeleteOrderOperateHistoryReq
	DeleteOrderOperateHistoryResp     = omsclient.DeleteOrderOperateHistoryResp
	DeleteOrderReturnApplyReq         = omsclient.DeleteOrderReturnApplyReq
	DeleteOrderReturnApplyResp        = omsclient.DeleteOrderReturnApplyResp
	DeleteOrderReturnReasonReq        = omsclient.DeleteOrderReturnReasonReq
	DeleteOrderReturnReasonResp       = omsclient.DeleteOrderReturnReasonResp
	DeleteOrderSettingReq             = omsclient.DeleteOrderSettingReq
	DeleteOrderSettingResp            = omsclient.DeleteOrderSettingResp
	DeliveryReq                       = omsclient.DeliveryReq
	DeliveryResp                      = omsclient.DeliveryResp
	OrderAddReq                       = omsclient.OrderAddReq
	OrderAddResp                      = omsclient.OrderAddResp
	OrderCancelReq                    = omsclient.OrderCancelReq
	OrderCancelResp                   = omsclient.OrderCancelResp
	OrderConfirmReq                   = omsclient.OrderConfirmReq
	OrderConfirmResp                  = omsclient.OrderConfirmResp
	OrderDeleteByIdReq                = omsclient.OrderDeleteByIdReq
	OrderDeleteReq                    = omsclient.OrderDeleteReq
	OrderDeleteResp                   = omsclient.OrderDeleteResp
	OrderDetailReq                    = omsclient.OrderDetailReq
	OrderDetailResp                   = omsclient.OrderDetailResp
	OrderItemData                     = omsclient.OrderItemData
	OrderItemListData                 = omsclient.OrderItemListData
	OrderListByMemberIdReq            = omsclient.OrderListByMemberIdReq
	OrderListByMemberIdResp           = omsclient.OrderListByMemberIdResp
	OrderListData                     = omsclient.OrderListData
	OrderListReq                      = omsclient.OrderListReq
	OrderListResp                     = omsclient.OrderListResp
	OrderOperateHistoryData           = omsclient.OrderOperateHistoryData
	OrderOperateHistoryListData       = omsclient.OrderOperateHistoryListData
	OrderRefundReq                    = omsclient.OrderRefundReq
	OrderRefundResp                   = omsclient.OrderRefundResp
	OrderReturnApplyListData          = omsclient.OrderReturnApplyListData
	OrderReturnReasonListData         = omsclient.OrderReturnReasonListData
	OrderSettingListData              = omsclient.OrderSettingListData
	OrderUpdateReq                    = omsclient.OrderUpdateReq
	OrderUpdateResp                   = omsclient.OrderUpdateResp
	QueryCartItemDetailReq            = omsclient.QueryCartItemDetailReq
	QueryCartItemDetailResp           = omsclient.QueryCartItemDetailResp
	QueryCartItemListReq              = omsclient.QueryCartItemListReq
	QueryCartItemListResp             = omsclient.QueryCartItemListResp
	QueryCompanyAddressDetailReq      = omsclient.QueryCompanyAddressDetailReq
	QueryCompanyAddressDetailResp     = omsclient.QueryCompanyAddressDetailResp
	QueryCompanyAddressListReq        = omsclient.QueryCompanyAddressListReq
	QueryCompanyAddressListResp       = omsclient.QueryCompanyAddressListResp
	QueryOrderItemDetailReq           = omsclient.QueryOrderItemDetailReq
	QueryOrderItemDetailResp          = omsclient.QueryOrderItemDetailResp
	QueryOrderItemListReq             = omsclient.QueryOrderItemListReq
	QueryOrderItemListResp            = omsclient.QueryOrderItemListResp
	QueryOrderListReq                 = omsclient.QueryOrderListReq
	QueryOrderOperateHistoryListReq   = omsclient.QueryOrderOperateHistoryListReq
	QueryOrderOperateHistoryListResp  = omsclient.QueryOrderOperateHistoryListResp
	QueryOrderReturnApplyDetailReq    = omsclient.QueryOrderReturnApplyDetailReq
	QueryOrderReturnApplyDetailResp   = omsclient.QueryOrderReturnApplyDetailResp
	QueryOrderReturnApplyListReq      = omsclient.QueryOrderReturnApplyListReq
	QueryOrderReturnApplyListResp     = omsclient.QueryOrderReturnApplyListResp
	QueryOrderReturnReasonDetailReq   = omsclient.QueryOrderReturnReasonDetailReq
	QueryOrderReturnReasonDetailResp  = omsclient.QueryOrderReturnReasonDetailResp
	QueryOrderReturnReasonListReq     = omsclient.QueryOrderReturnReasonListReq
	QueryOrderReturnReasonListResp    = omsclient.QueryOrderReturnReasonListResp
	QueryOrderSettingDetailReq        = omsclient.QueryOrderSettingDetailReq
	QueryOrderSettingDetailResp       = omsclient.QueryOrderSettingDetailResp
	QueryOrderSettingListReq          = omsclient.QueryOrderSettingListReq
	QueryOrderSettingListResp         = omsclient.QueryOrderSettingListResp
	ReleaseSkuStockLockData           = omsclient.ReleaseSkuStockLockData
	UpdateCartItemQuantityReq         = omsclient.UpdateCartItemQuantityReq
	UpdateCartItemQuantityResp        = omsclient.UpdateCartItemQuantityResp
	UpdateCartItemReq                 = omsclient.UpdateCartItemReq
	UpdateCartItemResp                = omsclient.UpdateCartItemResp
	UpdateCompanyAddressReq           = omsclient.UpdateCompanyAddressReq
	UpdateCompanyAddressResp          = omsclient.UpdateCompanyAddressResp
	UpdateCompanyAddressStatusReq     = omsclient.UpdateCompanyAddressStatusReq
	UpdateCompanyAddressStatusResp    = omsclient.UpdateCompanyAddressStatusResp
	UpdateMoneyInfoReq                = omsclient.UpdateMoneyInfoReq
	UpdateMoneyInfoResp               = omsclient.UpdateMoneyInfoResp
	UpdateNoteReq                     = omsclient.UpdateNoteReq
	UpdateNoteResp                    = omsclient.UpdateNoteResp
	UpdateOrderReturnApplyReq         = omsclient.UpdateOrderReturnApplyReq
	UpdateOrderReturnApplyResp        = omsclient.UpdateOrderReturnApplyResp
	UpdateOrderReturnReasonReq        = omsclient.UpdateOrderReturnReasonReq
	UpdateOrderReturnReasonResp       = omsclient.UpdateOrderReturnReasonResp
	UpdateOrderReturnReasonStatusReq  = omsclient.UpdateOrderReturnReasonStatusReq
	UpdateOrderReturnReasonStatusResp = omsclient.UpdateOrderReturnReasonStatusResp
	UpdateOrderSettingReq             = omsclient.UpdateOrderSettingReq
	UpdateOrderSettingResp            = omsclient.UpdateOrderSettingResp
	UpdateOrderStatusByOutTradeNoReq  = omsclient.UpdateOrderStatusByOutTradeNoReq
	UpdateOrderStatusByOutTradeNoResp = omsclient.UpdateOrderStatusByOutTradeNoResp
	UpdateReceiverInfoReq             = omsclient.UpdateReceiverInfoReq
	UpdateReceiverInfoResp            = omsclient.UpdateReceiverInfoResp

	OrderItemService interface {
		// 添加订单中所包含的商品
		AddOrderItem(ctx context.Context, in *AddOrderItemReq, opts ...grpc.CallOption) (*AddOrderItemResp, error)
		// 查询订单中所包含的商品详情
		QueryOrderItemDetail(ctx context.Context, in *QueryOrderItemDetailReq, opts ...grpc.CallOption) (*QueryOrderItemDetailResp, error)
		// 查询订单中所包含的商品列表
		QueryOrderItemList(ctx context.Context, in *QueryOrderItemListReq, opts ...grpc.CallOption) (*QueryOrderItemListResp, error)
	}

	defaultOrderItemService struct {
		cli zrpc.Client
	}
)

func NewOrderItemService(cli zrpc.Client) OrderItemService {
	return &defaultOrderItemService{
		cli: cli,
	}
}

// 添加订单中所包含的商品
func (m *defaultOrderItemService) AddOrderItem(ctx context.Context, in *AddOrderItemReq, opts ...grpc.CallOption) (*AddOrderItemResp, error) {
	client := omsclient.NewOrderItemServiceClient(m.cli.Conn())
	return client.AddOrderItem(ctx, in, opts...)
}

// 查询订单中所包含的商品详情
func (m *defaultOrderItemService) QueryOrderItemDetail(ctx context.Context, in *QueryOrderItemDetailReq, opts ...grpc.CallOption) (*QueryOrderItemDetailResp, error) {
	client := omsclient.NewOrderItemServiceClient(m.cli.Conn())
	return client.QueryOrderItemDetail(ctx, in, opts...)
}

// 查询订单中所包含的商品列表
func (m *defaultOrderItemService) QueryOrderItemList(ctx context.Context, in *QueryOrderItemListReq, opts ...grpc.CallOption) (*QueryOrderItemListResp, error) {
	client := omsclient.NewOrderItemServiceClient(m.cli.Conn())
	return client.QueryOrderItemList(ctx, in, opts...)
}
