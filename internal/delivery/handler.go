package delivery

import (
	"encoding/json"
	"net/http"
	"targeting-engine/internal/campaign"
	"targeting-engine/internal/db"
	"targeting-engine/internal/models"
)

// DeliveryHandler handles /v1/delivery endpoint
func DeliveryHandler(w http.ResponseWriter, r *http.Request) {
	app, country, os, err := extractAndValidateParams(r)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	matched := getMatchingCampaigns(app, country, os)
	if len(matched) == 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	writeResponse(w, matched)
}

// extractAndValidateParams reads query params and validates them
func extractAndValidateParams(r *http.Request) (app, country, os string, err error) {
	query := r.URL.Query()

	app = query.Get("app")
	country = query.Get("country")
	os = query.Get("os")

	switch {
	case app == "":
		return "", "", "", NewParamError("app")
	case country == "":
		return "", "", "", NewParamError("country")
	case os == "":
		return "", "", "", NewParamError("os")
	default:
		return app, country, os, nil
	}
}

// getMatchingCampaigns filters active campaigns based on rules
func getMatchingCampaigns(app, country, os string) []models.Campaign {
	var matched []models.Campaign

	for _, c := range db.Campaigns {
		if c.Status != "ACTIVE" {
			continue
		}
		for _, rule := range db.Rules {
			if rule.CampaignID == c.ID && campaign.MatchRule(rule, app, country, os) {
				matched = append(matched, c)
				break
			}
		}
	}

	return matched
}

// writeResponse sends matching campaigns to the client
func writeResponse(w http.ResponseWriter, campaigns []models.Campaign) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	var response []map[string]string
	for _, m := range campaigns {
		response = append(response, map[string]string{
			"cid": m.ID,
			"img": m.Image,
			"cta": m.CTA,
		})
	}

	_ = json.NewEncoder(w).Encode(response)
}

// writeJSONError sends a JSON error with status code
func writeJSONError(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"error": message,
	})
}

// NewParamError returns a formatted error for missing params
func NewParamError(param string) error {
	return &ParamError{Param: param}
}

// ParamError is a custom error type for missing params
type ParamError struct {
	Param string
}

func (e *ParamError) Error() string {
	return "missing " + e.Param + " param"
}
