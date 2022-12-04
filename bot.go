package durov

import (
	"context"
	"github.com/ayrat404/durov/client"
	"log"
	"time"
)

type TgBot struct {
	client  *client.TgClient
	params  BotParams
	handler Handler
	router  *commandRouter
}

func NewBot(token string, params BotParams) *TgBot {
	tgClient := client.NewClient(token)
	router := newRouter(params.commands, params.fallbackCommand)
	return &TgBot{
		client:  tgClient,
		params:  params,
		handler: composeHandlers(params.middlewares, router.Handle),
		router:  router,
	}
}

func (t *TgBot) Run(ctx context.Context) error {
	if _, err := t.client.GetMe(); err != nil {
		return err
	}

	err := t.setCommands()
	if err != nil {
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
				timeoutCtx, cancel := context.WithTimeout(ctx, time.Second*5)
				<-timeoutCtx.Done()
				cancel()
				continue
			}
			if len(updates) > 0 {
				offset = updates[len(updates)-1].UpdateId + 1
				for _, update := range updates {
					log.Printf("[DEBUG] received message from @%v: %v", update.Message.From.Username, update.Message.Text)
					t.process(&update)
				}
			}
		}
	}
}

func (t *TgBot) process(update *client.Update) {
	request := newRequestContext(update, t.client)
	t.handler(request)
}

func (t *TgBot) setCommands() error {
	params := createSetCommandsParams(t.params.commands)
	_, err := t.client.SetMyCommands(params)
	return err
}

func createSetCommandsParams(commands []BotCommand) *client.SetMyCommandsParams {
	setMyCommandsParams := &client.SetMyCommandsParams{
		Commands: make([]client.BotCommand, 0, len(commands)),
	}
	for _, botCommand := range commands {
		commandDef := botCommand.GetDefinition()
		if !commandDef.Display {
			continue
		}
		commandParams := client.BotCommand{
			Command:     commandDef.Name,
			Description: commandDef.Description,
		}
		setMyCommandsParams.Commands = append(setMyCommandsParams.Commands, commandParams)
	}
	return setMyCommandsParams
}
