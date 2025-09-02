# 🎯 Targeting Engine – Greedy Games Backend Assignment

## 📌 Overview
This project is a **Go microservice** that routes the right campaigns to the right delivery requests based on targeting rules.  
It exposes an HTTP endpoint `/v1/delivery` which matches **active campaigns** against request parameters (`app`, `country`, `os`) and returns the campaigns that qualify.

---

## 🗂️ Project Structure
```
targeting-engine/
├── cmd/
│ └── server/
│ └── main.go # App entrypoint
├── internal/
│ ├── delivery/
│ │ ├── handler.go # Delivery HTTP handler
│ │ └── handler_test.go # Unit tests for delivery
│ ├── campaign/
│ │ ├── matcher.go # Rule-matching logic
│ │ └── matcher_test.go # Unit tests for matcher
│ ├── db/
│ │ └── store.go # In-memory campaigns + rules (later DB/Redis)
│ └── models/
│ └── types.go # Struct definitions
├── migrations/ # (optional) SQL schema for Postgres
├── go.mod / go.sum
└── docker-compose.yml # (optional) Postgres/Redis setup
```

### Campaign
```go
type Campaign struct {
    ID     string `json:"cid"`
    Name   string
    Image  string `json:"img"`
    CTA    string `json:"cta"`
    Status string // "ACTIVE" or "INACTIVE"
}
```


### Targeting Rule
```go
type TargetingRule struct {
    CampaignID       string
    IncludeApps      []string
    ExcludeApps      []string
    IncludeCountries []string
    ExcludeCountries []string
    IncludeOS        []string
    ExcludeOS        []string
}
```

### API Usage
Delivery Endpoint

Request

```GET /v1/delivery?app=com.gametion.ludokinggame&country=us&os=android```


Response (200 OK)
```
[
  {
    "cid": "spotify",
    "img": "https://somelink",
    "cta": "Download"
  },
  {
    "cid": "subwaysurfer",
    "img": "https://somelink3",
    "cta": "Play"
  }
]
```

Response (204 No Content)
```
(no body)
```

Response (400 Bad Request)
```
{ "error": "missing app param" }
```

🧪 Testing

Run all tests:

```
go test ./...
```

Test cases cover:
Rule matching logic (inclusion/exclusion).
Delivery endpoint behavior (200, 204, 400).
