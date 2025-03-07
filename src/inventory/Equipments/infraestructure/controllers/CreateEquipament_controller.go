package equipamentcontrollers

import (
	equipmentusecases "gym-system/src/inventory/Equipments/application/useCases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEquipmentController struct {
	useCase             equipmentusecases.CreateEquipment
	notificationService equipmentusecases.MensajeUseCase // Servicio para enviar mensajes
}

func NewCreateEquipmentController(useCase *equipmentusecases.CreateEquipment, notificationService *equipmentusecases.MensajeUseCase) *CreateEquipmentController {
	return &CreateEquipmentController{useCase: *useCase, notificationService: *notificationService}
}

func (controller *CreateEquipmentController) Execute(c *gin.Context) {
	var requestBody struct {
		Name      string `json:"name"`
		Category  string `json:"category"`
		Condition string `json:"condition"`
	}

	if err := c.ShouldBindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	// Ejecuta la creación del equipo
	err := controller.useCase.Execute(requestBody.Name, requestBody.Category, requestBody.Condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el equipo"})
		return
	}

	// Enviar el mensaje a RabbitMQ
	err = controller.notificationService.SendMessage("PRODUCT", "EQUIPO AGREGADO EXITOSAMENTE")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al enviar el mensaje"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipo agregado exitosamente y mensaje enviado"})
}
