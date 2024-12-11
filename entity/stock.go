package entity

type ProductStock struct {
	ProductSkuId       int64            `json:"productSkuId"`
	StockType          *int64           `json:"stockType"`
	WarehouseStockList []WarehouseStock `json:"warehouseStockList"`
	ShippingMode       int64            `json:"shippingMode"`
	SupportSystemSync  *bool            `json:"supportSwitchToSystemSync"`
}

type WarehouseStock struct {
	WarehouseInfo    WarehouseInfo `json:"warehouseInfo"`
	WarehouseDisable bool          `json:"warehouseDisable"`
	StockAvailable   int64         `json:"stockAvailable"`
	SiteList         []Site        `json:"siteList"`
}

type WarehouseInfo struct {
	WarehouseDisable bool   `json:"warehouseDisable"`
	WarehouseId      string `json:"warehouseId"`
	WarehouseName    string `json:"warehouseName"`
}

type Site struct {
	SiteName string `json:"siteName"`
	SiteId   int64  `json:"siteId"`
}
