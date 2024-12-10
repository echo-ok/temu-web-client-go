package entity

type Product struct {
	productId               int
	productSkcId            int
	productName             string
	productType             int
	sourceType              int
	goodsId                 int
	leafCat                 Category
	categories              CategoryTree
	productProperties       []ProductProperty
	sizeTemplateIds         []int
	productTotalSalesVolume int
	extCode                 string
	skcStatus               int
	skcSiteStatus           int
	mainImageUrl            string
	last7DaysSalesVolume    *int
	productSkuSummaries     []ProductSku
	createdAt               int
	productSelected         bool
	hasDecoration           bool
	supplierSourceType      int
}

// 类目
type Category struct {
	CatId     int     `json:"catId"`
	CatName   string  `json:"catName"`
	CatEnName *string `json:"catEnName,omitempty"`
	CatType   int     `json:"catType"`
}

// 类目树
type CategoryTree struct {
	Cat1    Category `json:"cat1"`
	Cat2    Category `json:"cat2"`
	Cat3    Category `json:"cat3"`
	Cat4    Category `json:"cat4"`
	Cat5    Category `json:"cat5"`
	Cat6    Category `json:"cat6"`
	Cat7    Category `json:"cat7"`
	Cat8    Category `json:"cat8"`
	Cat9    Category `json:"cat9"`
	Cat10   Category `json:"cat10"`
	CatType int      `json:"catType"`
}

// 商品属性
type ProductProperty struct {
	TemplatePid      int    `json:"templatePid"`
	Pid              int    `json:"pid"`
	RefPid           int    `json:"refPid"`
	PropName         string `json:"propName"`
	Vid              int    `json:"vid"`
	PropValue        string `json:"propValue"`
	ValueUnit        string `json:"valueUnit"`
	ValueExtendInfo  string `json:"valueExtendInfo"`
	NumberInputValue string `json:"numberInputValue"`
}

// SKU信息
type ProductSku struct {
	ProductSkuId       int                 `json:"productSkuId"`
	ThumbUrl           string              `json:"thumbUrl"`
	ProductSkuSpecList []ProductSkuSpec    `json:"productSkuSpecList"`
	CurrencyType       string              `json:"currencyType"`
	SiteSupplierPrices []SiteSupplierPrice `json:"siteSupplierPrices"`
	ExtCode            string              `json:"extCode"`
	VirtualStock       int                 `json:"virtualStock"`
	TempLockNum        int                 `json:"tempLockNum"`
	Parent             interface{}         `json:"$parent,omitempty"`
}

// SKU规格
type ProductSkuSpec struct {
	ParentSpecId   int     `json:"parentSpecId"`
	ParentSpecName string  `json:"parentSpecName"`
	SpecId         int     `json:"specId"`
	SpecName       string  `json:"specName"`
	UnitSpecName   *string `json:"unitSpecName,omitempty"`
}

// 站点供应商价格
type SiteSupplierPrice struct {
	SiteId        int `json:"siteId"`
	SupplierPrice int `json:"supplierPrice"`
}
