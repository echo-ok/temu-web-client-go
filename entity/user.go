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
