package httpapi

import (
	"net/http/httptest"
	"strings"
	"testing"
)

func TestTaskBehavior(t *testing.T) {
	rr := httptest.NewRecorder()
	TaskHTTPHandler(rr, httptest.NewRequest("POST", "/task", nil))
	if rr.Code != 200 || !strings.Contains(rr.Body.String(), "inspection-1") || rr.Header().Get("Idempotent-Replay") != "true" {
		t.Fatalf("status=%d headers=%v body=%s", rr.Code, rr.Header(), rr.Body.String())
	}
}
