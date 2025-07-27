package models

type Campaign struct {
	ID     string `json:"cid"`
	Name   string
	Image  string `json:"img"`
	CTA    string `json:"cta"`
	Status string // ACTIVE or INACTIVE
}

type TargetingRule struct {
	CampaignID       string
	IncludeApps      []string
	ExcludeApps      []string
	IncludeCountries []string
	ExcludeCountries []string
	IncludeOS        []string
	ExcludeOS        []string
}
