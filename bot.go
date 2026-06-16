package main

import (
	"log"
	//"os"

	"github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	// Bot Object make and Error chto-to tam
	bot , err := tgbotapi.NewBotAPI("8953206131:AAEfpVc95VHc-NiAOxRmdyi6_PxJ716ZVe8");
	if err != nil {
		log.Panic("bot create error: ", err);
	}
	// Debug +
	bot.Debug = true;
	log.Printf(bot.Self.UserName);

	// Настраиваем Long Polling
	u := tgbotapi.NewUpdate(0);
	u.Timeout = 60;

	// Channel update 
	updates := bot.GetUpdatesChan(u);

	// Message get and answer
	for update := range updates {
		if update.Message == nil { continue } // if it not a massage ignore (maybe user input in group)
		log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text); // logging anything in console
	
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi!"); // ID - КОМУ ОТПРАВЛЯТЬ TEXT - ЧТО ОТПРАВИТЬ
		msg.ReplyToMessageID = update.Message.MessageID; // rEPLY MESSAGE 



		if _, err := bot.Send(msg); err != nil {
			log.Println(err);
		}

	}

}
