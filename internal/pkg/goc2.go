package pkg

import (
	"os"
	"strings"
)

func getDirectoriesV2(pathg string) ([]string, error) {
	var dirs []string
	f, err := os.Open(pathg)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	files, err := f.Readdir(-1)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if file.IsDir() && (strings.Contains(file.Name(), "Default") || strings.Contains(file.Name(), "Profile")) {
			dirs = append(dirs, file.Name())
		}
	}

	return dirs, nil
}
