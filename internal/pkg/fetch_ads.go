package pkg

import (
	"io/ioutil"
	"net/http"
)

func FetchFromGraphQl(url string, cookie string, access_token string) ([]byte, error) {
	client := &http.Client{}
	url += access_token
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("cookie", cookie)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
