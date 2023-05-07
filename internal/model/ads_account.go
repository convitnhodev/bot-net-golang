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
	AccountStatus          float64           `json:"account_status"`
	AccountID              string            `json:"account_id"`
	CreatedTime            string            `json:"created_time"`
	NextBillDate           string            `json:"next_bill_date"`
	Currency               string            `json:"currency"`
	AdTrustDSL             float64           `json:"adtrust_dsl"`
	TimezoneName           string            `json:"timezone_name"`
	TimezoneOffsetHoursUTC float64           `json:"timezone_offset_hours_utc"`
	BusinessCountryCode    string            `json:"business_country_code"`
	DisableReason          float64           `json:"disable_reason"`
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
	ThresholdAmount float64 `json:"threshold_amount"`
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
	Type            float64 `json:"type"`
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

type ApiResponseBM struct {
	Data []EntityBM `json:"data"`
}

type EntityBM struct {
	Name                   string     `json:"name"`
	IsDisabledForIntegrity bool       `json:"is_disabled_for_integrity_reasons"`
	CreatedTime            string     `json:"created_time"`
	VerificationStatus     string     `json:"verification_status"`
	OwnedAdAccounts        AdAccounts `json:"owned_ad_accounts"`
	ID                     string     `json:"id"`
	LimitAccount           string
	PermittedRoles         []string `json:"permitted_roles"`
}

type AdAccounts struct {
	Data []AdAccount `json:"data"`
}

type AdAccount struct {
	AdTrustDSL           float64        `json:"adtrust_dsl"`
	Balance              string         `json:"balance"`
	IsPrepayAccount      bool           `json:"is_prepay_account"`
	Currency             string         `json:"currency"`
	AccountID            string         `json:"account_id"`
	AccountStatus        float64        `json:"account_status"`
	Name                 string         `json:"name"`
	FundingSourceDetails FundingSource  `json:"funding_source_details"`
	AmountSpent          string         `json:"amount_spent"`
	Insights             Insights       `json:"insights"`
	Adspaymentcycle      AdPaymentCycle `json:"adspaymentcycle"`
	ID                   string         `json:"id"`
}

type FundingSource struct {
	ID   string  `json:"id"`
	Type float64 `json:"type"`
}

type Insights struct {
	Data   []Insight `json:"data"`
	Paging Paging    `json:"paging"`
}

type Insight struct {
	Spend     string `json:"spend"`
	DateStart string `json:"date_start"`
	DateStop  string `json:"date_stop"`
}

type Paging struct {
	Cursors Cursors `json:"cursors"`
}

type Cursors struct {
	Before string `json:"before"`
	After  string `json:"after"`
}

type AdPaymentCycle struct {
	Data   []Threshold `json:"data"`
	Paging Paging      `json:"paging"`
}

type Threshold struct {
	ThresholdAmount float64 `json:"threshold_amount"`
}

type DetaillPage struct {
	Fan          int    `json:"followers_count"`
	Verification string `json:"verification_status"`
}

type Page struct {
	AccessToken string `json:"access_token"`
	PageId      string `json:"id"`
	Name        string `json:"name"`
	Detail      DetaillPage
	Perms       []string `json:"perms"`
}

type ListPage struct {
	Data []Page `json:"data"`
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
