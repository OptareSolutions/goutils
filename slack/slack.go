package slack

import (
	"github.com/slack-go/slack"
)

type MESSAGE_COLOR string

const (
	COLOR_GOOD    = "good"
	COLOR_WARNING = "warning"
	COLOR_DANGER  = "danger"
	COLOR_NONE    = ""
)

type Receiver struct {
	Webhook string
	Channel string
}

func NewSlackReceiver(webhook, channel string) *Receiver {
	return &Receiver{
		Webhook: webhook,
		Channel: channel,
	}
}

type Message struct {
	Username string
	Icon     string
	Title    string
	Content  string
}

func NewMessage(username, icon, title, content string) *Message {
	return &Message{
		Username: username,
		Icon:     icon,
		Title:    title,
		Content:  content,
	}
}

// SendPlain sends a plain message to slack (without markdown formatting)
func (r Receiver) SendPlain(m Message) error {
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
func (r Receiver) SendMarkdown(m Message, c MESSAGE_COLOR) error {
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
