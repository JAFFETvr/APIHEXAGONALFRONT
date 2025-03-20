package equipmentrepository

import "gym-system/src/inventory/Equipments/domain/repository"

type ServerNotification struct {
    rabbitMQ repository.IMessageRepository
}

func NewServerNotification(rabbitMQ repository.IMessageRepository) *ServerNotification {
    return &ServerNotification{rabbitMQ: rabbitMQ}
}

func (s *ServerNotification) SendMessage(topic string, message string) error {
    return s.rabbitMQ.SendMessage(topic, message)
}
