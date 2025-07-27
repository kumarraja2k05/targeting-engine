package db

import "targeting-engine/internal/models"

var Campaigns = []models.Campaign{
	{ID: "spotify", Name: "Spotify", Image: "https://somelink", CTA: "Download", Status: "ACTIVE"},
	{ID: "duolingo", Name: "Duolingo", Image: "https://somelink2", CTA: "Install", Status: "ACTIVE"},
	{ID: "subwaysurfer", Name: "Subway Surfer", Image: "https://somelink3", CTA: "Play", Status: "ACTIVE"},
	{ID: "netflix", Name: "Netflix", Image: "https://somelink4", CTA: "Watch", Status: "INACTIVE"}, // inactive campaign
}

var Rules = []models.TargetingRule{
	{CampaignID: "spotify", IncludeCountries: []string{"US", "Canada"}},
	{CampaignID: "duolingo", IncludeOS: []string{"Android", "iOS"}, ExcludeCountries: []string{"US"}},
	{CampaignID: "subwaysurfer", IncludeOS: []string{"Android"}, IncludeApps: []string{"com.gametion.ludokinggame"}},
	{CampaignID: "netflix", IncludeOS: []string{"Android"}, IncludeCountries: []string{"India"}},
}
