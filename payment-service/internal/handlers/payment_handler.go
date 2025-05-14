package handlers

import (
	"net/http"
	"strconv"
	"paymentservice/internal/models"
	"paymentservice/internal/services"

	"github.com/gin-gonic/gin"
)

type PaymentHandler struct {
	service services.PaymentService
}

func RegisterPaymentRoutes(r *gin.Engine, service services.PaymentService) {
	handler := &PaymentHandler{service: service}

	p := r.Group("/api/payments")
	{
		p.POST("", handler.ProcessPayment)
		p.GET("/:id", handler.GetPaymentByID)
		p.GET("/user/:id", handler.GetPaymentsByUser)
		p.GET("/barber/:id", handler.GetPaymentsByBarber)
		p.POST("/webhook", handler.HandleWebhook)
	}
}

func (h *PaymentHandler) ProcessPayment(c *gin.Context) {
	var payment models.Payment
	if err := c.ShouldBindJSON(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.ProcessPayment(&payment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) GetPaymentByID(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payment ID"})
		return
	}
	payment, err := h.service.GetPaymentByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "payment not found"})
		return
	}
	c.JSON(http.StatusOK, payment)
}

func (h *PaymentHandler) GetPaymentsByUser(c *gin.Context) {
	uid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}
	payments, err := h.service.GetPaymentsByUser(uid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no payments found for user"})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) GetPaymentsByBarber(c *gin.Context) {
	bid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid barber ID"})
		return
	}
	payments, err := h.service.GetPaymentsByBarber(bid)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no payments found for barber"})
		return
	}
	c.JSON(http.StatusOK, payments)
}

func (h *PaymentHandler) HandleWebhook(c *gin.Context) {
	var payload map[string]interface{}
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.HandleWebhook(payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "webhook processed"})
}
