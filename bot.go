package durov

import (
	"context"
	"github.com/ayrat404/durov/client"
	"github.com/pkg/errors"
	"log"
	"time"
)

type TgBot struct {
	client      *client.TgClient
	params      BotParams
	handler     Handler
	cmdExecutor *CommandExecutor
}

func NewBot(token string, params BotParams) *TgBot {
	tgClient := client.NewClient(token)
	executor := newCommandExecutor(params.commands, params.fallbackCommand)
	cmdExecutor := executor.Execute
	if params.customCmdExecutor != nil {
		cmdExecutor = params.customCmdExecutor
	}
	return &TgBot{
		client:      tgClient,
		params:      params,
		handler:     composeHandlers(params.middlewares, cmdExecutor),
		cmdExecutor: executor,
	}
}

func (t *TgBot) Run(ctx context.Context) error {
	if _, err := t.client.GetMe(ctx); err != nil {
		return errors.Wrap(err, "failed to call GetMe")
	}

	err := t.setCommands()
	if err != nil {
		return errors.Wrap(err, "failed to call set commands")
	}

	log.Printf("start getting telegram updates")

	offset := 0
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			updates, err := t.client.GetUpdates(ctx, &client.GetUpdateParams{Timeout: 20, Offset: offset})
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
					t.process(&update)
				}
			}
		}
	}
}

func (t *TgBot) process(update *client.Update) {
	request := NewRequestContext(update, t.client, t.params.commands)
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
