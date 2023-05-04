package pkg

import (
	"encoding/json"
	"io/ioutil"
)

func FormatFile(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var m map[string]interface{}
	err = json.Unmarshal(data, &m)
	if err != nil {
		panic(err)
	}

	jsonData, err := json.MarshalIndent(m, "", "  ")
	if err != nil {
		panic(err)
	}

	err = ioutil.WriteFile(path, jsonData, 0644)
	if err != nil {
		panic(err)
	}

}
