package campaign

import (
	"strings"
	"targeting-engine/internal/models"
)

func MatchRule(r models.TargetingRule, app, country, os string) bool {
	if contains(r.ExcludeApps, app) || contains(r.ExcludeCountries, country) || contains(r.ExcludeOS, os) {
		return false
	}
	if len(r.IncludeApps) > 0 && !contains(r.IncludeApps, app) {
		return false
	}
	if len(r.IncludeCountries) > 0 && !contains(r.IncludeCountries, country) {
		return false
	}
	if len(r.IncludeOS) > 0 && !contains(r.IncludeOS, os) {
		return false
	}
	return true
}

func contains(slice []string, val string) bool {
	for _, s := range slice {
		if strings.EqualFold(s, val) {
			return true
		}
	}
	return false
}
