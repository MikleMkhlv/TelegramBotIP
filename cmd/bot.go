package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"teleframBot/internal/reposytory"
	"teleframBot/internal/service"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type Bot struct {
	service *service.Service
}

var numericKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("1"),
		tgbotapi.NewKeyboardButton("2"),
		tgbotapi.NewKeyboardButton("3"),
	),
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("4"),
		tgbotapi.NewKeyboardButton("5"),
		tgbotapi.NewKeyboardButton("6"),
	),
)

func main() {
	//TODO Connect database postgres...
	err := godotenv.Load(".env")
	if err != nil {
		logrus.Fatalf("File .env not found: %s", err)
	}

	db, err := repo()
	if err != nil {
		logrus.Fatalf("error conected to database: %s", err)
	}

	repo := reposytory.NewRepository(db)
	service := service.NewService(repo)
	//TODO ////////////////////////

	bot, err := tgbotapi.NewBotAPI(os.Getenv("BOT_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
			msg.Text = "hello"
			if _, err := bot.Send(msg); err != nil {
				log.Panic(err)

			}
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
		if update.Message.IsCommand() {
			switch update.Message.Command() {
			case "help":
				msg.Text = "I understand /admin and /status."
				bot.Send(msg)

			case "admin":
				a, err := service.GetAdmins(update.Message.Chat.ID, update.Message.From.FirstName, update.Message.From.LastName)
				if err != nil {
					logrus.Print(err)
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "You don't admin")
					bot.Send(msg)
					continue
				}

				msg = tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("Welcome admin: %v", a))
				msg.ReplyMarkup = numericKeyboard
				if _, err = bot.Send(msg); err != nil {
					panic(err)
				}
				continue
			case "status":
				msg.Text = fmt.Sprintf("%v", update.Message.Chat.ID)
			default:
				msg.Text = "I don't know that command"
			}

			//if _, err := bot.Send(msg); err != nil {
			//	log.Panic(err)
			//}

		}
		// Extract the command from the Message.

	}
}

func repo() (*sqlx.DB, error) {
	db, err := reposytory.NewPostgres(reposytory.Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("PORT"),
		UserName: os.Getenv("USERNAME"),
		Password: os.Getenv("PASSWORD"),
		DBName:   os.Getenv("DBNAME"),
		SSLMode:  os.Getenv("SSLMODE"),
	})
	if err != nil {
		return nil, err
	}

	return db, nil

}
