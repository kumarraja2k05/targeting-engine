package delivery

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestDeliveryHandler_MissingParams(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/delivery?country=IN&os=android", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(DeliveryHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusBadRequest {
		t.Errorf("expected 400, got %d", rr.Code)
	}
	expected := `{"error":"missing app param"}`
	if !strings.Contains(rr.Body.String(), expected) {
		t.Errorf("expected error message, got %s", rr.Body.String())
	}
}

func TestDeliveryHandler_NoMatch(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/delivery?app=nonexistent&country=antarctica&os=desktop", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(DeliveryHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusNoContent {
		t.Errorf("expected 204 No Content, got %d", rr.Code)
	}
}

func TestDeliveryHandler_ValidMatch(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1/delivery?app=com.gametion.ludokinggame&country=US&os=android", nil)
	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(DeliveryHandler)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", rr.Code)
	}
	if !strings.Contains(rr.Body.String(), `"cid":"subwaysurfer"`) {
		t.Errorf("expected matching campaign in response, got %s", rr.Body.String())
	}
}
