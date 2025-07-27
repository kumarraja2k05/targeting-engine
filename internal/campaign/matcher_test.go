package campaign

import (
	"testing"
	"targeting-engine/internal/models"
)

func TestMatchRule_InclusionOnly(t *testing.T) {
	rule := models.TargetingRule{
		IncludeApps:      []string{"com.example.app"},
		IncludeCountries: []string{"IN"},
		IncludeOS:        []string{"Android"},
	}

	match := MatchRule(rule, "com.example.app", "IN", "Android")
	if !match {
		t.Errorf("expected match to be true")
	}
}

func TestMatchRule_Exclusion(t *testing.T) {
	rule := models.TargetingRule{
		ExcludeCountries: []string{"US"},
	}

	match := MatchRule(rule, "anyapp", "US", "Android")
	if match {
		t.Errorf("expected match to be false due to exclusion")
	}
}

func TestMatchRule_NoMatchOnInclusion(t *testing.T) {
	rule := models.TargetingRule{
		IncludeCountries: []string{"IN"},
	}

	match := MatchRule(rule, "app", "US", "Android")
	if match {
		t.Errorf("expected match to be false as country does not match include list")
	}
}
