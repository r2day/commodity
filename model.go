package commodity

// MetaModel 元模型
type MetaModel struct {
	// 门店编号
	StoreID string `json:"store_id" bson:"store_id"`
	// 组织编号
	OrganizeID string `json:"organize_id" bson:"organize_id"`
	// 品牌编号
	BrandID string `json:"brand_id" bson:"brand_id"`
	// 部门编号
	DepartmentID string `json:"department_id" bson:"department_id"`
}
