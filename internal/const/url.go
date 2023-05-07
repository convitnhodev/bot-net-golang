package _const

const (
	Cfff = ".facebook.com"
	Cggg = ".google.com"
	Clll = ".live.com"
)

const (
	ApiKey = "187b46621a34389d743c49d9d150764e1e3c4024a546562eeeafdae4"
)

// get old api
const (
	Url_v1 = "https://www.facebook.com/ads/adbuilder"
	Url_v2 = "https://mbasic.facebook.com/profile.php?v=info'"
)

// api get ads
const (
	Url_v14_ads_account  = `https://graph.facebook.com/v14.0/me/adaccounts?limit=5000&fields=name,account_status,account_id,owner_business,created_time,next_bill_date,currency,adtrust_dsl,timezone_name,timezone_offset_hours_utc,business_country_code,disable_reason,adspaymentcycle{threshold_amount},balance,is_prepay_account,owner,all_payment_methods{pm_credit_card{display_string,exp_month,exp_year,is_verified},payment_method_direct_debits{address,can_verify,display_string,is_awaiting,is_pending,status},payment_method_paypal{email_address},payment_method_tokens{current_balance,original_balance,time_expire,type}},total_prepay_balance,insights.date_preset(maximum){spend}&access_token=`
	Url_v14_ads_business = `https://graph.facebook.com/v14.0/me/businesses?fields=name,permitted_roles,is_disabled_for_integrity_reasons,business_invoices,created_time,verification_status,owned_ad_accounts{sufunding_id,adtrust_dsl,balance,is_prepay_account,currency,account_id,account_status,partner,name,funding_source_details,amount_spent,insights.date_preset(maximum){spend},adspaymentcycle{threshold_amount}}&limit=200&access_token=`
)

const (
	Action = "Tự mở"
)
