package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"stripe-api-integration/model"
	"stripe-api-integration/service"
)

type CheckoutController struct {
	Service service.CheckoutService
}

func NewCheckoutController(service service.CheckoutService) *CheckoutController {
	return &CheckoutController{Service: service}
}

// CreateCheckoutSession godoc
// @Summary Cria uma sessão de checkout no Stripe
// @Description Cria e retorna os dados/URL de checkout do Stripe a partir do payload informado.
// @Tags checkout
// @Accept json
// @Produce json
// @Param body body model.CheckoutSessionInput true "Dados para criar a sessão de checkout"
// @Success 200 {object} model.CheckoutSessionOutput
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /checkout [post]
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
