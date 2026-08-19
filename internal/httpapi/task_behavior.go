package httpapi

import (
	"chargeguard/internal/charging"
	"encoding/json"
	"errors"
	"net/http"
)

func TaskHTTPHandler(w http.ResponseWriter, r *http.Request) {
	result, err := charging.SaveInspection("existing")
	if errors.Is(err, charging.ErrDuplicateInspection) {
		w.Header().Set("Idempotent-Replay", "true")
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(result)
		return
	}
	if errors.Is(err, charging.ErrInspectionStorage) {
		http.Error(w, "storage error", http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, "inspection error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}
