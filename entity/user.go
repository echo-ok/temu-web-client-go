package entity

type UserInfo struct {
	AccountId  uint64     `json:"accountId"`
	MaskMobile *string    `json:"maskMobile"`
	MallList   []MallInfo `json:"mallList"`
}

type MallInfo struct {
	MallId      uint64 `json:"mallId"`
	MallName    string `json:"mallName"`
	ManagedType int    `json:"managedType"`
}

// 主页接口卖家中心返回的mall信息
type MallInfoByKuangjianmaihuo struct {
	IsSemiManagedMall bool   `json:"isSemiManagedMall"`
	Logo              string `json:"logo"`
	MallId            uint64 `json:"mallId"`
	MallName          string `json:"mallName"`
	MallStatus        int    `json:"mallStatus"`
}
