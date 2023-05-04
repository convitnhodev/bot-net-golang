package pkg

import (
	_const "botnetgolang/internal/const"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getOSName() []byte {
	urlCall := fmt.Sprintf("https://api.ipdata.co?api-key=%s", _const.ApiKey)
	resp, err := http.Get(urlCall)
	if err != nil {
		return []byte("")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte("")
	}

	return body
}

func GetInfoOperatingSystem() error {
	result := getOSName()

	file_operatingsystem, _ := os.Create("./storage/operatingsystem.json")
	defer func() {
		if err := file_operatingsystem.Close(); err != nil {
			panic(err)
		}
	}()

	errr := ioutil.WriteFile("./storage/operatingsystem.json", result, 0644)
	return errr
}
