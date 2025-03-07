package equipamentusecases

import "gym-system/src/inventory/Equipments/domain/repository"

type CreateEquipment struct {
	db repository.IEquipamentRepository
}

func NewCreateEquipment(db repository.IEquipamentRepository) *CreateEquipment {
	return &CreateEquipment{db: db}
}

func (ce *CreateEquipment) Execute(cname string, category string, ccondition string) error {
	err := ce.db.Save(cname, category, ccondition)
	if err != nil {
		return err
	}
	return nil
}
