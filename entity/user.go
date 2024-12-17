package entity

type UserInfo struct {
	AccountId  int        `json:"accountId"`
	MaskMobile *string    `json:"maskMobile"`
	MallList   []MallInfo `json:"mallList"`
}

type MallInfo struct {
	MallId      int    `json:"mallId"`
	MallName    string `json:"mallName"`
	ManagedType int    `json:"managedType"`
}

// 主页接口卖家中心返回的mall信息
type AccountMallInfo struct {
	IsSemiManagedMall bool   `json:"isSemiManagedMall"`
	Logo              string `json:"logo"`
	MallId            int    `json:"mallId"`
	MallName          string `json:"mallName"`
	MallStatus        int    `json:"mallStatus"`
}
