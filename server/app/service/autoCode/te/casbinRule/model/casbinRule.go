// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package casbin_rule

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
type CasbinRule struct {
    PType string `orm:"p_type" json:"pType"` // 
    V0 string `orm:"v0" json:"v0"` // 
    V1 string `orm:"v1" json:"v1"` // 
    V2 string `orm:"v2" json:"v2"` // 
    V3 string `orm:"v3" json:"v3"` // 
    V4 string `orm:"v4" json:"v4"` // 
    V5 string `orm:"v5" json:"v5"` // 
}