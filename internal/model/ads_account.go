package model

type CreditCard struct {
	Display        string `json:"display_string"`
	ExpMonth       string `json:"exp_month"`
	ExpYear        string `json:"exp_year"`
	IsVerification bool   `json:"is_verified"`
}

type DataAdsPaymentCycle struct {
	ThresholdAmount int `json:"threshold_amount"`
}

type AdsPaymentCycle struct {
	Data DataAdsPaymentCycle `json:"data"`
}

type AdsAcount struct {
	AdsPaymentCycle AdsPaymentCycle `json:"adspaymentcycle"`
	Id              string          `json:"account_id"`
	Status          int             `json:"account_status"`
	CreatedTime     string          `json:"created_time"`
	NextBillDate    string          `json:"next_bill_date"`
	Currency        string          `json:"currency"`
	AdTrust         int             `json:"adtrust_dsl"`
	TimeZoneName    string          `json:"timezone_name"`
	TimeZoneHour    int             `json:"timezone_offset_hours_utc"`
	CountryCode     string          `json:"business_country_code"`
	ThreadAmount    string          `json:"thread_amount"`
	Balance         int             `json:"balance"`
	IsPrepayAccount bool            `json:"is_prepay_account"`
	Owner           string          `json:"owner"`
	Spend           string          `json:"spend"`
	CreditCard      []CreditCard
}

type Page struct {
	PageId string
	Name   string
	Fan    int
}

type AccountFbHoldAds struct {
	Id               string
	Name             string
	Coookie          string
	TotalAccountAds  int
	DetailAccountAds []AdsAcount
	TotalPage        int
	DetailPage       []Page
	Business         []Business
}

type Business struct {
	Id                 string
	Name               string
	IsDisable          bool
	CreatedTime        string
	VerificationStatus bool
	Role               []string
	DetailAccountAds   []AdsAcount
}

type AdsAccountObject struct {
	Data []AdsAcount `json:"data"`
}
