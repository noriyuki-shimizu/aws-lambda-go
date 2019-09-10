package main

import (
	"fmt"

	"./trainDelay"

	"github.com/line/line-bot-sdk-go/linebot"
)

// Line is struct
type Line struct {
	ChannelSecret string
	ChannelToken  string
	Bot           *linebot.Client
}

// SendTextMessage return error
func (r *Line) SendTextMessage(message string, replyToken string) error {
	return r.Reply(replyToken, linebot.NewTextMessage(message))
}

// SendTemplateMessage return error
func (r *Line) SendTemplateMessage(replyToken, altText string, template linebot.Template) error {
	return r.Reply(replyToken, linebot.NewTemplateMessage(altText, template))
}

// Reply return error
func (r *Line) Reply(replyToken string, message linebot.SendingMessage) error {
	if _, err := r.Bot.ReplyMessage(replyToken, message).Do(); err != nil {
		fmt.Printf("Reply Error: %v", err)
		return err
	}
	return nil
}

// NewCarouselColumn return *linebot.CarouselColumn
func (r *Line) NewCarouselColumn(thumbnailImageURL, title, text string, actions ...linebot.TemplateAction) *linebot.CarouselColumn {
	return &linebot.CarouselColumn{
		ThumbnailImageURL: thumbnailImageURL,
		Title:             title,
		Text:              text,
		Actions:           actions,
	}
}

// NewCarouselTemplate return *linebot.CarouselTemplate
func (r *Line) NewCarouselTemplate(columns ...*linebot.CarouselColumn) *linebot.CarouselTemplate {
	return &linebot.CarouselTemplate{
		Columns: columns,
	}
}

// New return error
func (r *Line) New(secret, token string) error {
	r.ChannelSecret = secret
	r.ChannelToken = token

	bot, err := linebot.New(
		r.ChannelSecret,
		r.ChannelToken,
	)
	if err != nil {
		return err
	}

	r.Bot = bot
	return nil
}

// EventRouter return void
func (r *Line) EventRouter(eve []*linebot.Event) {
	for _, event := range eve {
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				r.handleText(message, event.ReplyToken, event.Source.UserID)
			}
		}
	}
}

func (r *Line) handleText(message *linebot.TextMessage, replyToken, userID string) {
	trainDelayText := trainDelay.GetTrainDelayText(message.Text)
	r.SendTextMessage(trainDelayText, replyToken)
}
