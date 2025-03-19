package routes

import (
	"github.com/gin-gonic/gin"
	equipamentControllers "gym-system/src/inventory/Equipments/infraestructure/controllers"
	equipmentUseCases "gym-system/src/inventory/Equipments/application/useCases"
	"gym-system/src/inventory/Equipments/infraestructure/database"
	
)

func SetupRoutesEquipament(r *gin.Engine) {
	dbInstance := equipment.NewMySQLEquipament() 


	// Crear los controladores
	listEquipamentController := equipamentControllers.NewListEquipmentController(equipmentUseCases.NewListEquipment(dbInstance))
	createEquipamentController := equipamentControllers.NewCreateEquipmentController(
		equipmentUseCases.NewCreateEquipment(dbInstance),
	)
	getEquipmentById := equipamentControllers.NewEquipmentByIdController(equipmentUseCases.NewEquipmentById(dbInstance))
	getEquipmentCondition := equipamentControllers.NewEquipmentCondition(equipmentUseCases.NewEquipmentByCondition(dbInstance))
	updateEquipment := equipamentControllers.NewUpdateEquipmentController(equipmentUseCases.NewUpdateEquipment(dbInstance))
	deleteEquipment := equipamentControllers.NewDeleteEquipment(equipmentUseCases.NewDeleteEquipment(dbInstance))

	r.GET("/equipments", listEquipamentController.Execute)
	r.POST("/equipments", createEquipamentController.Execute)
	r.GET("/equipments/:id", getEquipmentById.Execute)
	r.GET("/equipments/condition/:condition", getEquipmentCondition.Execute)
	r.PUT("/equipments/:id", updateEquipment.Execute)
	r.DELETE("/equipments/:id", deleteEquipment.Execute)
}