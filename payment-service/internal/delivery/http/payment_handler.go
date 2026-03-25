package http

import (
	"ecom-ms/payment-service/internal/usecase/payment"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	usecase payment.Usecase
}

func NewOrderHandler(u payment.Usecase) *PaymentHandler {
	return &PaymentHandler{
		usecase: u,
	}
}

type CreatePaymentRequest struct {
	UserID      int64   `json:"user_id" binding:"required"`
	TotalAmount float64 `json:"total_amount" binding:"required,min=1"`
}

func RegisterPaymentRoutes(r *gin.Engine, h *PaymentHandler) {
	r.POST("/payment", h.CreatePayment())
}

func (h *PaymentHandler) CreatePayment(c *gin.Context) {
	var req CreatePaymentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	err := h.usecase.CreatePayment(ctx, req.UserID, req.TotalAmount)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "payment created",
	})
}
