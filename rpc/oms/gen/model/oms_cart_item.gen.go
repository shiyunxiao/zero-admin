// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameOmsCartItem = "oms_cart_item"

// OmsCartItem 购物车表
type OmsCartItem struct {
	ID                int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	ProductID         int64      `gorm:"column:product_id;not null;comment:商品id" json:"product_id"`                                                             // 商品id
	ProductSkuID      int64      `gorm:"column:product_sku_id;not null;comment:商品库存id" json:"product_sku_id"`                                                   // 商品库存id
	MemberID          int64      `gorm:"column:member_id;not null;comment:会员id" json:"member_id"`                                                               // 会员id
	Quantity          int32      `gorm:"column:quantity;not null;comment:购买数量" json:"quantity"`                                                                 // 购买数量
	Price             int64      `gorm:"column:price;not null;comment:添加到购物车的价格" json:"price"`                                                                  // 添加到购物车的价格
	ProductPic        string     `gorm:"column:product_pic;not null;comment:商品主图" json:"product_pic"`                                                           // 商品主图
	ProductName       string     `gorm:"column:product_name;not null;comment:商品名称" json:"product_name"`                                                         // 商品名称
	ProductSubTitle   string     `gorm:"column:product_sub_title;not null;comment:商品副标题（卖点）" json:"product_sub_title"`                                          // 商品副标题（卖点）
	ProductSkuCode    string     `gorm:"column:product_sku_code;not null;comment:商品sku条码" json:"product_sku_code"`                                              // 商品sku条码
	MemberNickname    string     `gorm:"column:member_nickname;not null;comment:会员昵称" json:"member_nickname"`                                                   // 会员昵称
	CreateTime        time.Time  `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP;comment:创建时间" json:"create_time"`                                 // 创建时间
	UpdateTime        *time.Time `gorm:"column:update_time;comment:更新时间" json:"update_time"`                                                                    // 更新时间
	DeleteStatus      int32      `gorm:"column:delete_status;not null;comment:是否删除" json:"delete_status"`                                                       // 是否删除
	ProductCategoryID int64      `gorm:"column:product_category_id;not null;comment:商品分类" json:"product_category_id"`                                           // 商品分类
	ProductBrand      string     `gorm:"column:product_brand;not null;comment:商品品牌" json:"product_brand"`                                                       // 商品品牌
	ProductSn         string     `gorm:"column:product_sn;not null;comment:货号" json:"product_sn"`                                                               // 货号
	ProductAttr       string     `gorm:"column:product_attr;not null;comment:商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]" json:"product_attr"` // 商品销售属性:[{"key":"颜色","value":"颜色"},{"key":"容量","value":"4G"}]
}

// TableName OmsCartItem's table name
func (*OmsCartItem) TableName() string {
	return TableNameOmsCartItem
}
