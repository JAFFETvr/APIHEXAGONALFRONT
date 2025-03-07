package equipamentusecases

import "gym-system/src/inventory/Equipments/domain/repository"

type MensajeUseCase struct {
	repo repository.IMessageRepository
}

func NewMensajeUseCase(repo repository.IMessageRepository) *MensajeUseCase {
	return &MensajeUseCase{repo: repo}
}

func (m *MensajeUseCase) SendMessage(queue string, message string) error {
	return m.repo.SendMessage(queue, message)
}
