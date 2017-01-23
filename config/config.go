package config

type info struct {
	Author   string   `json:"author"`
	Info     string   `json:"info"`
	Commands []string `json:"commands"`
}

type response struct {
	Text string `json:"text"`
	Bot  string `json:"bot"`
}

type Event struct {
	Text        string `json:"text"`
	Username    string `json:"username"`
	DisplayName string `json:"display_name"`
}

var (
	Port    = ":8080"
	BotName = "onair-bot"
	Command = `^/onair`
	LocMsc  = "Europe/Moscow"
	BotRes  = response{"Кажется, вещаем. Прямой эфир по ссылке https://radio-t.com/online/", BotName}
	ComRes  = response{"Прямой эфир по ссылке https://radio-t.com/online/", BotName}
	BotInfo = info{
		"Ildar Gilfanov @rabinzon",
		BotName + " - Каждую субботу в 11 вечера, отправляет в чат ссылку на онлайн вещание.",
		[]string{"`/onair` - ответит ссылкой на страницу вещания, независимо от времени."}}
)
