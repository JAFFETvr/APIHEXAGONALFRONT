package routes

import (
	"github.com/gin-gonic/gin"
	equipamentControllers "gym-system/src/inventory/Equipments/infraestructure/controllers"
	equipmentUseCases "gym-system/src/inventory/Equipments/application/useCases"
	"gym-system/src/inventory/Equipments/infraestructure/database" // Asegúrate de que se use correctamente
	"log"      // Para repositorios
)

func SetupRoutesEquipament(r *gin.Engine) {
	// Crear instancia de MySQLEquipament
	dbInstance := equipment.NewMySQLEquipament() // Aquí debes asegurarte de que tu base de datos implemente IEquipamentRepository

	// Crear la conexión a RabbitMQ
	rabbitMQ, err := equipment.NewRabbitMQ()
	if err != nil {
		log.Fatal("Error al conectar a RabbitMQ:", err)
	}

	// Crear el repositorio de mensajes y el servicio de notificación
	messageRepo := rabbitMQ // Usamos RabbitMQ como repositorio de mensajes
	notificationService := equipmentUseCases.NewMensajeUseCase(messageRepo)

	// Crear los controladores
	listEquipamentController := equipamentControllers.NewListEquipmentController(equipmentUseCases.NewListEquipment(dbInstance))
	createEquipamentController := equipamentControllers.NewCreateEquipmentController(
		equipmentUseCases.NewCreateEquipment(dbInstance),
		notificationService, // Pasa el servicio de notificación correctamente
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