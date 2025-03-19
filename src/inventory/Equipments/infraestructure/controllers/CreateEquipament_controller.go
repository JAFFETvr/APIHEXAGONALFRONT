package equipamentcontrollers

import (
	equipmentusecases "gym-system/src/inventory/Equipments/application/useCases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateEquipmentController struct {
	useCase equipmentusecases.CreateEquipment
}

func NewCreateEquipmentController(useCase *equipmentusecases.CreateEquipment) *CreateEquipmentController {
	return &CreateEquipmentController{useCase: *useCase}
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

	err := controller.useCase.Execute(requestBody.Name, requestBody.Category, requestBody.Condition)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al crear el equipo"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Equipo agregado exitosamente"})
}
