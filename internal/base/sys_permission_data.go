package base

// 数据权限
type Data struct {
	Model
	Name         string
	Entity       string
	ActionId     int
	DisplayValue string
	Remark       string // 备注
}
