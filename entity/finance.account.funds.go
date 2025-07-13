package entity

// FinanceAccountFunds 财务账户资金
type FinanceAccountFunds struct {
	TotalAmount                          string `json:"totalAmount"`
	AvailableBalance                     string `json:"availableBalance"`
	MallDepositThreshold                 string `json:"mallDepositThreshold"`
	MallDepositLimitAmount               string `json:"mallDepositLimitAmount"`
	MallDepositLimitAmountMinusThreshold string `json:"mallDepositLimitAmountMinusThreshold"`
	UsSiteBalance                        any    `json:"usSiteBalance"`
	OtherSiteBalance                     any    `json:"otherSiteBalance"`
	RequireRechargeAmountYuan            string `json:"requireRechargeAmountYuan"`
	MallDepositLimitRecharge             bool   `json:"mallDepositLimitRecharge"`
	WithdrawButtonMode                   int    `json:"withdrawButtonMode"`
	ShouldDisplayDetailPage              bool   `json:"shouldDisplayDetailPage"`
	CallShippingLabelDepositThreshold    struct {
		Value        int    `json:"value"`
		Symbol       string `json:"symbol"`
		CurrencyCode string `json:"currencyCode"`
		Sign         string `json:"sign"`
		DigitalText  string `json:"digitalText"`
	} `json:"callShippingLabelDepositThreshold"`
	Currency          string `json:"currency"`
	CurrencyDesc      string `json:"currencyDesc"`
	CurrencySymbol    string `json:"currencySymbol"`
	TotalAmountFormat struct {
		Value        int    `json:"value"`
		Symbol       string `json:"symbol"`
		CurrencyCode string `json:"currencyCode"`
		Sign         string `json:"sign"`
		DigitalText  string `json:"digitalText"`
	} `json:"totalAmountFormat"`
	AvailableBalanceFormat struct {
		Value        int    `json:"value"`
		Symbol       string `json:"symbol"`
		CurrencyCode string `json:"currencyCode"`
		Sign         string `json:"sign"`
		DigitalText  string `json:"digitalText"`
	} `json:"availableBalanceFormat"`
	MallDepositThresholdFormat struct {
		Value        int    `json:"value"`
		Symbol       string `json:"symbol"`
		CurrencyCode string `json:"currencyCode"`
		Sign         string `json:"sign"`
		DigitalText  string `json:"digitalText"`
	} `json:"mallDepositThresholdFormat"`
	MallDepositLimitAmountFormat struct {
		Value        int    `json:"value"`
		Symbol       string `json:"symbol"`
		CurrencyCode string `json:"currencyCode"`
		Sign         string `json:"sign"`
		DigitalText  string `json:"digitalText"`
	} `json:"mallDepositLimitAmountFormat"`
	MallDepositLimitAmountMinusThresholdFormat struct {
		Value        int    `json:"value"`
		Symbol       string `json:"symbol"`
		CurrencyCode string `json:"currencyCode"`
		Sign         string `json:"sign"`
		DigitalText  string `json:"digitalText"`
	} `json:"mallDepositLimitAmountMinusThresholdFormat"`
	UsSiteBalanceFormat         any `json:"usSiteBalanceFormat"`
	OtherSiteBalanceFormat      any `json:"otherSiteBalanceFormat"`
	RequireRechargeAmountFormat struct {
		Value        int    `json:"value"`
		Symbol       string `json:"symbol"`
		CurrencyCode string `json:"currencyCode"`
		Sign         string `json:"sign"`
		DigitalText  string `json:"digitalText"`
	} `json:"requireRechargeAmountFormat"`
	PlatformGovernDepositInfo struct {
		DynamicDepositModifyStrategyId any  `json:"dynamicDepositModifyStrategyId"`
		DynamicDepositTargetThreshold  any  `json:"dynamicDepositTargetThreshold"`
		NeedConfirmSiteThreshold       bool `json:"needConfirmSiteThreshold"`
	} `json:"platformGovernDepositInfo"`
	TotalBalance int `json:"totalBalance"`
}
