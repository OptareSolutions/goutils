package slack

import (
	"os"

	"github.com/slack-go/slack"
)

type MESSAGE_COLOR string

const (
	COLOR_GOOD    = "good"
	COLOR_WARNING = "warning"
	COLOR_DANGER  = "danger"
	COLOR_NONE    = ""
)

type WebhookReceiver struct {
	Webhook string
	Channel string
}

func NewSlackWebhookReceiver(webhook, channel string) *WebhookReceiver {
	return &WebhookReceiver{
		Webhook: webhook,
		Channel: channel,
	}
}

type BotReceiver struct {
	Token   string
	Channel string
}

func NewSlackBotReceiver(token, channel string) *BotReceiver {
	return &BotReceiver{
		Token:   token,
		Channel: channel,
	}
}

type TextMessage struct {
	Username string
	Icon     string
	Title    string
	Content  string
}

func NewTextMessage(username, icon, title, content string) *TextMessage {
	return &TextMessage{
		Username: username,
		Icon:     icon,
		Title:    title,
		Content:  content,
	}
}

type FileMessage struct {
	FilePath  string
	Filename  string
	ChannelID string
}

func NewFileMessage(filePath, filename, channelID string) *FileMessage {
	return &FileMessage{
		FilePath:  filePath,
		Filename:  filename,
		ChannelID: channelID,
	}
}

// SendPlain sends a plain message to slack (without markdown formatting)
func (r WebhookReceiver) SendPlain(m TextMessage) error {
	msg := slack.WebhookMessage{
		Username: m.Username,
		IconURL:  m.Icon,
		Channel:  r.Channel,
		Text:     m.Content,
	}

	err := slack.PostWebhook(r.Webhook, &msg)
	if err != nil {
		return err
	}

	return nil
}

// SendMarkdown sends a markdown formatted message to slack
func (r WebhookReceiver) SendMarkdown(m TextMessage, c MESSAGE_COLOR) error {
	// blocks: send rich text (italics, code lines..)
	// attachments: send files, html, code blocks/markdown...
	att := slack.Attachment{
		Title:      m.Title,
		Color:      string(c),
		Text:       "```" + m.Content + "```",
		MarkdownIn: []string{"text"}, // interpretamos markdown del campo text
	}

	msg := slack.WebhookMessage{
		Username:    m.Username,
		IconURL:     m.Icon,
		Channel:     r.Channel,
		Attachments: []slack.Attachment{att}, // enviar attachments (markdown y tal)
	}

	err := slack.PostWebhook(r.Webhook, &msg)
	if err != nil {
		return err
	}

	return nil
}

// SendFile sends a file to slack
func (r BotReceiver) SendFile(m FileMessage) error {
	f, err := os.Stat(m.FilePath)
	if err != nil {
		return err
	}

	api := slack.New(r.Token)

	uploadParams := slack.UploadFileV2Parameters{
		Filename: m.Filename,
		File:     m.FilePath,
		FileSize: int(f.Size()),
		Channel:  r.Channel,
	}

	_, err = api.UploadFileV2(uploadParams)
	if err != nil {
		return err
	}

	return nil
}
