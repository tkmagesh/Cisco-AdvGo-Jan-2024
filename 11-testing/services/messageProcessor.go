package services

type MessageProcessor struct {
	messageService MessageService
}

func (mp MessageProcessor) Process(msg string) bool {
	return mp.messageService.Send(msg)
	// return true
}

func NewMessageProcessor(msgService MessageService) MessageProcessor {
	return MessageProcessor{
		messageService: msgService,
	}
}
