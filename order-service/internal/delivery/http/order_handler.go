package http

import (
	"ecom-ms/order-service/internal/usecase/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	usecase order.Usecase
}

func NewOrderHandler(u order.Usecase) *OrderHandler {
	return &OrderHandler{
		usecase: u,
	}
}

type CreateOrderRequest struct {
	ProductID int64   `json:"product_id" binding:"required"`
	Quantity  int     `json:"quantity" binding:"required,min=1"`
	Note      *string `json:"note"`
}

func RegisterOrderRoutes(r *gin.Engine, h *OrderHandler) {
	r.POST("/orders", h.CreateOrder)
}

func (h *OrderHandler) CreateOrder(c *gin.Context) {
	var req CreateOrderRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx := c.Request.Context()

	err := h.usecase.CreateOrder(ctx, req.ProductID, req.Quantity, req.Note)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "order created",
	})
}
