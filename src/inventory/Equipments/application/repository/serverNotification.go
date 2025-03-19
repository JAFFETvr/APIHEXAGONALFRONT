package equipmentrepository

import "gym-system/src/inventory/Equipments/domain/repository"

type ServerNotification struct {
    rabbitMQ repository.IMessageRepository
}

func NewServerNotification(rabbitMQ repository.IMessageRepository) *ServerNotification {
    return &ServerNotification{rabbitMQ: rabbitMQ}
}

func (s *ServerNotification) NotifyEquipmentAdded() error {
    return s.rabbitMQ.SendMessage("PRODUCT", "EQUIPO AGREGADO EXITOSAMENTE")
}
