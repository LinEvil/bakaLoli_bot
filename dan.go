package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

var stickerIDs0 []string
var stickerIDs1 []string
var stickerIDs2 []string

func main() {
	botToken := os.Args[1]
	if len(botToken) < 10 {
		fmt.Printf("No botToken!")
		return
	}
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("Bot start")
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	stickerIDs0 = getStickerSet("maimai_dx", bot)
	stickerIDs1 = getStickerSet("nanpuyue_favorite", bot)
	stickerIDs2 = getStickerSet("adashima", bot)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		log.Printf("[%s %d] %s", update.Message.From.UserName, update.Message.From.ID, update.Message.Text)

		if update.Message.IsCommand() {
			switch cmd := update.Message.Command(); cmd {
			case "sets":
				// commandSet(update.Message, bot)

			case "maimai":
				go commandMaimai(update.Message, bot)

			case "qiao":
				go commandQiao(update.Message, bot)

			case "adachi":
				go commandAdashima(update.Message, bot)

			case "shima":
				go commandAdashima(update.Message, bot)

			case "json":
				go commandJson(update.Message, bot)
			}
		}
	}
}

func getStickerSet(setID string, bot *tgbotapi.BotAPI) []string {
	stickerSet, _ := bot.GetStickerSet(tgbotapi.GetStickerSetConfig{Name: setID})
	stickerIDs := []string{}
	for _, s := range stickerSet.Stickers {
		stickerIDs = append(stickerIDs, s.FileID)
	}
	return stickerIDs
}

// func commandSet(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
// 	stickerSetID := message.CommandArguments()
// 	getStickerSet(stickerSetID, bot)
// 	bot.Send(tgbotapi.NewMessage(message.Chat.ID, "success! "+stickerSetID))
// }

func commandMaimai(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	rand.Seed(time.Now().Unix())
	stickerID := stickerIDs0[rand.Intn(len(stickerIDs0))]
	msgSend := tgbotapi.NewStickerShare(message.Chat.ID, stickerID)
	msgSend.ReplyToMessageID = message.MessageID
	bot.Send(msgSend)
}

func commandQiao(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	rand.Seed(time.Now().Unix())
	stickerID := stickerIDs1[rand.Intn(len(stickerIDs1))]
	msgSend := tgbotapi.NewStickerShare(message.Chat.ID, stickerID)
	msgSend.ReplyToMessageID = message.MessageID
	bot.Send(msgSend)
}

func commandAdashima(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	rand.Seed(time.Now().Unix())
	stickerID := stickerIDs2[rand.Intn(len(stickerIDs2))]
	msgSend := tgbotapi.NewStickerShare(message.Chat.ID, stickerID)
	msgSend.ReplyToMessageID = message.MessageID
	bot.Send(msgSend)
}

func commandJson(message *tgbotapi.Message, bot *tgbotapi.BotAPI) {
	if message.ReplyToMessage == nil {
		return
	}
	// log.Printf("%s", message.ReplyToMessage)
	msgSend := tgbotapi.NewMessage(message.Chat.ID, "in dev...")
	msgSend.ReplyToMessageID = message.MessageID
	bot.Send(msgSend)
}
