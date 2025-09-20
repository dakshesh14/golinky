package healthcheck

type HealthCheckResponse struct {
	DB    string `json:"db"`
	Cache string `json:"cache"`
}
