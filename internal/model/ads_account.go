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

//type AdsPaymentCycle struct {
//	Data DataAdsPaymentCycle `json:"data"`
//}

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

type ApiResponse struct {
	Data []Account `json:"data"`
}

type Account struct {
	Name                   string            `json:"name"`
	AccountStatus          int               `json:"account_status"`
	AccountID              string            `json:"account_id"`
	CreatedTime            string            `json:"created_time"`
	NextBillDate           string            `json:"next_bill_date"`
	Currency               string            `json:"currency"`
	AdTrustDSL             int               `json:"adtrust_dsl"`
	TimezoneName           string            `json:"timezone_name"`
	TimezoneOffsetHoursUTC int               `json:"timezone_offset_hours_utc"`
	BusinessCountryCode    string            `json:"business_country_code"`
	DisableReason          int               `json:"disable_reason"`
	AdsPaymentCycle        AdsPaymentCycle   `json:"adspaymentcycle"`
	Balance                string            `json:"balance"`
	IsPrepayAccount        bool              `json:"is_prepay_account"`
	Owner                  string            `json:"owner"`
	AllPaymentMethods      AllPaymentMethods `json:"all_payment_methods"`
	TotalPrepayBalance     PrepayBalance     `json:"total_prepay_balance"`
	Insights               Insights          `json:"insights"`
	ID                     string            `json:"id"`
}

type AdsPaymentCycle struct {
	Data []ThresholdAmount `json:"data"`
}

type ThresholdAmount struct {
	ThresholdAmount int `json:"threshold_amount"`
}

type AllPaymentMethods struct {
	PaymentMethodTokens PaymentMethodTokens `json:"payment_method_tokens"`
}

type PaymentMethodTokens struct {
	Data []PaymentMethodData `json:"data"`
}

type PaymentMethodData struct {
	CurrentBalance  Balance `json:"current_balance"`
	OriginalBalance Balance `json:"original_balance"`
	TimeExpire      string  `json:"time_expire"`
	Type            int     `json:"type"`
}

type Balance struct {
	Amount             string `json:"amount"`
	AmountInHundredths string `json:"amount_in_hundredths"`
	Currency           string `json:"currency"`
	OffsettedAmount    string `json:"offsetted_amount"`
}

type PrepayBalance struct {
	Amount             string `json:"amount"`
	AmountInHundredths string `json:"amount_in_hundredths"`
	Currency           string `json:"currency"`
	OffsettedAmount    string `json:"offsetted_amount"`
}

type Insights struct {
	Data []SpendData `json:"data"`
}

type SpendData struct {
	Spend     string `json:"spend"`
	DateStart string `json:"date_start"`
	DateStop  string `json:"date_stop"`
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
