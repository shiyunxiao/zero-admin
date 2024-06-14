// Code generated by goctl. DO NOT EDIT.
// Source: ums.proto

package memberlevelservice

import (
	"context"

	"github.com/feihua/zero-admin/rpc/ums/umsclient"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddGrowthChangeHistoryReq                  = umsclient.AddGrowthChangeHistoryReq
	AddGrowthChangeHistoryResp                 = umsclient.AddGrowthChangeHistoryResp
	AddIntegrationChangeHistoryReq             = umsclient.AddIntegrationChangeHistoryReq
	AddIntegrationChangeHistoryResp            = umsclient.AddIntegrationChangeHistoryResp
	AddIntegrationConsumeSettingReq            = umsclient.AddIntegrationConsumeSettingReq
	AddIntegrationConsumeSettingResp           = umsclient.AddIntegrationConsumeSettingResp
	AddMemberBrandAttentionReq                 = umsclient.AddMemberBrandAttentionReq
	AddMemberBrandAttentionResp                = umsclient.AddMemberBrandAttentionResp
	AddMemberLevelReq                          = umsclient.AddMemberLevelReq
	AddMemberLevelResp                         = umsclient.AddMemberLevelResp
	AddMemberMemberTagRelationReq              = umsclient.AddMemberMemberTagRelationReq
	AddMemberMemberTagRelationResp             = umsclient.AddMemberMemberTagRelationResp
	AddMemberProductCategoryRelationReq        = umsclient.AddMemberProductCategoryRelationReq
	AddMemberProductCategoryRelationResp       = umsclient.AddMemberProductCategoryRelationResp
	AddMemberProductCollectionReq              = umsclient.AddMemberProductCollectionReq
	AddMemberProductCollectionResp             = umsclient.AddMemberProductCollectionResp
	AddMemberReadHistoryReq                    = umsclient.AddMemberReadHistoryReq
	AddMemberReadHistoryResp                   = umsclient.AddMemberReadHistoryResp
	AddMemberReceiveAddressReq                 = umsclient.AddMemberReceiveAddressReq
	AddMemberReceiveAddressResp                = umsclient.AddMemberReceiveAddressResp
	AddMemberReq                               = umsclient.AddMemberReq
	AddMemberResp                              = umsclient.AddMemberResp
	AddMemberRuleSettingReq                    = umsclient.AddMemberRuleSettingReq
	AddMemberRuleSettingResp                   = umsclient.AddMemberRuleSettingResp
	AddMemberStatisticsInfoReq                 = umsclient.AddMemberStatisticsInfoReq
	AddMemberStatisticsInfoResp                = umsclient.AddMemberStatisticsInfoResp
	AddMemberTagReq                            = umsclient.AddMemberTagReq
	AddMemberTagResp                           = umsclient.AddMemberTagResp
	AddMemberTaskReq                           = umsclient.AddMemberTaskReq
	AddMemberTaskResp                          = umsclient.AddMemberTaskResp
	DeleteGrowthChangeHistoryReq               = umsclient.DeleteGrowthChangeHistoryReq
	DeleteGrowthChangeHistoryResp              = umsclient.DeleteGrowthChangeHistoryResp
	DeleteIntegrationChangeHistoryReq          = umsclient.DeleteIntegrationChangeHistoryReq
	DeleteIntegrationChangeHistoryResp         = umsclient.DeleteIntegrationChangeHistoryResp
	DeleteIntegrationConsumeSettingReq         = umsclient.DeleteIntegrationConsumeSettingReq
	DeleteIntegrationConsumeSettingResp        = umsclient.DeleteIntegrationConsumeSettingResp
	DeleteMemberBrandAttentionReq              = umsclient.DeleteMemberBrandAttentionReq
	DeleteMemberBrandAttentionResp             = umsclient.DeleteMemberBrandAttentionResp
	DeleteMemberLevelReq                       = umsclient.DeleteMemberLevelReq
	DeleteMemberLevelResp                      = umsclient.DeleteMemberLevelResp
	DeleteMemberLoginLogReq                    = umsclient.DeleteMemberLoginLogReq
	DeleteMemberLoginLogResp                   = umsclient.DeleteMemberLoginLogResp
	DeleteMemberMemberTagRelationReq           = umsclient.DeleteMemberMemberTagRelationReq
	DeleteMemberMemberTagRelationResp          = umsclient.DeleteMemberMemberTagRelationResp
	DeleteMemberProductCollectionReq           = umsclient.DeleteMemberProductCollectionReq
	DeleteMemberProductCollectionResp          = umsclient.DeleteMemberProductCollectionResp
	DeleteMemberReadHistoryReq                 = umsclient.DeleteMemberReadHistoryReq
	DeleteMemberReadHistoryResp                = umsclient.DeleteMemberReadHistoryResp
	DeleteMemberReceiveAddressReq              = umsclient.DeleteMemberReceiveAddressReq
	DeleteMemberReceiveAddressResp             = umsclient.DeleteMemberReceiveAddressResp
	DeleteMemberReq                            = umsclient.DeleteMemberReq
	DeleteMemberResp                           = umsclient.DeleteMemberResp
	DeleteMemberRuleSettingReq                 = umsclient.DeleteMemberRuleSettingReq
	DeleteMemberRuleSettingResp                = umsclient.DeleteMemberRuleSettingResp
	DeleteMemberTagReq                         = umsclient.DeleteMemberTagReq
	DeleteMemberTagResp                        = umsclient.DeleteMemberTagResp
	DeleteMemberTaskReq                        = umsclient.DeleteMemberTaskReq
	DeleteMemberTaskResp                       = umsclient.DeleteMemberTaskResp
	GrowthChangeHistoryListData                = umsclient.GrowthChangeHistoryListData
	IntegrationChangeHistoryListData           = umsclient.IntegrationChangeHistoryListData
	IntegrationConsumeSettingListData          = umsclient.IntegrationConsumeSettingListData
	MemberBrandAttentionListData               = umsclient.MemberBrandAttentionListData
	MemberLevelListData                        = umsclient.MemberLevelListData
	MemberListData                             = umsclient.MemberListData
	MemberLoginLogListData                     = umsclient.MemberLoginLogListData
	MemberLoginReq                             = umsclient.MemberLoginReq
	MemberLoginResp                            = umsclient.MemberLoginResp
	MemberMemberTagRelationListData            = umsclient.MemberMemberTagRelationListData
	MemberProductCategoryRelationListData      = umsclient.MemberProductCategoryRelationListData
	MemberProductCollectionListData            = umsclient.MemberProductCollectionListData
	MemberReadHistoryListData                  = umsclient.MemberReadHistoryListData
	MemberReceiveAddressListData               = umsclient.MemberReceiveAddressListData
	MemberRuleSettingListData                  = umsclient.MemberRuleSettingListData
	MemberTagListData                          = umsclient.MemberTagListData
	MemberTaskListData                         = umsclient.MemberTaskListData
	QueryGrowthChangeHistoryDetailReq          = umsclient.QueryGrowthChangeHistoryDetailReq
	QueryGrowthChangeHistoryDetailResp         = umsclient.QueryGrowthChangeHistoryDetailResp
	QueryGrowthChangeHistoryListReq            = umsclient.QueryGrowthChangeHistoryListReq
	QueryGrowthChangeHistoryListResp           = umsclient.QueryGrowthChangeHistoryListResp
	QueryIntegrationChangeHistoryDetailReq     = umsclient.QueryIntegrationChangeHistoryDetailReq
	QueryIntegrationChangeHistoryDetailResp    = umsclient.QueryIntegrationChangeHistoryDetailResp
	QueryIntegrationChangeHistoryListReq       = umsclient.QueryIntegrationChangeHistoryListReq
	QueryIntegrationChangeHistoryListResp      = umsclient.QueryIntegrationChangeHistoryListResp
	QueryIntegrationConsumeSettingDetailReq    = umsclient.QueryIntegrationConsumeSettingDetailReq
	QueryIntegrationConsumeSettingDetailResp   = umsclient.QueryIntegrationConsumeSettingDetailResp
	QueryIntegrationConsumeSettingListReq      = umsclient.QueryIntegrationConsumeSettingListReq
	QueryIntegrationConsumeSettingListResp     = umsclient.QueryIntegrationConsumeSettingListResp
	QueryMemberBrandAttentionDetailReq         = umsclient.QueryMemberBrandAttentionDetailReq
	QueryMemberBrandAttentionDetailResp        = umsclient.QueryMemberBrandAttentionDetailResp
	QueryMemberBrandAttentionListReq           = umsclient.QueryMemberBrandAttentionListReq
	QueryMemberBrandAttentionListResp          = umsclient.QueryMemberBrandAttentionListResp
	QueryMemberDetailReq                       = umsclient.QueryMemberDetailReq
	QueryMemberDetailResp                      = umsclient.QueryMemberDetailResp
	QueryMemberLevelDetailReq                  = umsclient.QueryMemberLevelDetailReq
	QueryMemberLevelDetailResp                 = umsclient.QueryMemberLevelDetailResp
	QueryMemberLevelListReq                    = umsclient.QueryMemberLevelListReq
	QueryMemberLevelListResp                   = umsclient.QueryMemberLevelListResp
	QueryMemberListReq                         = umsclient.QueryMemberListReq
	QueryMemberListResp                        = umsclient.QueryMemberListResp
	QueryMemberLoginLogListReq                 = umsclient.QueryMemberLoginLogListReq
	QueryMemberLoginLogListResp                = umsclient.QueryMemberLoginLogListResp
	QueryMemberMemberTagRelationListReq        = umsclient.QueryMemberMemberTagRelationListReq
	QueryMemberMemberTagRelationListResp       = umsclient.QueryMemberMemberTagRelationListResp
	QueryMemberProductCategoryRelationListReq  = umsclient.QueryMemberProductCategoryRelationListReq
	QueryMemberProductCategoryRelationListResp = umsclient.QueryMemberProductCategoryRelationListResp
	QueryMemberProductCollectionDetailReq      = umsclient.QueryMemberProductCollectionDetailReq
	QueryMemberProductCollectionDetailResp     = umsclient.QueryMemberProductCollectionDetailResp
	QueryMemberProductCollectionListReq        = umsclient.QueryMemberProductCollectionListReq
	QueryMemberProductCollectionListResp       = umsclient.QueryMemberProductCollectionListResp
	QueryMemberReadHistoryDetailReq            = umsclient.QueryMemberReadHistoryDetailReq
	QueryMemberReadHistoryDetailResp           = umsclient.QueryMemberReadHistoryDetailResp
	QueryMemberReadHistoryListReq              = umsclient.QueryMemberReadHistoryListReq
	QueryMemberReadHistoryListResp             = umsclient.QueryMemberReadHistoryListResp
	QueryMemberReceiveAddressDetailReq         = umsclient.QueryMemberReceiveAddressDetailReq
	QueryMemberReceiveAddressDetailResp        = umsclient.QueryMemberReceiveAddressDetailResp
	QueryMemberReceiveAddressListReq           = umsclient.QueryMemberReceiveAddressListReq
	QueryMemberReceiveAddressListResp          = umsclient.QueryMemberReceiveAddressListResp
	QueryMemberRuleSettingDetailReq            = umsclient.QueryMemberRuleSettingDetailReq
	QueryMemberRuleSettingDetailResp           = umsclient.QueryMemberRuleSettingDetailResp
	QueryMemberRuleSettingListReq              = umsclient.QueryMemberRuleSettingListReq
	QueryMemberRuleSettingListResp             = umsclient.QueryMemberRuleSettingListResp
	QueryMemberStatisticsInfoDetailReq         = umsclient.QueryMemberStatisticsInfoDetailReq
	QueryMemberStatisticsInfoDetailResp        = umsclient.QueryMemberStatisticsInfoDetailResp
	QueryMemberStatisticsInfoListData          = umsclient.QueryMemberStatisticsInfoListData
	QueryMemberStatisticsInfoListReq           = umsclient.QueryMemberStatisticsInfoListReq
	QueryMemberStatisticsInfoListResp          = umsclient.QueryMemberStatisticsInfoListResp
	QueryMemberTagDetailReq                    = umsclient.QueryMemberTagDetailReq
	QueryMemberTagDetailResp                   = umsclient.QueryMemberTagDetailResp
	QueryMemberTagListReq                      = umsclient.QueryMemberTagListReq
	QueryMemberTagListResp                     = umsclient.QueryMemberTagListResp
	QueryMemberTaskDetailReq                   = umsclient.QueryMemberTaskDetailReq
	QueryMemberTaskDetailResp                  = umsclient.QueryMemberTaskDetailResp
	QueryMemberTaskListReq                     = umsclient.QueryMemberTaskListReq
	QueryMemberTaskListResp                    = umsclient.QueryMemberTaskListResp
	UpdateIntegrationConsumeSettingReq         = umsclient.UpdateIntegrationConsumeSettingReq
	UpdateIntegrationConsumeSettingResp        = umsclient.UpdateIntegrationConsumeSettingResp
	UpdateIntegrationConsumeSettingStatusReq   = umsclient.UpdateIntegrationConsumeSettingStatusReq
	UpdateIntegrationConsumeSettingStatusResp  = umsclient.UpdateIntegrationConsumeSettingStatusResp
	UpdateMemberIntegrationReq                 = umsclient.UpdateMemberIntegrationReq
	UpdateMemberIntegrationResp                = umsclient.UpdateMemberIntegrationResp
	UpdateMemberLevelReq                       = umsclient.UpdateMemberLevelReq
	UpdateMemberLevelResp                      = umsclient.UpdateMemberLevelResp
	UpdateMemberLevelStatusReq                 = umsclient.UpdateMemberLevelStatusReq
	UpdateMemberLevelStatusResp                = umsclient.UpdateMemberLevelStatusResp
	UpdateMemberReceiveAddressReq              = umsclient.UpdateMemberReceiveAddressReq
	UpdateMemberReceiveAddressResp             = umsclient.UpdateMemberReceiveAddressResp
	UpdateMemberReceiveAddressStatusReq        = umsclient.UpdateMemberReceiveAddressStatusReq
	UpdateMemberReceiveAddressStatusResp       = umsclient.UpdateMemberReceiveAddressStatusResp
	UpdateMemberReq                            = umsclient.UpdateMemberReq
	UpdateMemberResp                           = umsclient.UpdateMemberResp
	UpdateMemberRuleSettingReq                 = umsclient.UpdateMemberRuleSettingReq
	UpdateMemberRuleSettingResp                = umsclient.UpdateMemberRuleSettingResp
	UpdateMemberRuleSettingStatusReq           = umsclient.UpdateMemberRuleSettingStatusReq
	UpdateMemberRuleSettingStatusResp          = umsclient.UpdateMemberRuleSettingStatusResp
	UpdateMemberStatusReq                      = umsclient.UpdateMemberStatusReq
	UpdateMemberStatusResp                     = umsclient.UpdateMemberStatusResp
	UpdateMemberTagReq                         = umsclient.UpdateMemberTagReq
	UpdateMemberTagResp                        = umsclient.UpdateMemberTagResp
	UpdateMemberTagStatusReq                   = umsclient.UpdateMemberTagStatusReq
	UpdateMemberTagStatusResp                  = umsclient.UpdateMemberTagStatusResp
	UpdateMemberTaskReq                        = umsclient.UpdateMemberTaskReq
	UpdateMemberTaskResp                       = umsclient.UpdateMemberTaskResp
	UpdateMemberTaskStatusReq                  = umsclient.UpdateMemberTaskStatusReq
	UpdateMemberTaskStatusResp                 = umsclient.UpdateMemberTaskStatusResp

	MemberLevelService interface {
		// 添加会员等级表
		AddMemberLevel(ctx context.Context, in *AddMemberLevelReq, opts ...grpc.CallOption) (*AddMemberLevelResp, error)
		// 删除会员等级表
		DeleteMemberLevel(ctx context.Context, in *DeleteMemberLevelReq, opts ...grpc.CallOption) (*DeleteMemberLevelResp, error)
		// 更新会员等级表
		UpdateMemberLevel(ctx context.Context, in *UpdateMemberLevelReq, opts ...grpc.CallOption) (*UpdateMemberLevelResp, error)
		// 更新会员等级表状态
		UpdateMemberLevelStatus(ctx context.Context, in *UpdateMemberLevelStatusReq, opts ...grpc.CallOption) (*UpdateMemberLevelStatusResp, error)
		// 查询会员等级表详情
		QueryMemberLevelDetail(ctx context.Context, in *QueryMemberLevelDetailReq, opts ...grpc.CallOption) (*QueryMemberLevelDetailResp, error)
		// 查询会员等级表列表
		QueryMemberLevelList(ctx context.Context, in *QueryMemberLevelListReq, opts ...grpc.CallOption) (*QueryMemberLevelListResp, error)
	}

	defaultMemberLevelService struct {
		cli zrpc.Client
	}
)

func NewMemberLevelService(cli zrpc.Client) MemberLevelService {
	return &defaultMemberLevelService{
		cli: cli,
	}
}

// 添加会员等级表
func (m *defaultMemberLevelService) AddMemberLevel(ctx context.Context, in *AddMemberLevelReq, opts ...grpc.CallOption) (*AddMemberLevelResp, error) {
	client := umsclient.NewMemberLevelServiceClient(m.cli.Conn())
	return client.AddMemberLevel(ctx, in, opts...)
}

// 删除会员等级表
func (m *defaultMemberLevelService) DeleteMemberLevel(ctx context.Context, in *DeleteMemberLevelReq, opts ...grpc.CallOption) (*DeleteMemberLevelResp, error) {
	client := umsclient.NewMemberLevelServiceClient(m.cli.Conn())
	return client.DeleteMemberLevel(ctx, in, opts...)
}

// 更新会员等级表
func (m *defaultMemberLevelService) UpdateMemberLevel(ctx context.Context, in *UpdateMemberLevelReq, opts ...grpc.CallOption) (*UpdateMemberLevelResp, error) {
	client := umsclient.NewMemberLevelServiceClient(m.cli.Conn())
	return client.UpdateMemberLevel(ctx, in, opts...)
}

// 更新会员等级表状态
func (m *defaultMemberLevelService) UpdateMemberLevelStatus(ctx context.Context, in *UpdateMemberLevelStatusReq, opts ...grpc.CallOption) (*UpdateMemberLevelStatusResp, error) {
	client := umsclient.NewMemberLevelServiceClient(m.cli.Conn())
	return client.UpdateMemberLevelStatus(ctx, in, opts...)
}

// 查询会员等级表详情
func (m *defaultMemberLevelService) QueryMemberLevelDetail(ctx context.Context, in *QueryMemberLevelDetailReq, opts ...grpc.CallOption) (*QueryMemberLevelDetailResp, error) {
	client := umsclient.NewMemberLevelServiceClient(m.cli.Conn())
	return client.QueryMemberLevelDetail(ctx, in, opts...)
}

// 查询会员等级表列表
func (m *defaultMemberLevelService) QueryMemberLevelList(ctx context.Context, in *QueryMemberLevelListReq, opts ...grpc.CallOption) (*QueryMemberLevelListResp, error) {
	client := umsclient.NewMemberLevelServiceClient(m.cli.Conn())
	return client.QueryMemberLevelList(ctx, in, opts...)
}
