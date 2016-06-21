package hubspot

type Property struct {
	Property string      `json:"property"`
	Value    interface{} `json:"value"`
}

type PropertyDeal struct {
	Property string      `json:"name"`
	Value    interface{} `json:"value"`
}
