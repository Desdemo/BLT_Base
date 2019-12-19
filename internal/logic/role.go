package logic

import "BLT_Base/internal/base"

// 查看角色分配用户
func GetUserForRole(rid int64) (u2r []base.User, err error) {
	data, err := base.Common{
		Table: "user_role",
		Data:  u2r,
		Id:    rid,
		Name:  "uid",
	}.IdList()
	err = base.Engine.Where("id in ?", data.([]int64)).Find(u2r)
	return
}

// 获取角色菜单权限
func GetMenuForRole() {

}
