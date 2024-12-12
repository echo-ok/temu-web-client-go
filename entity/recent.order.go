package entity

type RecentOrder struct {
	OrderList      []RecentOrderItem `json:"orderList"`
	ParentOrderMap ParentOrderInfo   `json:"parentOrderMap"`
}

type RecentOrderItem struct {
	OrderSn            string      `json:"orderSn"`
	SkcId              int         `json:"skcId"`
	OrderPackageStatus interface{} `json:"orderPackageStatus"`
	GoodsId            int         `json:"goodsId"`
	OrderLabel         []struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"orderLabel"`
	OrderPackageInfoList             interface{} `json:"orderPackageInfoList"`
	OrderStatus                      int         `json:"orderStatus"`
	WarehouseName                    string      `json:"warehouseName"`
	Spec                             string      `json:"spec"`
	HitFulfillmentDeliveryRestricted interface{} `json:"hitFulfillmentDeliveryRestricted"`
	FulfillmentQuantity              int         `json:"fulfillmentQuantity"`
	IsWalmartPackage                 bool        `json:"isWalmartPackage"`
	ShippedQuantity                  int         `json:"shippedQuantity"`
	ShippingChannelId                interface{} `json:"shippingChannelId"`
	ThumbUrl                         string      `json:"thumbUrl"`
	GoodsName                        string      `json:"goodsName"`
	SkuId                            string      `json:"skuId"`
	FulfillmentAppealDisplay         interface{} `json:"fulfillmentAppealDisplay"`
	OrderSendInfo                    struct {
		MatchOrderSubPackageCondition bool        `json:"matchOrderSubPackageCondition"`
		LabelTaskBatchId              interface{} `json:"labelTaskBatchId"`
		RefuseSendReason              interface{} `json:"refuseSendReason"`
	} `json:"orderSendInfo"`
	ProductSkuIdList      []string    `json:"productSkuIdList"`
	DeliveryRestrictedDTO interface{} `json:"deliveryRestrictedDTO"`
	Quantity              int         `json:"quantity"`
	IsPartialCancellation bool        `json:"isPartialCancellation"`
	ExtCodeList           []string    `json:"extCodeList"`
	ProductInfoList       []struct {
		ProductSkuId    string `json:"productSkuId"`
		ProductSkcId    string `json:"productSkcId"`
		ProductSpuId    string `json:"productSpuId"`
		ProductQuantity int    `json:"productQuantity"`
	} `json:"productInfoList"`
	UnShippedQuantity    int         `json:"unShippedQuantity"`
	CancelQuantity       int         `json:"cancelQuantity"`
	LackStockApplyStatus interface{} `json:"lackStockApplyStatus"`
	WarehouseId          string      `json:"warehouseId"`
	OrderFulfillmentMode int         `json:"orderFulfillmentMode"`
	IsUStoCA             bool        `json:"isUStoCA"`
	OriginQuantity       int         `json:"originQuantity"`
}

type ParentOrderInfo struct {
	CwBatchOrderInfo              interface{} `json:"cwBatchOrderInfo"`
	UnbindVisitorOrder            bool        `json:"unbindVisitorOrder"`
	ChangeAddressStatus           interface{} `json:"changeAddressStatus"`
	UrgeTag                       interface{} `json:"urgeTag"`
	RegionName1                   string      `json:"regionName1"`
	AfterSalesOrderType           int         `json:"afterSalesOrderType"`
	CancelType                    interface{} `json:"cancelType"`
	ParentShippingTimeStr         interface{} `json:"parentShippingTimeStr"`
	AddressAbormalAppealInfo      interface{} `json:"addressAbormalAppealInfo"`
	SiteName                      string      `json:"siteName"`
	NeedsSignatureServiceReminder bool        `json:"needsSignatureServiceReminder"`
	RegionId1                     int         `json:"regionId1"`
	ParentReceiptTimeStr          interface{} `json:"parentReceiptTimeStr"`
	AddressSnapshotInfo           struct {
		AddressSnapshotId     string      `json:"addressSnapshotId"`
		LogisticsAbnormalType interface{} `json:"logisticsAbnormalType"`
		ErrorType             interface{} `json:"errorType"`
	} `json:"addressSnapshotInfo"`
	PoListDisplayButtonList []struct {
		Type       int         `json:"type"`
		Name       string      `json:"name"`
		LinkUrl    string      `json:"linkUrl"`
		Elsn       int         `json:"elsn"`
		ButtonType string      `json:"buttonType"`
		IsEnabled  bool        `json:"isEnabled"`
		ReasonDesc interface{} `json:"reasonDesc"`
	} `json:"poListDisplayButtonList"`
	UserInfoStatusType           int           `json:"userInfoStatusType"`
	WaybillInfoList              interface{}   `json:"waybillInfoList"`
	ParentShipMode               int           `json:"parentShipMode"`
	ParentOrderSn                string        `json:"parentOrderSn"`
	ParentAfterSalesTag          interface{}   `json:"parentAfterSalesTag"`
	HasUs2CaOrders               bool          `json:"hasUs2CaOrders"`
	ExpectShipLatestTime         int           `json:"expectShipLatestTime"`
	StatusWarnInfo               int           `json:"statusWarnInfo"`
	SellerNoteDTO                interface{}   `json:"sellerNoteDTO"`
	ParentRiskWarningStatus      interface{}   `json:"parentRiskWarningStatus"`
	CanShipping                  bool          `json:"canShipping"`
	IsRemoteAreaOrder            bool          `json:"isRemoteAreaOrder"`
	ParentOrderTimeStr           string        `json:"parentOrderTimeStr"`
	CwBatchOrderTracking         interface{}   `json:"cwBatchOrderTracking"`
	OnlineShippingSwitch         bool          `json:"onlineShippingSwitch"`
	ExpectDeliveryEndTimeStr     string        `json:"expectDeliveryEndTimeStr"`
	IsCombinedShipping           bool          `json:"isCombinedShipping"`
	NeedBuySignServiceReasonList interface{}   `json:"needBuySignServiceReasonList"`
	ParentConfirmTimeStr         string        `json:"parentConfirmTimeStr"`
	ParentPackageStatus          int           `json:"parentPackageStatus"`
	MerchantFeedbackAddressFlag  bool          `json:"merchantFeedbackAddressFlag"`
	MaskedRegionName1            string        `json:"maskedRegionName1"`
	ExpectShipLatestTimeStr      string        `json:"expectShipLatestTimeStr"`
	FulfillmentExtendedDays      interface{}   `json:"fulfillmentExtendedDays"`
	AddressInvisibleDesc         interface{}   `json:"addressInvisibleDesc"`
	MaskedRegionName3            string        `json:"maskedRegionName3"`
	MaskedRegionName2            string        `json:"maskedRegionName2"`
	CombinedShippingList         []interface{} `json:"combinedShippingList"`
	ParentOrderPendingEndTimeStr string        `json:"parentOrderPendingEndTimeStr"`
	ParentOrderStatus            int           `json:"parentOrderStatus"`
}
