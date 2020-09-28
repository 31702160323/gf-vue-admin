// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package customers

import (
	"server/app/model/admins"

	"github.com/gogf/gf/os/gtime"
)

// RecordNotFound 根据条件判断数据是否存在
// 有数据返回false
// 没数据 true
func RecordNotFound(where ...interface{}) bool {
	return Model.RecordNotFound(where...)
}

func (m *arModel) RecordNotFound(where ...interface{}) bool {
	r, err := m.M.FindOne(where...)
	if r == nil && err == nil {
		return true
	}
	return false
}

// Custom Model
type Customers struct {
	Id                 uint        `orm:"id,primary" json:"ID"`                            // 自增ID
	CreateAt           *gtime.Time `orm:"create_at" json:"CreatedAt"`                      // 创建时间
	UpdateAt           *gtime.Time `orm:"update_at" json:"UpdatedAt"`                      // 更新时间
	DeleteAt           *gtime.Time `orm:"delete_at" json:"DeletedAt"`                      // 删除时间
	CustomerName       string      `orm:"customer_name" json:"customerName"`               // 客户名
	CustomerPhoneData  string      `orm:"customer_phone_data" json:"customerPhoneData"`    // 客户电话
	SysUserId          int         `orm:"sys_user_id" json:"sysUserId"`                    // 负责员工id
	SysUserAuthorityId string      `orm:"sys_user_authority_id" json:"sysUserAuthorityID"` // 负责员工角色
}

type CustomerHasOneAdmin struct {
	*Customers
	Admin *admins.Admin `json:"sysUser"`
}
