package handlers

import (
	"net/http"
	"strconv"
	"paymentservice/internal/models"
	"paymentservice/internal/services"

	"github.com/gin-gonic/gin"
)

// PaymentHandler handles payment-related HTTP requests.
type PaymentHandler struct {
	service services.PaymentService
}

// RegisterPaymentRoutes registers payment endpoints.
func RegisterPaymentRoutes(r *gin.Engine, service services.PaymentService) {
	handler := &PaymentHandler{
		service: service,
	}

	paymentGroup := r.Group("/api/payments")
	{
		paymentGroup.POST("", handler.ProcessPayment)
		paymentGroup.GET("/:id", handler.GetPaymentByID)
		paymentGroup.GET("/user/:id", handler.GetPaymentsByUser)
		paymentGroup.POST("/webhook", handler.HandleWebhook)
	}
}

// ProcessPayment godoc
// @Summary Process a new payment
// @Description Process a new payment for a booking
// @Tags payments
// @Accept json
// @Produce json
// @Param payment body models.Payment true "Payment Information"
// @Success 200 {object} models.Payment
// @Failure 400 {object} map[string]interface{}
// @Router /api/payments [post]
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

// GetPaymentByID godoc
// @Summary Get payment details by ID
// @Description Retrieve payment details by payment ID
// @Tags payments
// @Produce json
// @Param id path int true "Payment ID"
// @Success 200 {object} models.Payment
// @Failure 404 {object} map[string]interface{}
// @Router /api/payments/{id} [get]
func (h *PaymentHandler) GetPaymentByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
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

// GetPaymentsByUser godoc
// @Summary Get payment history for a user
// @Description Retrieve all payments for a specific user
// @Tags payments
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {array} models.Payment
// @Failure 404 {object} map[string]interface{}
// @Router /api/payments/user/{id} [get]
func (h *PaymentHandler) GetPaymentsByUser(c *gin.Context) {
	idStr := c.Param("id")
	userID, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid user ID"})
		return
	}

	payments, err := h.service.GetPaymentsByUser(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no payments found for user"})
		return
	}

	c.JSON(http.StatusOK, payments)
}

// HandleWebhook godoc
// @Summary Handle payment gateway webhooks
// @Description Process webhook notifications from the payment gateway
// @Tags payments
// @Accept json
// @Produce json
// @Param payload body map[string]interface{} true "Webhook Payload"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]interface{}
// @Router /api/payments/webhook [post]
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
