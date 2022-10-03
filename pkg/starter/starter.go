package starter

import (
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type app struct {
	bot *tgbotapi.BotAPI
}

func NewApplication() *app {
	return &app{}
}

func (a *app) setDebugMode() {
	a.bot.Debug = enableDebug()
}

func (a *app) setBot() {
	log.Println("Setting botAPI...")
	log.Println("Setting botAPI...")

	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_BOT_TOKEN"))
	if err != nil {
		log.Fatalln(err)
	}
	a.bot = bot
}

func (a *app) Run() {
	log.Println("Program started...")
	log.Println("Program started...")
	a.setBot()
	a.setDebugMode()
	log.Println("Program stopped...")
	log.Println("Program stopped...")
	log.Println("Program stopped...")
	log.Println("Program stopped...")
}

//--

func enableDebug() bool {
	log.Println("Checking debug flag...")
	log.Println("Checking debug flag...")
	if os.Getenv("TELEGRAM_BOT_DEBUG") == "1" ||
		os.Getenv("TELEGRAM_BOT_DEBUG") == "TRUE" ||
		os.Getenv("TELEGRAM_BOT_DEBUG") == "True" ||
		os.Getenv("TELEGRAM_BOT_DEBUG") == "true" {
		return true
	}
	log.Println("Checking debug flag...")
	return false
}

//--
