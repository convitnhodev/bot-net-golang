package pkg

import (
	"io/ioutil"
	"path/filepath"
	"strings"
)

func GetDirectoriesV2(path string) ([]string, error) {
	// Đọc danh sách các file và thư mục con trong đường dẫn đã cho
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	// Lọc ra các thư mục có tên chứa chuỗi "Default" hoặc "Profile "
	var dirs []string
	for _, file := range files {
		if file.IsDir() && (strings.Contains(file.Name(), "Default") || strings.Contains(file.Name(), "Profile ")) {
			dir := filepath.Join(path, file.Name())
			dirs = append(dirs, dir)
		}
	}

	return dirs, nil
}

func GetBW(textGoc, texts, texts2 string) string {
	var textClear []string
	if textGoc == "" {
		return ""
	}
	if textGoc != "" {
		textClear = strings.Split(textGoc, texts)
	}
	if len(textClear) > 1 {
		textClear = strings.Split(textClear[1], texts2)
		return textClear[0]
	}
	return ""
}

func ReadTextFile(path string) (string, error) {
	filePath := path
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	text := string(bytes)
	return text, nil
}
