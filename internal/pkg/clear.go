package pkg

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// get all path of sub dir belong paths
func GetDirectories(pathg string) ([]string, error) {
	var dirs []string

	files, err := ioutil.ReadDir(pathg)
	if err != nil {
		return dirs, err
	}

	for _, file := range files {
		if file.IsDir() {
			absPath, err := filepath.Abs(filepath.Join(pathg, file.Name()))
			if err != nil {
				return dirs, err
			}
			dirs = append(dirs, absPath)
		}
	}
	return dirs, nil
}

func DeleteFolderRecursive(directoryPath string) error {
	if _, err := os.Stat(directoryPath); os.IsNotExist(err) {
		// Thư mục không tồn tại
		return nil
	}

	files, err := filepath.Glob(filepath.Join(directoryPath, "*"))
	if err != nil {
		return err
	}

	for _, file := range files {
		if fi, err := os.Stat(file); err == nil {
			if fi.IsDir() {
				// Đệ quy xóa thư mục con
				if err := DeleteFolderRecursive(file); err != nil {
					return err
				}
			} else {
				// Xóa tệp
				if err := os.Remove(file); err != nil {
					return err
				}
			}
		}
	}

	// Xóa thư mục cha
	if err := os.Remove(directoryPath); err != nil {
		return err
	}

	return nil
}

func ReadFiles(dirname string) error {
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return err
	}

	for _, file := range files {
		filename := file.Name()
		if strings.Contains(filename, ".markdown") || strings.Contains(filename, "LICENSE") || strings.Contains(filename, "license") || strings.Contains(filename, "docs") || strings.Contains(filename, "example") || strings.Contains(filename, "test") || strings.Contains(filename, "samples") || strings.Contains(filename, ".d.ts") || strings.Contains(filename, ".txt") || strings.Contains(filename, ".md") {
			path := filepath.Join(dirname, filename)
			if err := os.Remove(path); err != nil {
				return err
			}

			if file.IsDir() {
				if err := os.RemoveAll(path); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func CheckFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
