// Code generated by goctl. DO NOT EDIT.
// Source: sms.proto

package couponservice

import (
	"context"

	"zero-admin/rpc/sms/smsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	CouponAddReq                            = smsclient.CouponAddReq
	CouponAddResp                           = smsclient.CouponAddResp
	CouponDeleteReq                         = smsclient.CouponDeleteReq
	CouponDeleteResp                        = smsclient.CouponDeleteResp
	CouponFindByIdReq                       = smsclient.CouponFindByIdReq
	CouponFindByIdResp                      = smsclient.CouponFindByIdResp
	CouponFindByIdsReq                      = smsclient.CouponFindByIdsReq
	CouponFindByIdsResp                     = smsclient.CouponFindByIdsResp
	CouponHistoryAddReq                     = smsclient.CouponHistoryAddReq
	CouponHistoryAddResp                    = smsclient.CouponHistoryAddResp
	CouponHistoryDeleteReq                  = smsclient.CouponHistoryDeleteReq
	CouponHistoryDeleteResp                 = smsclient.CouponHistoryDeleteResp
	CouponHistoryListData                   = smsclient.CouponHistoryListData
	CouponHistoryListReq                    = smsclient.CouponHistoryListReq
	CouponHistoryListResp                   = smsclient.CouponHistoryListResp
	CouponHistoryUpdateReq                  = smsclient.CouponHistoryUpdateReq
	CouponHistoryUpdateResp                 = smsclient.CouponHistoryUpdateResp
	CouponListData                          = smsclient.CouponListData
	CouponListReq                           = smsclient.CouponListReq
	CouponListResp                          = smsclient.CouponListResp
	CouponProductCategoryRelationAddReq     = smsclient.CouponProductCategoryRelationAddReq
	CouponProductCategoryRelationAddResp    = smsclient.CouponProductCategoryRelationAddResp
	CouponProductCategoryRelationDeleteReq  = smsclient.CouponProductCategoryRelationDeleteReq
	CouponProductCategoryRelationDeleteResp = smsclient.CouponProductCategoryRelationDeleteResp
	CouponProductCategoryRelationListData   = smsclient.CouponProductCategoryRelationListData
	CouponProductCategoryRelationListReq    = smsclient.CouponProductCategoryRelationListReq
	CouponProductCategoryRelationListResp   = smsclient.CouponProductCategoryRelationListResp
	CouponProductCategoryRelationUpdateReq  = smsclient.CouponProductCategoryRelationUpdateReq
	CouponProductCategoryRelationUpdateResp = smsclient.CouponProductCategoryRelationUpdateResp
	CouponProductRelationAddReq             = smsclient.CouponProductRelationAddReq
	CouponProductRelationAddResp            = smsclient.CouponProductRelationAddResp
	CouponProductRelationDeleteReq          = smsclient.CouponProductRelationDeleteReq
	CouponProductRelationDeleteResp         = smsclient.CouponProductRelationDeleteResp
	CouponProductRelationListData           = smsclient.CouponProductRelationListData
	CouponProductRelationListReq            = smsclient.CouponProductRelationListReq
	CouponProductRelationListResp           = smsclient.CouponProductRelationListResp
	CouponProductRelationUpdateReq          = smsclient.CouponProductRelationUpdateReq
	CouponProductRelationUpdateResp         = smsclient.CouponProductRelationUpdateResp
	CouponUpdateReq                         = smsclient.CouponUpdateReq
	CouponUpdateResp                        = smsclient.CouponUpdateResp
	FlashPromotionAddReq                    = smsclient.FlashPromotionAddReq
	FlashPromotionAddResp                   = smsclient.FlashPromotionAddResp
	FlashPromotionDeleteReq                 = smsclient.FlashPromotionDeleteReq
	FlashPromotionDeleteResp                = smsclient.FlashPromotionDeleteResp
	FlashPromotionListByDateReq             = smsclient.FlashPromotionListByDateReq
	FlashPromotionListByDateResp            = smsclient.FlashPromotionListByDateResp
	FlashPromotionListData                  = smsclient.FlashPromotionListData
	FlashPromotionListReq                   = smsclient.FlashPromotionListReq
	FlashPromotionListResp                  = smsclient.FlashPromotionListResp
	FlashPromotionLogAddReq                 = smsclient.FlashPromotionLogAddReq
	FlashPromotionLogAddResp                = smsclient.FlashPromotionLogAddResp
	FlashPromotionLogDeleteReq              = smsclient.FlashPromotionLogDeleteReq
	FlashPromotionLogDeleteResp             = smsclient.FlashPromotionLogDeleteResp
	FlashPromotionLogListData               = smsclient.FlashPromotionLogListData
	FlashPromotionLogListReq                = smsclient.FlashPromotionLogListReq
	FlashPromotionLogListResp               = smsclient.FlashPromotionLogListResp
	FlashPromotionLogUpdateReq              = smsclient.FlashPromotionLogUpdateReq
	FlashPromotionLogUpdateResp             = smsclient.FlashPromotionLogUpdateResp
	FlashPromotionProductRelationAddReq     = smsclient.FlashPromotionProductRelationAddReq
	FlashPromotionProductRelationAddResp    = smsclient.FlashPromotionProductRelationAddResp
	FlashPromotionProductRelationDeleteReq  = smsclient.FlashPromotionProductRelationDeleteReq
	FlashPromotionProductRelationDeleteResp = smsclient.FlashPromotionProductRelationDeleteResp
	FlashPromotionProductRelationListData   = smsclient.FlashPromotionProductRelationListData
	FlashPromotionProductRelationListReq    = smsclient.FlashPromotionProductRelationListReq
	FlashPromotionProductRelationListResp   = smsclient.FlashPromotionProductRelationListResp
	FlashPromotionProductRelationUpdateReq  = smsclient.FlashPromotionProductRelationUpdateReq
	FlashPromotionProductRelationUpdateResp = smsclient.FlashPromotionProductRelationUpdateResp
	FlashPromotionSessionAddReq             = smsclient.FlashPromotionSessionAddReq
	FlashPromotionSessionAddResp            = smsclient.FlashPromotionSessionAddResp
	FlashPromotionSessionByTimeReq          = smsclient.FlashPromotionSessionByTimeReq
	FlashPromotionSessionByTimeResp         = smsclient.FlashPromotionSessionByTimeResp
	FlashPromotionSessionDeleteReq          = smsclient.FlashPromotionSessionDeleteReq
	FlashPromotionSessionDeleteResp         = smsclient.FlashPromotionSessionDeleteResp
	FlashPromotionSessionListData           = smsclient.FlashPromotionSessionListData
	FlashPromotionSessionListReq            = smsclient.FlashPromotionSessionListReq
	FlashPromotionSessionListResp           = smsclient.FlashPromotionSessionListResp
	FlashPromotionSessionUpdateReq          = smsclient.FlashPromotionSessionUpdateReq
	FlashPromotionSessionUpdateResp         = smsclient.FlashPromotionSessionUpdateResp
	FlashPromotionUpdateReq                 = smsclient.FlashPromotionUpdateReq
	FlashPromotionUpdateResp                = smsclient.FlashPromotionUpdateResp
	HomeAdvertiseAddReq                     = smsclient.HomeAdvertiseAddReq
	HomeAdvertiseAddResp                    = smsclient.HomeAdvertiseAddResp
	HomeAdvertiseDeleteReq                  = smsclient.HomeAdvertiseDeleteReq
	HomeAdvertiseDeleteResp                 = smsclient.HomeAdvertiseDeleteResp
	HomeAdvertiseListData                   = smsclient.HomeAdvertiseListData
	HomeAdvertiseListReq                    = smsclient.HomeAdvertiseListReq
	HomeAdvertiseListResp                   = smsclient.HomeAdvertiseListResp
	HomeAdvertiseUpdateReq                  = smsclient.HomeAdvertiseUpdateReq
	HomeAdvertiseUpdateResp                 = smsclient.HomeAdvertiseUpdateResp
	HomeBrandAddData                        = smsclient.HomeBrandAddData
	HomeBrandAddReq                         = smsclient.HomeBrandAddReq
	HomeBrandAddResp                        = smsclient.HomeBrandAddResp
	HomeBrandDeleteReq                      = smsclient.HomeBrandDeleteReq
	HomeBrandDeleteResp                     = smsclient.HomeBrandDeleteResp
	HomeBrandListData                       = smsclient.HomeBrandListData
	HomeBrandListReq                        = smsclient.HomeBrandListReq
	HomeBrandListResp                       = smsclient.HomeBrandListResp
	HomeBrandUpdateReq                      = smsclient.HomeBrandUpdateReq
	HomeBrandUpdateResp                     = smsclient.HomeBrandUpdateResp
	HomeNewProductAddData                   = smsclient.HomeNewProductAddData
	HomeNewProductAddReq                    = smsclient.HomeNewProductAddReq
	HomeNewProductAddResp                   = smsclient.HomeNewProductAddResp
	HomeNewProductDeleteReq                 = smsclient.HomeNewProductDeleteReq
	HomeNewProductDeleteResp                = smsclient.HomeNewProductDeleteResp
	HomeNewProductListData                  = smsclient.HomeNewProductListData
	HomeNewProductListReq                   = smsclient.HomeNewProductListReq
	HomeNewProductListResp                  = smsclient.HomeNewProductListResp
	HomeNewProductUpdateReq                 = smsclient.HomeNewProductUpdateReq
	HomeNewProductUpdateResp                = smsclient.HomeNewProductUpdateResp
	HomeRecommendProductAddData             = smsclient.HomeRecommendProductAddData
	HomeRecommendProductAddReq              = smsclient.HomeRecommendProductAddReq
	HomeRecommendProductAddResp             = smsclient.HomeRecommendProductAddResp
	HomeRecommendProductDeleteReq           = smsclient.HomeRecommendProductDeleteReq
	HomeRecommendProductDeleteResp          = smsclient.HomeRecommendProductDeleteResp
	HomeRecommendProductListData            = smsclient.HomeRecommendProductListData
	HomeRecommendProductListReq             = smsclient.HomeRecommendProductListReq
	HomeRecommendProductListResp            = smsclient.HomeRecommendProductListResp
	HomeRecommendProductUpdateReq           = smsclient.HomeRecommendProductUpdateReq
	HomeRecommendProductUpdateResp          = smsclient.HomeRecommendProductUpdateResp
	HomeRecommendSubjectAddData             = smsclient.HomeRecommendSubjectAddData
	HomeRecommendSubjectAddReq              = smsclient.HomeRecommendSubjectAddReq
	HomeRecommendSubjectAddResp             = smsclient.HomeRecommendSubjectAddResp
	HomeRecommendSubjectDeleteReq           = smsclient.HomeRecommendSubjectDeleteReq
	HomeRecommendSubjectDeleteResp          = smsclient.HomeRecommendSubjectDeleteResp
	HomeRecommendSubjectListData            = smsclient.HomeRecommendSubjectListData
	HomeRecommendSubjectListReq             = smsclient.HomeRecommendSubjectListReq
	HomeRecommendSubjectListResp            = smsclient.HomeRecommendSubjectListResp
	HomeRecommendSubjectUpdateReq           = smsclient.HomeRecommendSubjectUpdateReq
	HomeRecommendSubjectUpdateResp          = smsclient.HomeRecommendSubjectUpdateResp

	CouponService interface {
		CouponAdd(ctx context.Context, in *CouponAddReq, opts ...grpc.CallOption) (*CouponAddResp, error)
		CouponList(ctx context.Context, in *CouponListReq, opts ...grpc.CallOption) (*CouponListResp, error)
		CouponUpdate(ctx context.Context, in *CouponUpdateReq, opts ...grpc.CallOption) (*CouponUpdateResp, error)
		CouponDelete(ctx context.Context, in *CouponDeleteReq, opts ...grpc.CallOption) (*CouponDeleteResp, error)
		// 根据优惠券id查询优惠券
		CouponFindById(ctx context.Context, in *CouponFindByIdReq, opts ...grpc.CallOption) (*CouponFindByIdResp, error)
		// 根据优惠券ids查询优惠券
		CouponFindByIds(ctx context.Context, in *CouponFindByIdsReq, opts ...grpc.CallOption) (*CouponFindByIdsResp, error)
	}

	defaultCouponService struct {
		cli zrpc.Client
	}
)

func NewCouponService(cli zrpc.Client) CouponService {
	return &defaultCouponService{
		cli: cli,
	}
}

func (m *defaultCouponService) CouponAdd(ctx context.Context, in *CouponAddReq, opts ...grpc.CallOption) (*CouponAddResp, error) {
	client := smsclient.NewCouponServiceClient(m.cli.Conn())
	return client.CouponAdd(ctx, in, opts...)
}

func (m *defaultCouponService) CouponList(ctx context.Context, in *CouponListReq, opts ...grpc.CallOption) (*CouponListResp, error) {
	client := smsclient.NewCouponServiceClient(m.cli.Conn())
	return client.CouponList(ctx, in, opts...)
}

func (m *defaultCouponService) CouponUpdate(ctx context.Context, in *CouponUpdateReq, opts ...grpc.CallOption) (*CouponUpdateResp, error) {
	client := smsclient.NewCouponServiceClient(m.cli.Conn())
	return client.CouponUpdate(ctx, in, opts...)
}

func (m *defaultCouponService) CouponDelete(ctx context.Context, in *CouponDeleteReq, opts ...grpc.CallOption) (*CouponDeleteResp, error) {
	client := smsclient.NewCouponServiceClient(m.cli.Conn())
	return client.CouponDelete(ctx, in, opts...)
}

// 根据优惠券id查询优惠券
func (m *defaultCouponService) CouponFindById(ctx context.Context, in *CouponFindByIdReq, opts ...grpc.CallOption) (*CouponFindByIdResp, error) {
	client := smsclient.NewCouponServiceClient(m.cli.Conn())
	return client.CouponFindById(ctx, in, opts...)
}

// 根据优惠券ids查询优惠券
func (m *defaultCouponService) CouponFindByIds(ctx context.Context, in *CouponFindByIdsReq, opts ...grpc.CallOption) (*CouponFindByIdsResp, error) {
	client := smsclient.NewCouponServiceClient(m.cli.Conn())
	return client.CouponFindByIds(ctx, in, opts...)
}
