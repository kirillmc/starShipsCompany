package consumer

import "context"

// MessageHandler — обработчик сообщений.
type MessageHandler func(ctx context.Context, msg Message) error
