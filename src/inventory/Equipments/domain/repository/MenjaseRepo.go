package repository

type IMessageRepository interface {
	SendMessage(queue string, message string) error
}