package http

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"notification-engine/internal/core/usecases"
)

type NotificationHandler struct {
	usecase usecases.NotificationService
}

func NewNotificationHandler(usecase usecases.NotificationService) *NotificationHandler {
	return &NotificationHandler{
		usecase: usecase,
	}
}

// NotificationRequest defines the JSON payload for triggering a notification.
type NotificationRequest struct {
	Level   string `json:"level" example:"INFO"`
	Source  string `json:"source" example:"Quant Engine"`
	Message string `json:"message" example:"Signal generated: BUY BTC/USDT"`
}

// SendNotification godoc
// @Summary Send a push notification
// @Description Sends a notification via configured messengers (e.g., Telegram)
// @Tags notification
// @Accept json
// @Produce json
// @Param request body NotificationRequest true "Notification details"
// @Security ApiKeyAuth
// @Success 200 {object} map[string]string "Success response"
// @Failure 400 {object} map[string]string "Bad request"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /api/v1/notify [post]
func (h *NotificationHandler) SendNotification(c echo.Context) error {
	var req NotificationRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "invalid request format"})
	}

	if req.Message == "" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "message is required"})
	}

	if req.Level == "" {
		req.Level = "INFO"
	}
	if req.Source == "" {
		req.Source = "System"
	}

	err := h.usecase.ProcessNotification(req.Level, req.Source, req.Message)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "notification sent successfully"})
}
