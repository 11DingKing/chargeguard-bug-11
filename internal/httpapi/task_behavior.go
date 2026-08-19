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
		_ = json.NewEncoder(w).Encode(result)
		return
	}
	if err != nil {
		http.Error(w, "storage error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(result)
}
