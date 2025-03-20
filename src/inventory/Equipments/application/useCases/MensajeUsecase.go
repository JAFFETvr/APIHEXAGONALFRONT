package equipamentusecases

import "gym-system/src/inventory/Equipments/domain/repository"

type MensajeUseCase struct {
	repo repository.IMessageRepository
}

func NewMensajeUseCase(repo repository.IMessageRepository) *MensajeUseCase {
	return &MensajeUseCase{repo: repo}
}

func (m *MensajeUseCase) ExecuteEquipmentAdded() error {

	err := m.repo.SendMessage("PRODUCT", "EQUIPO AGREGADO EXITOSAMENTE")
	if err != nil {
		return err
	}

	return nil
}
