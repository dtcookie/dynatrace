package services

import (
	"encoding/json"

	"github.com/dtcookie/dynatrace/rest"
	"github.com/dtcookie/dynatrace/rest/credentials"
)

// Service TODO: documentation
type Service struct {
	client *rest.Client
}

// NewService TODO: documentation
// "https://#######.live.dynatrace.com/api/config/v1", "###########"
func NewService(baseURL string, token string) *Service {
	rest.Verbose = false
	credentials := credentials.New(token)
	config := rest.Config{}
	client := rest.NewClient(&config, baseURL, credentials)

	return &Service{client: client}
}

// Update TODO: documentation
func (cs *Service) Update(config *ServiceAnomalyDetectionConfig) error {
	if _, err := cs.client.PUT("/anomalyDetection/services", config, 204); err != nil {
		return err
	}
	return nil
}

// Get TODO: documentation
func (cs *Service) Get() (*ServiceAnomalyDetectionConfig, error) {
	var err error
	var bytes []byte

	if bytes, err = cs.client.GET("/anomalyDetection/services", 200); err != nil {
		return nil, err
	}
	var response ServiceAnomalyDetectionConfig
	if err = json.Unmarshal(bytes, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
