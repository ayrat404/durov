package durov

import (
	"context"
	"github.com/ayrat404/durov/client"
	"log"
	"time"
)

type TgBot struct {
	client *client.TgClient
}

func NewBot(token string) *TgBot {
	client := client.NewClient(token)
	return &TgBot{
		client: client,
	}
}

func (t *TgBot) Run(ctx context.Context) error {
	if _, err := t.client.GetMe(); err != nil {
		return err
	}

	offset := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			updates, err := t.client.GetUpdates(&client.GetUpdateParams{Timeout: 20, Offset: offset})
			if err != nil {
				log.Printf("[ERROR] error getting updates: %s", err)
				time.Sleep(time.Second * 5) // TODO take timeout from config
				continue
			}
			if len(updates) > 0 {
				offset = updates[len(updates)-1].UpdateId + 1
				for _, update := range updates {
					log.Printf("[DEBUG] received message from @%v: %v", update.Message.From.Username, update.Message.Text)
				}
			}
		}
	}
}
