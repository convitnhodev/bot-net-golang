package model

type MessageTelegram struct {
	Time            string
	ReceivedDataBot string
	BotType         string
	IPAddress       string
	WindowVersion   string
	UserName        string
	Country         string
	Browser         string
	BrowserProfile  int
	FacebookAccount int
	AdsAccount      int
	Page            int
	Bm              int
	CreditCard      int
	Password        int
	Extension       []string
	Telegram        []string
	Discord         []string
	TimeActive      string
	LastUpdate      string
}
