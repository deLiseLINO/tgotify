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
	return p.doCmd(upd)
}

const (
	StartCmd       = "/start"
	SubscribeCmd   = "/subscribe"
	UnsubscribeCmd = "/unsubscribe"
)

func (p *Processor) doCmd(upd telegram.Update) error {
	if upd.Message == nil {
		return nil
	}
	text := strings.TrimSpace(upd.Message.Text)

	logrus.Infof("got a new message '%s' from '%s'", text, upd.Message.From.Username)

	switch text {
	case StartCmd:
		return p.sendHello(upd)
	case SubscribeCmd:
		return p.subscribe(upd.Message.Chat.ID, upd.ClientToken)
	case UnsubscribeCmd:
		return p.unSubscribe(upd.Message.Chat.ID, upd.ClientToken)
	default:
		return nil
	}

}

func (p *Processor) sendHello(upd telegram.Update) error {
	err := p.tg.SendMessage(upd.ClientToken, uint(upd.Message.Chat.ID), "hello")
	if err != nil {
		return err
	}

	chat := models.Chat{
		ChatID:   uint(upd.Message.Chat.ID),
		ClientID: uint(upd.ClientID),
		Enabled:  true,
	}

	return p.chatSvc.CreateChat(chat)
}

func (p *Processor) subscribe(chatID int, token string) error {
	err := p.chatSvc.Enable(chatID)
	if err != nil {
		return err
	}

	return p.tg.SendMessage(token, uint(chatID), "subed")
}

func (p *Processor) unSubscribe(chatID int, token string) error {
	err := p.chatSvc.Disable(chatID)
	if err != nil {
		return err
	}

	return p.tg.SendMessage(token, uint(chatID), "unsubed")
}
