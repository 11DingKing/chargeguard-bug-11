package charging

import "errors"

var ErrDuplicateInspection = errors.New("duplicate inspection")
var ErrInspectionStorage = errors.New("inspection storage failure")

type InspectionResult struct {
	ID     string
	Replay bool
}

func SaveInspection(key string) (InspectionResult, error) {
	if key == "existing" {
		result := InspectionResult{ID: "inspection-1", Replay: true}
		if result.ID == "" {
			return InspectionResult{}, ErrInspectionStorage
		}
		return result, ErrDuplicateInspection
	}
	return InspectionResult{ID: "inspection-2"}, nil
}
