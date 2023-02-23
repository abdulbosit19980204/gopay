package main

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("👑GoPay Premium"),
		tgbotapi.NewKeyboardButton("💳Kartalarim"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("💸To'lov"),
		tgbotapi.NewKeyboardButton("💰Balans"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🔀O'tkazmalar"),
		tgbotapi.NewKeyboardButton("📆 To'lovlar Tarixi"),
	),

	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("↙️Kiruvchi hisoblar"),
		tgbotapi.NewKeyboardButton("🌟Saralangan to'lovlar"),

		// 	tgbotapi.NewKeyboardButton("👑Click Premium"),
		// 	tgbotapi.NewKeyboardButton("💳Kartalarim"),

	),

	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("💠Click-Hamyon"),
		tgbotapi.NewKeyboardButton("😍KATTA CASHBACK"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("🧮Mening qarzlarim"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("📍Joylarda to'lov"),
		tgbotapi.NewKeyboardButton("⚙️Sozlamlar"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("💌Biz bilan aloqa"),
		tgbotapi.NewKeyboardButton("🆕Nima Yangiliklar"),
	),
)

func main() {
	bot, err := tgbotapi.NewBotAPI("1069768597:AAHlO1zhlyh7PsTUwLxZ2DkLmPhoj5qK7MM")
	// bot, err := tgbotapi.NewBotAPI("282679704:AAEQoq5g5t2S3YEZ5SxvvSNQVxLS7vYPUpM")
	//gopay 6282679704:AAEQoq5g5t2S3YEZ5SxvvSNQVxLS7vYPUpM token
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore non-Message updates
			continue
		}

		msg := tgbotapi.NewMessage(update.Message.Chat.ID, update.Message.Text)

		switch update.Message.Text {
		case "/start":
			msg.ReplyMarkup = numericKeyboard
		case "close":
			msg.ReplyMarkup = tgbotapi.NewRemoveKeyboard(true)
		}

		if _, err := bot.Send(msg); err != nil {
			log.Panic(err)
		}
	}
}
