package pkg

import (
	"fmt"
	"github.com/kbinani/screenshot"
	"image/png"
	"os"
)

func SnapImage() {
	// Lấy ảnh của màn hình đầu tiên
	img, err := screenshot.CaptureDisplay(0)
	if err != nil {
		panic(err)
	}

	// Lưu ảnh vào file png
	file, err := os.Create("./storage/screenshot.png")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		panic(err)
	}

	fmt.Println("Đã chụp ảnh màn hình và lưu vào file screenshot.png")
}
