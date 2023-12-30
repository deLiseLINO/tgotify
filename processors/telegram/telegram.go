package telegram

import (
	"strings"
	telegram "tgotify/client"
	models "tgotify/storage"

	"github.com/sirupsen/logrus"
)

type ChatService interface {
	CreateChat(chat models.Chat) error
	GetToken(chatID int) (string, error)
	Enable(chatID int) error
	Disable(chatID int) error
}

type Processor struct {
	tg      *telegram.Client
	offset  int
	chatSvc ChatService
}

func New(client *telegram.Client, chatSvc ChatService) *Processor {
	return &Processor{
		tg:      client,
		chatSvc: chatSvc,
	}
}

func (p *Processor) Fetch(limit int) ([]telegram.Update, error) {
	updates := p.tg.Updates(p.offset, limit)

	if len(updates) == 0 {
		return nil, nil
	}

	p.offset = updates[len(updates)-1].ID + 1

	return updates, nil
}

func (p *Processor) ProcessMessage(upd telegram.Update) error {
	return p.doCmd(upd.Message.Text, upd.Message.Chat.ID, upd.Message.From.Username, upd.ClientToken)
}

const (
	StartCmd       = "/start"
	SubscribeCmd   = "/subscribe"
	UnsubscribeCmd = "/unsubscribe"
)

func (p *Processor) doCmd(text string, chatID int, username string, token string) error {
	text = strings.TrimSpace(text)

	logrus.Printf("got a new message '%s' from '%s'", text, username)

	switch text {
	// case StartCmd:
	// 	return p.sendHello(chatID, token)
	// case SubscribeCmd:
	// 	return p.subscribe(chatID, token)
	// case UnsubscribeCmd:
	// 	return p.unSubscribe(chatID, token)
	default:
		return nil
	}

}

// func (p *Processor) sendHello(chatID int, token string) error {

// 	err := p.tg.SendMessage(token, uint(chatID), "hello")
// 	if err != nil {
// 		return err
// 	}

// 	return p.chatSvc.Enable(chatID)
// }

// func (p *Processor) subscribe(chatID int, token string) error {
// 	err := p.chatSvc.Enable(chatID)
// 	if err != nil {
// 		return err
// 	}

// 	return p.tg.SendMessage(token, uint(chatID), "subed")
// }

// func (p *Processor) unSubscribe(chatID int, token string) error {
// 	err := p.chatSvc.Disable(chatID)
// 	if err != nil {
// 		return err
// 	}

// 	return p.tg.SendMessage(token, uint(chatID), "unsubed")
// }
