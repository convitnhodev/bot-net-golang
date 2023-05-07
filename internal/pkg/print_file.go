package pkg

import (
	"fmt"
	"os"
)

func Prinfile(path string,
	browser string,
	cookie string,
	ads_account []map[string]interface{},
	profile string) {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println(err)
		return
	}

	var writer_file string

	writer_file = browser
	writer_file += "\n" + profile
	writer_file += "\n" + "Cookie: " + cookie

	tong := len(ads_account[0])
	writer_file += "\n" + "Tổng số tài khoản quảng cáo: " + string(tong)

	defer file.Close()

}
