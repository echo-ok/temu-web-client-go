package entity

type ProductStock struct {
	ProductSkuId       int              `json:"productSkuId"`
	StockType          *int             `json:"stockType"`
	WarehouseStockList []WarehouseStock `json:"warehouseStockList"`
	ShippingMode       int              `json:"shippingMode"`
	SupportSystemSync  *bool            `json:"supportSwitchToSystemSync"`
}

type WarehouseStock struct {
	WarehouseInfo    WarehouseInfo `json:"warehouseInfo"`
	WarehouseDisable bool          `json:"warehouseDisable"`
	StockAvailable   int           `json:"stockAvailable"`
	SiteList         []Site        `json:"siteList"`
}

type WarehouseInfo struct {
	WarehouseDisable bool   `json:"warehouseDisable"`
	WarehouseId      string `json:"warehouseId"`
	WarehouseName    string `json:"warehouseName"`
}

type Site struct {
	SiteName string `json:"siteName"`
	SiteId   int    `json:"siteId"`
}
