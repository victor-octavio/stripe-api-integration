package controller

import (
	"net/http"

	"back_end_mirumuh/model"
	"back_end_mirumuh/service"

	"github.com/gin-gonic/gin"
)

type CheckoutController struct {
	Service service.CheckoutService
}

func NewCheckoutController(service service.CheckoutService) *CheckoutController {
	return &CheckoutController{Service: service}
}

func (c *CheckoutController) CreateCheckoutSession(ctx *gin.Context) {
	var input model.CheckoutSessionInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := c.Service.CreateCheckoutSession(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
