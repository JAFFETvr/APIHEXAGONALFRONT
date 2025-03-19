package equipamentusecases

import "gym-system/src/inventory/Equipments/domain/repository"

type MensajeUseCase struct {
	repo repository.IMessageRepository
}

func NewMensajeUseCase(repo repository.IMessageRepository) *MensajeUseCase {
	return &MensajeUseCase{repo: repo}
}

func (m *MensajeUseCase) SendEquipmentAddedMessage() error {
	return m.repo.SendMessage("PRODUCT", "EQUIPO AGREGADO EXITOSAMENTE")
}
