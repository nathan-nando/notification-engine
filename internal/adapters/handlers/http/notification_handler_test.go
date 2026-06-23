package http_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	adapterHttp "notification-engine/internal/adapters/handlers/http"
)

type mockUseCase struct{}

func (m *mockUseCase) ProcessNotification(level, source, message string) error {
	return nil
}

func TestSendNotification(t *testing.T) {
	e := echo.New()
	reqBody := adapterHttp.NotificationRequest{
		Level:   "INFO",
		Source:  "Test",
		Message: "Test Message",
	}
	body, _ := json.Marshal(reqBody)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/notify", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	handler := adapterHttp.NewNotificationHandler(&mockUseCase{})
	err := handler.SendNotification(c)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if rec.Code != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, rec.Code)
	}
}
