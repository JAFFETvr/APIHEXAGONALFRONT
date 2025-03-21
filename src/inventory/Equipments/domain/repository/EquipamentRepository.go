package repository

type IEquipamentRepository interface {
    Save(cname string, category string, ccondition string) error
    GetAll() ([]map[string]interface{}, error)
    GetById(id int) ([]map[string]interface{}, error)
    GetCondition(condition string) ([]map[string]interface{}, error)
    Update(id int, cname string, category string, ccondition string) error
    Delete(id int) error

}
    
